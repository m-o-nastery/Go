package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func ExtractLog(inputFilename string, start, end time.Time) ([]string, error) {
	f, _ := os.Open(inputFilename)
	flag := false
	logs := []string{}
	defer f.Close()
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines := strings.Split(line, " ")
		date, _ := time.Parse("02.01.2006", lines[0])
		if date.Equal(start) {
			flag = true
		}
		if date.Equal(end) {
			logs = append(logs, line)
			flag = false
		}
		if flag == true {
			logs = append(logs, line)
		}
	}
	if len(logs) == 0{
		return logs, fmt.Errorf("Error")
	}
	return logs, nil
}

// func main() {
// 	start, _ := time.Parse("02.01.2006", "17.12.2022")
// 	end, _ := time.Parse("02.01.2006", "20.12.2022")
// 	fmt.Println(ExtractLog("new_file.txt", start, end))
// }