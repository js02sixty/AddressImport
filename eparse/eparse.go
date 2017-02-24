package eparse

import (
	"bufio"
	"io"
	"os"
)

// Parse file
func Parse(source string, destination string) (err error) {

	// open input file
	fi, err := os.Open(source)
	if err != nil {
		return err
	}

	// close fi on exit and check for its returned error
	defer func() error {
		if err = fi.Close(); err != nil {
			return err
		}
		return nil
	}()

	// open output file
	fo, err := os.Create(destination)
	if err != nil {
		return err
	}
	// close fo on exit and check for its returned error
	defer func() error {
		if err = fo.Close(); err != nil {
			return err
		}
		return nil
	}()

	r := bufio.NewReader(fi) // read buffer
	w := bufio.NewWriter(fo) // write buffer

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		var n int
		n, err = r.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		// write a chunk
		if _, err = w.Write(buf[:n]); err != nil {
			return err
		}
	}

	if err = w.Flush(); err != nil {
		return err
	}
	return nil

}
