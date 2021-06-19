package lib

import (
	"errors"
	"fmt"
	"os/exec"
)

const tmuxBinary string = "tmux"

func KillSession(sessionName string) ([]byte, error) {
	args := []string{
		"kill-session",
		"-t",
		sessionName,
	}
	return exec.Command(tmuxBinary, args...).Output()
}

func CreateSession(session Session) error {
	args := []string{
		"new",
		"-s",
		session.Name,
		"-d",
		"-n",
		session.Windows[0].Name,
		"-c",
		session.Windows[0].Path,
	}

	if _, err := exec.Command(tmuxBinary, args...).Output(); err != nil {
		return errors.New("error creating session")
	}

	for k, v := range session.Env {
		if _, err := SetSessionEnvironment(session.Name, k, v); err != nil {
			return errors.New("error setting environment")
		}
	}

	if len(session.Windows[0].Cmd) > 0 {
		if _, err := SendCommandToPane(session.Name, session.Windows[0].Name, session.Windows[0].Cmd); err != nil {
			return errors.New("error sending command to initial window")
		}
	}

	if len(session.Windows) > 1 {
		for i := 1; i < len(session.Windows); i++ {
			if _, err := CreateWindow(session.Name, session.Windows[i]); err != nil {
				return errors.New("error creating window")
			}
		}
	}

	return nil
}

func CreateWindow(sessionName string, window Window) ([]byte, error) {
	args := []string{
		"neww",
		"-t",
		sessionName,
		"-n",
		window.Name,
		"-d",
		"-c",
		window.Path,
	}

	stdout, error := exec.Command(tmuxBinary, args...).CombinedOutput()

	if len(window.Cmd) > 0 {
		return SendCommandToPane(sessionName, window.Name, window.Cmd)
	}

	return stdout, error
}

func SetSessionEnvironment(sessionName string, key string, value string) ([]byte, error) {
	args := []string{
		"setenv",
		"-t",
		sessionName,
		key,
		value,
	}
	return exec.Command(tmuxBinary, args...).CombinedOutput()
}

func SetPaneOption(sessionName string, paneName string, option string, value string) ([]byte, error) {
	args := []string{
		"set",
		"-t",
		fmt.Sprintf("%s:%s", sessionName, paneName),
		option,
		value,
	}
	return exec.Command(tmuxBinary, args...).CombinedOutput()
}

func SendCommandToPane(sessionName string, paneName string, command string) ([]byte, error) {
	args := []string{
		"send-keys",
		"-t",
		fmt.Sprintf("%s:%s", sessionName, paneName),
		command,
		"C-m",
	}
	return exec.Command(tmuxBinary, args...).CombinedOutput()
}
