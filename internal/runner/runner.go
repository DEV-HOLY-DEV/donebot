package runner

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// RunCommand executes a command and prints its standard output and error.
func RunCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running: %s %s\n", name, strings.Join(args, " "))
	return cmd.Run()
}
