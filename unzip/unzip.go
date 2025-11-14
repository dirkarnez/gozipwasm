package unzip

import (
	"archive/zip"
	"bytes"
	"strings"

	yekaZip "github.com/yeka/zip"
)

func UnzipSplittedFilesWithPassword(password string, byteArray []byte) ([]*yekaZip.File, error) {
	r, err := yekaZip.NewReader(bytes.NewReader(byteArray), int64(len(byteArray)))
	if err != nil {
		return nil, err
	}

	// f.SetPassword("12345")
	// r.File[0].SetPassword()

	return r.File, nil
}

func UnzipSplittedFiles(password string, byteArray []byte) ([]*zip.File, error) {
	if len(strings.TrimSpace(password)) > 0 {
		// return UnzipSplittedFilesWithPassword(password, byteArray)
	} else {
		zipReader, err := zip.NewReader(bytes.NewReader(byteArray), int64(len(byteArray)))
		if err != nil {
			return nil, err
		}
	}
	// Create a new file to write the concatenated content
	// dstFile, err := os.Create(baseFilename)
	// if err != nil {
	// 	return err
	// }
	// defer dstFile.Close()

	// // Open the concatenated file
	// zipReader, err := zip.OpenReader(baseFilename)
	// if err != nil {
	// 	return err
	// }
	// defer zipReader.Close()

	return zipReader.File, nil
}
