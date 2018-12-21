package main

import "golang.org/x/tour/reader"

type myReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r myReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func exercise07() {
	reader.Validate(myReader{})
}
