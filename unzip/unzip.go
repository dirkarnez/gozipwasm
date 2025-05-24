package unzip

import (
	"archive/zip"
	"bytes"

	yekaZip "github.com/yeka/zip"
)

// f.SetPassword("12345")
func UnzipSplitFilesWithPassword(byteArray []byte, password string) ([]*yekaZip.File, error) {
	r, err := yekaZip.NewReader(bytes.NewReader(byteArray), int64(len(byteArray)))
	if err != nil {
		return nil, err
	}

	return r.File, nil

	// for _, f := range r.File {
	// 	if f.IsEncrypted() {
	// 		f.SetPassword("12345")
	// 	}

	// 	r, err := f.Open()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	buf, err := io.ReadAll(r)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer r.Close()

	// 	fmt.Printf("Size of %v: %v byte(s)\n", f.Name, len(buf))
	// }
}

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
