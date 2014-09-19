package main

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"strings"
)

const (
	singleFileByteLimit = 107374182400 // 1 GB
	chunkSize           = 4096         // 4 KB
)

func copyContents(r io.Reader, w io.Writer) error {
	var size int64
	b := make([]byte, chunkSize)
	for {
		// we limit the size to avoid zip bombs
		size += chunkSize
		if size > singleFileByteLimit {
			return errors.New("file too large, please contact us for assistance")
		}
		// read chunk into memory
		length, err := r.Read(b[:cap(b)])
		if err != nil {
			if err != io.EOF {
				return err
			}
			if length == 0 {
				break
			}
		}
		// write chunk to zip file
		_, err = w.Write(b[:length])
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	target := "temp/test.zip"
	sources := make([]string, 2)
	sources[0] = "temp/a.csv"
	sources[1] = "temp/b.csv"

	fw, _ := os.Create(target)
	defer fw.Close()

	wt := zip.NewWriter(fw)
	defer wt.Close()

	for _, source := range sources {
		index := strings.LastIndex(source, "/")
		fileName := source[index+1:]
		w, _ := wt.Create(fileName)
		rf, _ := os.Open(source)
		copyContents(rf, w)
		rf.Close()
		os.Remove(source)
	}

}
