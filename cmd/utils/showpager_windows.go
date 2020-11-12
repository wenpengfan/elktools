// +build windows

package utils

import (
	"io"
	"os/exec"
)

func showPagedReader(w io.Writer, r io.Reader) error {
	var (
		pager = "more"
		cmd   *exec.Cmd
	)

	cmd = exec.Command(pager)
	cmd.Stdout = w
	cmd.Stderr = w
	cmd.Stdin = r
	return cmd.Run()
}
