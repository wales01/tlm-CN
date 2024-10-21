package shell

import (
	"bytes"
	"context"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	ollama "github.com/jmorganca/ollama/api"
	"os"
	"os/exec"
)

var Version string

func Ok() string {
	style := lipgloss.NewStyle()

	style = style.Bold(true)
	style = style.Foreground(lipgloss.Color("2"))

	return style.Render("(ok)")
}

func SuccessMessage(message string) string {
	style := lipgloss.NewStyle()
	style = style.Foreground(lipgloss.Color("2"))
	return style.Render(message)
}

func WarnMessage(message string) string {
	style := lipgloss.NewStyle()
	style = style.Foreground(lipgloss.Color("202"))
	return style.Render(message)
}

func Err() string {
	style := lipgloss.NewStyle()

	style = style.Bold(true)
	style = style.Foreground(lipgloss.Color("9"))

	return style.Render("(err)")
}

func Warn() string {
	style := lipgloss.NewStyle()

	style = style.Bold(true)
	style = style.Foreground(lipgloss.Color("202"))

	return style.Render("(warn)")
}

func Exec2(command string) (*exec.Cmd, *bytes.Buffer, *bytes.Buffer) {
	var stdout, stderr bytes.Buffer
	var cmd *exec.Cmd

	if GetShell() == "powershell" {
		cmd = exec.Command(GetShell(), "-Command", command)
	} else {
		cmd = exec.Command(GetShell(), "-c", command)
	}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	return cmd, &stdout, &stderr
}

func CheckOllamaIsUp(api *ollama.Client) error {
	_, err := api.Version(context.Background())
	if err != nil {
		fmt.Println(Err() + " " + err.Error())
		fmt.Println(Err() + " Ollama connection failed. Please check your Ollama if it's running or configured correctly.")
		os.Exit(-1)
	}
	return nil
}
