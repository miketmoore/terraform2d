package terraform2d

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/deanobob/tmxreader"
	"github.com/faiface/pixel"
)

// Load loads the tmx files
func Load(tilemapFiles []string, tilemapDir string) map[string]tmxreader.TmxMap {
	tmxMapData := map[string]tmxreader.TmxMap{}
	for _, name := range tilemapFiles {
		path := fmt.Sprintf("%s%s.tmx", tilemapDir, name)
		tmxMapData[name] = parseTmxFile(path)
	}
	return tmxMapData
}

func parseTmxFile(filename string) tmxreader.TmxMap {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	tmxMap, err := tmxreader.Parse(raw)
	if err != nil {
		panic(err)
	}

	return tmxMap
}

type mapDrawData struct {
	Rect     pixel.Rect
	SpriteID int
}

// MapData represents data for one map
type MapData struct {
	Name string
	Data []mapDrawData
}

// BuildMapDrawData builds draw data and stores it in a map
func BuildMapDrawData(dir string, files []string, tileSize float64) map[string]MapData {

	// load all TMX file data for each map
	tmxMapData := Load(files, dir)

	all := map[string]MapData{}

	for mapName, mapData := range tmxMapData {

		md := MapData{
			Name: mapName,
			Data: []mapDrawData{},
		}

		layers := mapData.Layers
		for _, layer := range layers {

			records := ParseCSV(strings.TrimSpace(layer.Data.Value) + ",")
			for row := 0; row <= len(records); row++ {
				if len(records) > row {
					for col := 0; col < len(records[row])-1; col++ {
						y := float64(11-row) * tileSize
						x := float64(col) * tileSize

						record := records[row][col]
						spriteID, err := strconv.Atoi(record)
						if err != nil {
							panic(err)
						}
						mrd := mapDrawData{
							Rect:     pixel.R(x, y, x+tileSize, y+tileSize),
							SpriteID: spriteID,
						}
						md.Data = append(md.Data, mrd)
					}
				}

			}
			all[mapName] = md
		}
	}

	return all
}
