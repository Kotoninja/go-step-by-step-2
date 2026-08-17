package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFilePart(inputFilename, outFileName string, startpos int) error {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	_, err = inputFile.Seek(int64(startpos), io.SeekStart)
	if err != nil {
		return err
	}

	outFile, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, inputFile)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := CopyFilePart("/home/user/code/go-step-by-step-2/step2/text.txt", "/home/user/code/go-step-by-step-2/step2/output.txt", 2); err != nil {
		fmt.Println(err)
	}

}
