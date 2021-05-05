package myopengl

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func LoadPng(filepath string) *image.RGBA {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("error opening file")
	}
	im, err := png.Decode(file)
	if err != nil {
		log.Fatal("error decoding file")
	}
	img := image.NewRGBA(im.Bounds())
	draw.Draw(img, img.Bounds(), im, image.Pt(0, 0), draw.Src)
	return img
}
