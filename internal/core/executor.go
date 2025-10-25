package core

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
)

// FindPythonInterpreter locates an available Python interpreter
func FindPythonInterpreter() (string, error) {
	interpreters := []string{"python3", "python"}

	for _, interpreter := range interpreters {
		path, err := exec.LookPath(interpreter)
		if err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("python is not installed or not in your PATH\n\nPlease install Python 3.x from https://www.python.org/downloads/\n\nAfter installation, ensure Python is added to your system PATH")
}

// ExecutePythonCode executes Python code and returns the output
func ExecutePythonCode(ctx context.Context, code string) error {
	pythonPath, err := FindPythonInterpreter()
	if err != nil {
		return err
	}

	tmpFile, err := os.CreateTemp("", "pseudolang_*.py")
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer func() {
		_ = os.Remove(tmpFile.Name())
	}()

	if _, err := tmpFile.WriteString(code); err != nil {
		_ = tmpFile.Close()
		return fmt.Errorf("failed to write Python code to temporary file: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to close temporary file: %w", err)
	}

	cmd := exec.CommandContext(ctx, pythonPath, tmpFile.Name())

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	if stdout.Len() > 0 {
		fmt.Print(stdout.String())
	}

	if stderr.Len() > 0 {
		fmt.Fprint(os.Stderr, stderr.String())
	}

	if err != nil {
		exitMsg := ""
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitMsg = fmt.Sprintf(" (exit code %d)", exitErr.ExitCode())
		}
		return fmt.Errorf("python execution failed%s", exitMsg)
	}

	return nil
}

// ExecutePythonFile executes a Python file directly
func ExecutePythonFile(ctx context.Context, filepath string) error {
	pythonPath, err := FindPythonInterpreter()
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, pythonPath, filepath)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	if stdout.Len() > 0 {
		fmt.Print(stdout.String())
	}

	if stderr.Len() > 0 {
		fmt.Fprint(os.Stderr, stderr.String())
	}

	if err != nil {
		exitMsg := ""
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitMsg = fmt.Sprintf(" (exit code %d)", exitErr.ExitCode())
		}
		return fmt.Errorf("python execution failed%s", exitMsg)
	}

	return nil
}
