package main

import (
	"io"
	"os"
	"strings"
)

type Rot13Reader struct {
	r io.Reader
}

func (r Rot13Reader) Read(b []byte) (int, error) {
	s, err := r.r.Read(b)
	if err != nil {
		return s, err
	}
	
	for i, v := range b {
		// uppercase letters
		if v <= 95 && v >= 65 {
			new := v + 13
			if new > 95 {
				new -= 26
			}
			b[i] = new
		}
		// lowercase letters
		if v <= 122 && v >= 97 {
			new := v + 13
			if new > 122 {
				new -= 26
			}
			b[i] = new
		}
	}
	return s, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{s}
	io.Copy(os.Stdout, &r)
}