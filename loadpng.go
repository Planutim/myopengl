package myopengl

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

func LoadPng(filepath string) *image.RGBA {
	file, err := os.Open(filepath)
	defer file.Close()
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

func LoadImage(filepath string) *image.RGBA {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatal("error opening file")
	}
	im, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("error decoding file")
	}
	img := image.NewRGBA(im.Bounds())
	draw.Draw(img, img.Bounds(), im, image.Pt(0, 0), draw.Src)
	return img
}
