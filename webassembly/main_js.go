//go:build js && wasm
// +build js,wasm

package main

//go:generate cp $GOROOT/misc/wasm/wasm_exec.js .

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/dirkarnez/gozipwasm/unzip"

	"syscall/js"
)

func UnzipSplitFiles_JavaScript(this js.Value, args []js.Value) interface{} {
	dstBuffer := new(bytes.Buffer)

	for _, arg := range args {
		buffer := make([]byte, arg.Length())
		js.CopyBytesToGo(buffer, arg)

		_, err := io.Copy(dstBuffer, bytes.NewReader(buffer))
		if err != nil {
			return err
		}
	}

	files, err := unzip.UnzipSplitFiles(dstBuffer.Bytes())
	if err != nil {
		return err
	}

	var jsArray = js.Global().Get("Array").New(len(files))

	// Extract the contents of the ZIP archive
	for i, file := range files {
		zippedFile, err := file.Open()
		if err != nil {
			return err
		}
		defer zippedFile.Close()

		// Read the contents of the file into a []byte slice
		fileBytes, err := ioutil.ReadAll(zippedFile)
		if err != nil {
			return err
		}

		dst := js.Global().Get("Uint8Array").New(len(fileBytes))

		js.CopyBytesToJS(dst, fileBytes)

		jsArray.SetIndex(i, dst)
	}

	// // 计算md5的值
	// bytes, _ := dirkCaf.ConvertCafToMidi(buffer)

	// dst := js.Global().Get("Uint8Array").New(len(bytes))
	// js.CopyBytesToJS(dst, bytes)

	return jsArray
}

func main() {
	js.Global().Set("UnzipSplitFilesGo", js.FuncOf(UnzipSplitFiles_JavaScript))
	select {} // block the main thread forever
}
