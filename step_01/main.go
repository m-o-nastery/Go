package main

import (
	"fmt"
	"io"
	"strings"
)

func Read(r io.Reader, seq []byte) error {
	data := make([]byte, 4)
	_, err := r.Read(data[:])
	for {
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if string(data) == string(seq) {
			fmt.Println("YES")
			break
		}
		copy(data, data[1:])
		_, err = r.Read(data[len(data)-1:])
		fmt.Println(string(data))
	}
	return nil
}

func Contains(r io.Reader, seq []byte) (bool, error) {
	buff := make([]byte, len(seq))
	_, err := r.Read(buff)
    for {
        if err != nil && err != io.EOF {
			return false, err
        }
        if string(buff) == string(seq) {
			return true, nil
		}
        if err == io.EOF {
            return false, nil
		}
		copy(buff, buff[1:])
		_, err = r.Read(buff[len(buff)-1:])
    }
}

func main() {
	 r := strings.NewReader("AAAADDDDVVVVAAADDDVVVAADDVVxxmamay")
	 fmt.Println(Contains(r, []byte("mama")))

}