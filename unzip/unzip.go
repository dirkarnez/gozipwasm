package unzip

import (
	"archive/zip"
	"bytes"
)

func UnzipSplitFiles(byteArray []byte) ([]*zip.File, error) {
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

	zipReader, err := zip.NewReader(bytes.NewReader(byteArray), int64(len(byteArray)))
	if err != nil {
		return nil, err
	}

	return zipReader.File, nil
}
