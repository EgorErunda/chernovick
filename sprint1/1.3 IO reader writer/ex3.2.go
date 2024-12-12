package main

import (
	"io"
)

func ReadString(r io.Reader) (string, error) {
	buf := make([]byte, 1024)
	var data []byte

	for {
		n, err := r.Read(buf)
		if n > 0 {
			data = append(data, buf[:n]...)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
	}
	return string(data), nil
}
