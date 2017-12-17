package tool

import (
	"image"
	"image/png"
	"os"
)

//LoadImage load image from path
func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

//SaveImage save image as png
func SaveImage(path string, img image.Image) (err error) {
	imgfile, err := os.Create(path)
	defer imgfile.Close()

	err = png.Encode(imgfile, img)
	return
}
