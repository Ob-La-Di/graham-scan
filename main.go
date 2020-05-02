package main

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"graham-scan/game"
	"image/color"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const (
	WIDTH  = 720
	HEIGHT = 480
)

var g *game.Game

func update(screen *ebiten.Image) error {
	//
	//for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
	//	if ebiten.IsKeyPressed(k) {
	//		g.SortByPivotAngle()
	//	}
	//}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	pixels := make([]byte, WIDTH*HEIGHT*4)

	for _, gift := range g.Gifts {
		if gift.IsSelected {
			pixels[(int(gift.Point.Y)*WIDTH+int(gift.Point.X))*4] = 0x00
			pixels[(int(gift.Point.Y)*WIDTH+int(gift.Point.X))*4+2] = 0x00
		} else {
			pixels[(int(gift.Point.Y)*WIDTH+int(gift.Point.X))*4] = 0xFF
			pixels[(int(gift.Point.Y)*WIDTH+int(gift.Point.X))*4+2] = 0xFF
		}
		pixels[(int(gift.Point.Y)*WIDTH+int(gift.Point.X))*4+1] = 0xFF
		pixels[(int(gift.Point.Y)*WIDTH+int(gift.Point.X))*4+3] = 0xFF
	}
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFont := truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	screen.ReplacePixels(pixels)

	for index, gift := range g.Gifts {
		if gift.IsSelected == true {
			text.Draw(screen, strconv.Itoa(index+1), mplusNormalFont, int(gift.Point.X), int(gift.Point.Y)+20, color.White)
		}
	}

	return nil
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	g = game.New(WIDTH, HEIGHT, 20)
	g.Wrap()
	if err := ebiten.Run(update, WIDTH, HEIGHT, 1, "Hello world!"); err != nil {
		panic(err)
	}
}
