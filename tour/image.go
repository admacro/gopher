package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 133, 133)
}

func (img Image) At(x, y int) (cc color.Color) {
	return color.RGBA{55, 55, 111, 111}
}

func main() {
	m := Image{}
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
	err := ioutil.WriteFile("image.html", imageBytes, perm)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("file image.html successfully created")
	}
}
