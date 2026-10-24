package main

import "os"

func CopyFilePart(inputFilename, outFilename string, startpos int) error {
	f1, err := os.Open(inputFilename)

	if err != nil {
		return err
	}
	defer f1.Close()

	f1.Seek(int64(startpos), 0)
	buffer := make([]byte, 1024)
	n, err := f1.Read(buffer)

	err = os.WriteFile(outFilename, buffer[:n], 0666)
	if err != nil {
		return err
	}
	return nil
}

// func main() {
// 	CopyFilePart("new_file.txt", "new_file_2.txt", 2)
// }