package utils

import (
	"crypto/md5"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
	"strings"
)

const w = 5
const scale = 30

// ExportImage creates the identicon image to filepath
func ExportImage(img *image.RGBA, filepath string) {
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("error creating file")
	}
	defer file.Close()

	png.Encode(file, img)
}

// BuildImage contains the logic to draw the image using a simple and stupid algorithm
func BuildImage(idAsString, location string) error {

	id := []byte(strings.ToLower(idAsString))
	idHash := md5.Sum(id)
	idHashString := fmt.Sprintf("%x", idHash)

	identiconColor := color.RGBA{uint8(idHash[0]), uint8(idHash[1]), uint8(idHash[2]), uint8(idHash[3])}
	identicon := image.NewRGBA(image.Rect(0, 0, (w * scale), (w * scale)))

	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			charAt := idHashString[(i*5)+j+6]
			start := image.Point{i * scale, j * scale}
			end := image.Point{i*scale + scale, j*scale + scale}
			square := image.Rectangle{start, end}

			if int(math.Mod(float64(charAt), 2)) == 0 {
				draw.Draw(identicon, square, &image.Uniform{identiconColor}, image.ZP, draw.Src)
			} else {
				draw.Draw(identicon, square, &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.ZP, draw.Src)
			}
		}
	}

	ExportImage(identicon, location)
	return nil
}
