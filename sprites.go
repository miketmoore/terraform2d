package terraform2d

import (
	"fmt"
	"image"
	"math"
	"os"

	"github.com/faiface/pixel"
)

func loadPicture(path string) pixel.Picture {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Could not open the picture:")
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Could not decode the picture:")
		fmt.Println(err)
		os.Exit(1)
	}
	return pixel.PictureDataFromImage(img)
}

// LoadAndBuildSpritesheet this is a map of pixel engine sprites
func LoadAndBuildSpritesheet(path string, tileSize float64) map[int]*pixel.Sprite {
	pic := loadPicture(path)

	cols := pic.Bounds().W() / tileSize
	rows := pic.Bounds().H() / tileSize

	maxIndex := (rows * cols) - 1.0

	index := maxIndex
	id := maxIndex + 1
	spritesheet := map[int]*pixel.Sprite{}
	for row := (rows - 1); row >= 0; row-- {
		for col := (cols - 1); col >= 0; col-- {
			x := col
			y := math.Abs(rows-row) - 1
			spritesheet[int(id)] = pixel.NewSprite(pic, pixel.R(
				x*tileSize,
				y*tileSize,
				x*tileSize+tileSize,
				y*tileSize+tileSize,
			))
			index--
			id--
		}
	}
	return spritesheet
}
