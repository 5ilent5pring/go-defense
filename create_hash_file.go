package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// Specify the folder path
	folderPath := "/path/to/your/folder"

	// List files in the folder
	files, err := listFiles(folderPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Open the output file for writing
	outputFile, err := os.Create("hash.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Generate and write hashes for each file to the output file
	for _, file := range files {
		filePath := filepath.Join(folderPath, file)

		// Generate hashes
		md5Hash, err := generateFileHash(filePath, md5.New())
		if err != nil {
			fmt.Printf("Error generating MD5 hash for file %s: %v\n", file, err)
		}

		sha1Hash, err := generateFileHash(filePath, sha1.New())
		if err != nil {
			fmt.Printf("Error generating SHA-1 hash for file %s: %v\n", file, err)
		}

		sha256Hash, err := generateFileHash(filePath, sha256.New())
		if err != nil {
			fmt.Printf("Error generating SHA-256 hash for file %s: %v\n", file, err)
		}

		// Write hashes to the output file
		fmt.Fprintf(outputFile, "File: %s\n", file)
		fmt.Fprintf(outputFile, "MD5: %s\n", md5Hash)
		fmt.Fprintf(outputFile, "SHA-1: %s\n", sha1Hash)
		fmt.Fprintf(outputFile, "SHA-256: %s\n", sha256Hash)
		fmt.Fprintln(outputFile, "------------------------")
	}

	fmt.Println("Hashes written to hash.txt")
}

func listFiles(folderPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, info.Name())
		}
		return nil
	})
	return files, err
}

func generateFileHash(filePath string, hashAlgo io.Writer) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := io.Copy(hashAlgo, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hashAlgo.(hashAlgoWriter).Sum(nil)), nil
}

type hashAlgoWriter interface {
	io.Writer
	Sum(b []byte) []byte
}
