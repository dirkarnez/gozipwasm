package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	err := UnzipSplitFiles("BEYOND.zip")
	if err != nil {
		fmt.Println("Error unzipping files:", err)
	}
}

func UnzipSplitFiles(filename string) error {
	baseFilename := filename[:len(filename)-4] // Remove the ".zip" extension
	zipFiles, err := filepath.Glob(baseFilename + ".zip.*")
	if err != nil {
		return err
	}

	// Create a new file to write the concatenated content
	// dstFile, err := os.Create(baseFilename)
	// if err != nil {
	// 	return err
	// }
	// defer dstFile.Close()

	dstBuffer := new(bytes.Buffer)

	// Concatenate the split files into a single file
	for _, zipFile := range zipFiles {
		srcFile, err := os.Open(zipFile)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		_, err = io.Copy(dstBuffer, srcFile)
		if err != nil {
			return err
		}
	}

	zipReader, err := zip.NewReader(bytes.NewReader(dstBuffer.Bytes()), int64(dstBuffer.Len()))
	if err != nil {
		return err
	}

	// Extract the contents of the ZIP archive
	for _, file := range zipReader.File {
		err := extractFile(file)
		if err != nil {
			return err
		}
	}

	fmt.Println("Unzipping complete.")
	return nil
}

func extractFile(file *zip.File) error {
	// Open the file from the ZIP archive
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Create the destination file
	dstPath := filepath.Join(".", file.Name)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the file contents to the destination file
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}
