// https://tour.golang.org/moretypes/18
package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
)

// Implement Pic.
// It should return a slice of length dy, each element of which is a slice of
// dx 8-bit unsigned integers. When you run the program, it will display your
// picture, interpreting the integers as grayscale (well, bluescale) values.

// The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
// (You need to use a loop to allocate each []uint8 inside the [][]uint8.)
// (Use uint8(intValue) to convert between types.)
func Pic(dx, dy int) (s [][]uint8) {
	s = make([][]uint8, dy)
	for i := range s {
		dxSlice := make([]uint8, dx)
		for j := range dxSlice {
			dxSlice[j] = uint8((dx + dy) / 3)
		}
		s[i] = dxSlice
	}
	return s
}

func main() {
	Show(Pic)
}

// pic
func Show(f func(int, int) [][]uint8) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	WriteHtmlFile(GetImageBytes(m))
}

func GetImageBytes(m image.Image) string {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func WriteHtmlFile(imageStrEncoding string) {
	html := fmt.Sprintf("<img src=\"data:image/png;base64,%v\">", imageStrEncoding)
	imageBytes := []byte(html)
	var perm = os.FileMode(0644)
	err := ioutil.WriteFile("show_pic.html", imageBytes, perm)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("file image.html successfully created")
	}
}
