package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/dirkarnez/gozipwasm/unzip"
)

func main() {
	// if len(os.Args) == 2 {
	// 	err := UnzipSplittedFiles_Desktop(os.Args[1:2][0])
	// 	if err != nil {
	// 		fmt.Println("Error unzipping files:", err)
	// 	}
	// } else {
	// 	log.Fatal("please drag a file with extension .zip.001")
	// }
	UnzipSplittedFilesWithPassword_Desktop("Downloads.zip.001", "123")

}

func UnzipSplittedFiles_Desktop(filename string) error {
	baseFilename := filename[:len(filename)-8] // Remove the ".zip.001" extension
	zipFiles, err := filepath.Glob(baseFilename + ".zip.*")
	if err != nil {
		return err
	}
	dstBuffer := new(bytes.Buffer)

	// Concatenate the split files into a single file
	for _, zipFile := range zipFiles {
		srcFile, err := os.Open(zipFile)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		_, err = io.Copy( /*dstFile*/ dstBuffer, srcFile)
		if err != nil {
			return err
		}
	}

	//zipReader, err := zip.NewReader(bytes.NewReader(dstBuffer.Bytes()), int64(dstBuffer.Len()))
	files, err := unzip.UnzipSplittedFiles(dstBuffer.Bytes())
	if err != nil {
		return err
	}

	// Extract the contents of the ZIP archive
	for _, file := range files {
		err := extractFile(file)
		if err != nil {
			return err
		}
	}

	fmt.Println("Unzipping complete.")
	return nil
}

func UnzipSplittedFilesWithPassword_Desktop(filename string, password string) error {
	baseFilename := filename[:len(filename)-8] // Remove the ".zip.001" extension
	zipFiles, err := filepath.Glob(baseFilename + ".zip.*")
	if err != nil {
		return err
	}
	dstBuffer := new(bytes.Buffer)

	// Concatenate the split files into a single file
	for _, zipFile := range zipFiles {
		srcFile, err := os.Open(zipFile)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		_, err = io.Copy( /*dstFile*/ dstBuffer, srcFile)
		if err != nil {
			return err
		}
	}

	//zipReader, err := zip.NewReader(bytes.NewReader(dstBuffer.Bytes()), int64(dstBuffer.Len()))
	files, err := unzip.UnzipSplittedFilesWithPassword(dstBuffer.Bytes())
	if err != nil {
		return err
	}

	// Extract the contents of the ZIP archive
	for _, f := range files {
		// err := extractFile(file)
		// if err != nil {
		// 	return err
		// }

		if f.IsEncrypted() {
			f.SetPassword(password)
		}

		r, err := f.Open()
		if err != nil {
			return err
		}

		buf, err := io.ReadAll(r)
		if err != nil {
			return err
		}
		defer r.Close()

		fmt.Printf("Size of %v: %v byte(s)\n", f.Name, len(buf))
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
