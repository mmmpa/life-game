package main

import (
	"fmt"
	"github.com/mmmpa/life-game/life-game"
	"log"
	"encoding/json"
	"io/ioutil"
	"os"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"flag"
)

func main() {
	fmt.Print("life game.\n")

	file := flag.String("file", "sample/pentadecathlon.json", "")
	mode := flag.String("mode", "json", "")
	wait := flag.Int("wait", 100, "")

	flag.Parse()

	var c Configuration

	if *mode == "json" {
		c = toConfigurationFromJSON(*file)
	} else {
		c = toConfigurationFromImage(*file)
	}

	life_game.Run(*wait, c.Field.W, c.Field.H, c.Lives)
}

type Configuration struct {
	Field struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"field"`
	Lives []life_game.Position `json:"lives"`
}

func toConfigurationFromJSON(path string) Configuration {
	var err error

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var c Configuration
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func toConfigurationFromImage(path string) Configuration {
	var err error

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	bound := img.Bounds()

	c := Configuration{}

	c.Field.W = bound.Max.X
	c.Field.H = bound.Max.Y

	for y := 0; y < bound.Max.Y; y++ {
		for x := 0; x < bound.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			if(r + g + b == 0) {
				c.Lives = append(c.Lives, life_game.Position{x, y})
			}
		}
	}

	return c
}
