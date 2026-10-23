package main

import (
	"fmt"
	"io"
	"strings"
)

func ReadString(r io.Reader) (string, error) {
	data := make([]byte, 1024)
	bytesRead, err := r.Read(data)
	if err != io.EOF {
		return "", nil
	}
	return string(data[:bytesRead]), nil
}

func main() {
	input := strings.NewReader("")
	str, err := ReadString(input)
	fmt.Println(str, err)
}
