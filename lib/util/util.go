package util

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func ServerLog(msg string) {
	log.Println(color.BlueString("[grpc-server] => " + msg))
}

func ClientLog(msg string) {
	log.Println(color.YellowString("[grpc-client] => " + msg))
}

func ErrLog(msg string) {
	log.Println(color.RedString("error: " + msg))
}

func ReadFilesFromDir(dir string) ([][]byte, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var fileContents [][]byte
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(dir, file.Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("Failed to read file %s: %v", filePath, err)
			continue
		}

		fileContents = append(fileContents, content)
	}

	return fileContents, nil
}

func WriteFile(directory, fileName, content string) error {
	// Create the directory if it doesn't exist
	err := os.MkdirAll(directory, 0755)
	if err != nil {
		return err
	}

	// Create the full file path by joining the directory path and file name
	filePath := filepath.Join(directory, fileName)

	// Open the file for writing (creates the file if it doesn't exist)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
