package main

import "os"

func ModifyFile(filename string, pos int, val string) {
	f, _ := os.OpenFile(filename, os.O_RDWR, 0600)
	defer f.Close()
	f.Seek(int64(pos), 0)
	f.WriteString(val)
}

// func main() {
// 	filename := "new_file.txt"
// 	ModifyFile(filename, 2, "69")
// }