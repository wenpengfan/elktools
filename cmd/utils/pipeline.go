package utils

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/desertbit/grumble"
)

func PipelineFlags(f *grumble.Flags) {
	f.BoolL("number", false, "number all output lines")
	f.BoolL("less", false, "opposite of more")
	f.BoolL("wc", false, "print the newline counts")
	f.StringL("grep", "", "print lines matching a pattern")
}

func PipelinePrintln(data []byte, c *grumble.Context) error {
	number := c.Flags.Bool("number")
	less := c.Flags.Bool("less")
	wc := c.Flags.Bool("wc")
	grep := c.Flags.String("grep")

	var (
		n   int
		out strings.Builder
		buf *bufio.Reader
	)

	buf = bufio.NewReader(bytes.NewBuffer(data))
	for {
		line, err := buf.ReadString('\n')
		if err != nil && !errors.Is(err, io.EOF) {
			return err
		}
		line = strings.TrimSpace(line)
		if errors.Is(err, io.EOF) && line == "" {
			break
		}
		if !strings.Contains(line, grep) {
			continue
		}

		if number {
			out.WriteString(fmt.Sprintf("%-4d ", n))
		}
		out.WriteString(fmt.Sprintf("%s\n", line))
		n++
	}

	if wc {
		c.App.Println(n)
		return nil
	}

	if less {
		ShowPaged(os.Stdout, out.String())
		return nil
	}
	c.App.Println(out.String())
	return nil
}
