package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (x MyReader) Read(ans []byte) (int, error) {
	len := len(ans)
	for i := 0; i < len; i++ {
		ans[i] = byte('A')
	}

	return len, nil
}

func main() {
	reader.Validate(MyReader{})
}
