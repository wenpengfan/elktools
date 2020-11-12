package utils

import (
	"io"
	"strings"
)

func ShowPaged(w io.Writer, text string) error {
	return showPagedReader(w, strings.NewReader(text))
}

func ShowPagedReader(w io.Writer, r io.Reader) error {
	return showPagedReader(w, r)
}
