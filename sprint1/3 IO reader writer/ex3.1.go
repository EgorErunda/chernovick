package main

import (
	"errors"
	"io"
)

func WriterString(s string, w io.Writer) error {
	data := []byte(s)
	n, err := w.Write(data)
	if err != nil {
		return errors.New("ERROR")
	}

	if n != len(data) {
		return io.ErrShortWrite
	}
	return nil
}
