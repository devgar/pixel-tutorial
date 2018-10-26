package main

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 640, 480),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	pic, err := loadPicture("hiking.png")
	if err != nil {
		panic(err)
	}

	// * Mojave fix
	win.SetPos(win.GetPos().Add(pixel.V(0, 1)))

	sprite := pixel.NewSprite(pic, pic.Bounds())

	win.Clear(colornames.Firebrick)

	angle := 0.5

	for !win.Closed() {
		angle += 0.05

		mat := pixel.IM
		mat = mat.Rotated(pixel.ZV, angle)
		mat = mat.Moved(win.Bounds().Center())
		sprite.Draw(win, mat)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
