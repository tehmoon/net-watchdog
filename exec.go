package main

import (
	"os/exec"
	"strings"
	"os"
	"io"
)

func execString(command string) (error) {
	return execve(strings.Split(command, " "), []string{}, nil, os.Stdout, os.Stderr)
}

func execve(command []string, env []string, stdin io.Reader, stdout, stderr io.Writer) (error) {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Env = env
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	return cmd.Run()
}
