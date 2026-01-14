<div align="center">

# 🌙 Dream Journal CLI

**Capture your subconscious. Decrypt your dreams.**

[![Release](https://img.shields.io/github/v/release/fardhanrasya/dream-journal?color=f472b6&labelColor=333&style=flat-square)](https://github.com/fardhanrasya/dream-journal/releases)
[![License](https://img.shields.io/github/license/fardhanrasya/dream-journal?color=a78bfa&labelColor=333&style=flat-square)](LICENSE)
[![Go](https://img.shields.io/badge/Maintained%20with-Go-00ADD8?style=flat-square&logo=go)](https://golang.org/)

<p align="center">
  <img src="https://media1.giphy.com/media/v1.Y2lkPTc5MGI3NjExMjRybHpmNjZ6aW53aW53aW53aW53aW53aW53aW53aW53aW53aQ/LMc7w0qQy0qQy0qQ/giphy.gif" alt="Dream TUI" width="600" />
</p>

*A minimal, blazing fast CLI tool to record, manage, and explore your dreams from the comfort of your terminal.*

[Features](#-features) • [Installation](#-installation) • [Usage](#-usage) • [TUI Mode](#-interactive-tui) • [Contributing](#-contributing)

</div>

---

## ✨ Features

- **🚀 Quick Entry**: Capture dreams instantly before they fade away.
- **🖥️ Interactive TUI**: Beautiful, keyboard-driven interface built with [Bubble Tea](https://github.com/charmbracelet/bubbletea).
- **📝 Editor Integration**: Writes dream content using your favorite `$EDITOR` (Vim, Nano, VS Code, Notepad).
- **🔍 Full-Text Search**: Instantly find that one dream about "flying pizza".
- **🏷️ Auto-Title**: Automatically generates titles if you're too sleepy to think of one.
- **🔒 Local & Private**: All data stored locally in a lightweight SQLite database.

## 📦 Installation

### One-Line Install ⚡
**Windows (PowerShell):**
```powershell
iwr https://raw.githubusercontent.com/fardhanrasya/dream-journal/main/install.ps1 -useb | iex
```

**Linux / MacOS:**
```bash
curl -sL https://raw.githubusercontent.com/fardhanrasya/dream-journal/main/install.sh | bash
```

### Manual Install
If you prefer not to use the script, you can:
1. Download binary from [Releases](https://github.com/fardhanrasya/dream-journal/releases).
2. Or build from source: `go install fardhan.dev/dreamjournal@latest`

*Or clone and build:*
```bash
git clone https://github.com/fardhanrasya/dream-journal.git
cd dream-journal
go build -o dream
```

## 🚀 Usage

### Quick Add
```bash
dream add "I was swimming in a sea of orange soda..."
# Opens your default editor if no content is provided:
dream add
```

### List Dreams
```bash
dream list
```

### Edit a Dream
```bash
dream edit <ID>
```

### Delete a Dream
```bash
dream delete <ID>
```

## 🎨 Interactive TUI

Launch the full experience with:
```bash
dream tui
```

| Key | Action |
| :--- | :--- |
| `↑` / `↓` / `j` / `k` | Navigate list |
| `/` | Search dreams |
| `Enter` | View details |
| `a` / `n` | Add new dream |
| `e` | Edit selected dream |
| `d` / `x` | Delete selected dream |
| `Esc` / `q` | Go back / Quit |

## 🛠️ Configuration

The database is stored at `~/dream-journal.db` by default.

To set your preferred editor, ensure the `EDITOR` environment variable is set in your shell (e.g., in `.bashrc` or `.zshrc`):
```bash
export EDITOR=vim
```
*On Windows, it defaults to Notepad if not set.*

## 🤝 Contributing

Contributions are welcome! Feel free to open an issue or submit a Pull Request.

1. Fork it
2. Create your feature branch (`git checkout -b feature/cool-feature`)
3. Commit your changes (`git commit -am 'Add some cool feature'`)
4. Push to the branch (`git push origin feature/cool-feature`)
5. Create a new Pull Request

## 📄 License

MIT © [Fardhan Rasya](https://github.com/fardhanrasya)
