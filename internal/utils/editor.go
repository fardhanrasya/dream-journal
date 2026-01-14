package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func OpenEditor(initialContent string) (string, error) {
	cmd, tmpfile, err := PrepareEditor(initialContent)
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpfile)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error running editor: %v", err)
	}

	return ReadEditorFile(tmpfile)
}

func PrepareEditor(initialContent string) (*exec.Cmd, string, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		if strings.Contains(strings.ToLower(os.Getenv("OS")), "windows") {
			editor = "notepad"
		} else {
			editor = "vim"
		}
	}

	tmpfile, err := ioutil.TempFile("", "dream_*.md")
	if err != nil {
		return nil, "", fmt.Errorf("error creating temp file: %v", err)
	}

	if initialContent != "" {
		if err := ioutil.WriteFile(tmpfile.Name(), []byte(initialContent), 0644); err != nil {
			tmpfile.Close()
			os.Remove(tmpfile.Name())
			return nil, "", fmt.Errorf("error writing initial content: %v", err)
		}
	}
	
	tmpfile.Close() 

	cmd := exec.Command(editor, tmpfile.Name())
	return cmd, tmpfile.Name(), nil
}

func ReadEditorFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading temp file: %v", err)
	}
	return string(content), nil
}
