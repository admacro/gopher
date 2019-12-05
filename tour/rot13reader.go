package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (ttl int, err error) {
	rb := make([]byte, 8)
	for {
		// https://stackoverflow.com/questions/59191346/different-behaviour-of-break
		// -vs-return-in-infinite-loop-when-implementing-io-re
		n, readErr := rr.r.Read(rb) // variable shadowing if readErr is named err
		err = readErr
		if err == nil {
			for i, c := range rb[:n] {
				b[i+ttl] = decodeRot13(c)
			}
			ttl += n
		} else if err == io.EOF {
			break
		}
	}
	return ttl, err
}

func decodeRot13(c byte) byte {
	if c >= 97 && c <= 122 { // a-z: 97 122
		c += 13
		if c > 122 {
			c -= 26
		}
	} else if c >= 65 && c <= 90 { // A-Z: 65 90
		c += 13
		if c > 90 {
			c -= 26
		}
	}
	return c
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
