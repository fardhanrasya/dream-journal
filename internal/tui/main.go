package tui

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"fardhan.dev/dreamjournal/internal/model"
	"fardhan.dev/dreamjournal/internal/repository"
	"fardhan.dev/dreamjournal/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	dream model.Dream
}

func (i item) Title() string       { return i.dream.Title }
func (i item) Description() string { return i.dream.CreatedAt.Format("2006-01-02 15:04") }
func (i item) FilterValue() string { return i.dream.Title + " " + i.dream.Content }

type editorFinishedMsg struct {
	err      error
	filename string
	isEdit   bool
	dreamID  int
}

type Model struct {
	list     list.Model
	repo     *repository.DreamRepository
	selected *model.Dream
	quitting bool
	err      error
}

func Start(db *sql.DB) error {
	repo := repository.NewDreamRepository(db)
	dreams, err := repo.GetDreams()
	if err != nil {
		return err
	}

	items := make([]list.Item, len(dreams))
	for i, d := range dreams {
		items[i] = item{dream: d}
	}

	delegate := list.NewDefaultDelegate()
	l := list.New(items, delegate, 0, 0)
	l.Title = "Dream Journal"
	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(key.WithKeys("a", "n"), key.WithHelp("a/n", "add dream")),
			key.NewBinding(key.WithKeys("e"), key.WithHelp("e", "edit dream")),
			key.NewBinding(key.WithKeys("d", "x"), key.WithHelp("d/x", "delete dream")),
		}
	}

	m := Model{
		list: l,
		repo: repo,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := lipgloss.NewStyle().Margin(1, 2).GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	
	case editorFinishedMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, nil
		}
		
		content, err := utils.ReadEditorFile(msg.filename)
		os.Remove(msg.filename) // Clean up
		if err != nil {
			m.err = err
			return m, nil
		}

		if strings.TrimSpace(content) == "" {
			return m, nil // Canceled
		}

		if msg.isEdit {
			// Edit existing
			dream, err := m.repo.GetDreamByID(msg.dreamID)
			if err != nil {
				m.err = err
				return m, nil
			}
			dream.Content = content
			// Optionally update title logic? Keeping simple for now
			if err := m.repo.UpdateDream(dream); err != nil {
				m.err = err
				return m, nil
			}
			m.selected = dream // Update selected view
		} else {
			// Add new
			title := utils.GenerateAutoTitle(content)
			dream := &model.Dream{
				Title:   title,
				Content: content,
			}
			if err := m.repo.CreateDream(dream); err != nil {
				m.err = err
				return m, nil
			}
		}
		// Refresh list
		return m.refreshDreams()

	case tea.KeyMsg:
		if m.selected != nil {
			switch msg.String() {
			case "esc", "q":
				m.selected = nil
				return m, nil
			case "e":
				// Edit from detail view
				return m, m.openEditor(m.selected.Content, true, m.selected.ID)
			}
		} else {
			if m.list.FilterState() == list.Filtering {
				break
			}
			switch msg.String() {
			case "ctrl+c", "q":
				m.quitting = true
				return m, tea.Quit
			case "enter":
				i, ok := m.list.SelectedItem().(item)
				if ok {
					m.selected = &i.dream
				}
				return m, nil
			case "a", "n":
				return m, m.openEditor("", false, 0)
			case "e":
				i, ok := m.list.SelectedItem().(item)
				if ok {
					return m, m.openEditor(i.dream.Content, true, i.dream.ID)
				}
			case "d", "x":
				i, ok := m.list.SelectedItem().(item)
				if ok {
					if err := m.repo.DeleteDream(i.dream.ID); err != nil {
						m.err = err
					}
					return m.refreshDreams()
				}
			}
		}
	}

	if m.selected == nil {
		m.list, cmd = m.list.Update(msg)
	}
	return m, cmd
}

func (m Model) View() string {
	if m.selected != nil {
		titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")).MarginBottom(1)
		dateStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).MarginBottom(1)
		contentStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("252"))

		return lipgloss.NewStyle().Margin(1, 2).Render(
			fmt.Sprintf(
				"%s\n%s\n\n%s\n\n(Press e to edit, esc or q to go back)",
				titleStyle.Render(m.selected.Title),
				dateStyle.Render(m.selected.CreatedAt.Format("2006-01-02 15:04")),
				contentStyle.Render(m.selected.Content),
			),
		)
	}
	return lipgloss.NewStyle().Margin(1, 2).Render(m.list.View())
}

func (m Model) openEditor(content string, isEdit bool, id int) tea.Cmd {
	cmd, filename, err := utils.PrepareEditor(content)
	if err != nil {
		return func() tea.Msg { return editorFinishedMsg{err: err} }
	}
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return editorFinishedMsg{err: err, filename: filename, isEdit: isEdit, dreamID: id}
	})
}

func (m Model) refreshDreams() (Model, tea.Cmd) {
	dreams, err := m.repo.GetDreams()
	if err != nil {
		return m, func() tea.Msg { return nil } // Handle error properly?
	}
	items := make([]list.Item, len(dreams))
	for i, d := range dreams {
		items[i] = item{dream: d}
	}
	cmd := m.list.SetItems(items)
	return m, cmd
}
