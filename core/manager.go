package core

import (
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

// Manager handles command running
type Manager struct {
	logger *logrus.Logger
}

// NewManager creates a new instance of Manager
func NewManager(logger *logrus.Logger) *Manager {
	return &Manager{logger}
}

func (m *Manager) commandForOS(commands ...string) *exec.Cmd {
	cmd := exec.Command(commands[0], commands[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd
}

// RunCmd runs a given docker command
func (m *Manager) RunCmd(command, containerID string) {
	switch command {
	case "exec":
		options := []string{
			"/bin/bash",
			"/bin/zsh",
			"/bin/sh",
		}
		for _, bin := range options {
			err := m.commandForOS("docker", "exec", "-it", containerID, bin).Run()
			if err == nil {
				break
			}
		}
	case "logs":
		cmd := exec.Command("docker", "logs", containerID)
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			m.logger.Fatal(err)
		}
	case "stop":
		cmd := exec.Command("docker", "stop", containerID)
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			m.logger.Fatal(err)
		}
	case "exit":
		os.Exit(0)
	}
}
