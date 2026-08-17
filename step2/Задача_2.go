package main

import (
	"bufio"
	"fmt"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	f, err := os.Open(inputFilename)
	if err != nil {
		return ""
	}

	defer f.Close()

	fileScanner := bufio.NewScanner(f)

	line := 1
	for fileScanner.Scan() {
		if line == lineNum {
			return fileScanner.Text()
		}
		line++
	}

	if errScanner := fileScanner.Err(); errScanner != nil {
		return ""
	}

	return ""
}

// func main() {
// 	fmt.Println(LineByNum("/home/user/code/go-step-by-step-2/step2/text.txt", 0))

// }
