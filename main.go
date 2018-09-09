
package main

import (
  "fmt"
  "io/ioutil"
  "image/color"
  _ "image/png"

  "github.com/golang/freetype/truetype"
  "golang.org/x/image/font"
  "github.com/hajimehoshi/ebiten"
  "github.com/hajimehoshi/ebiten/text"
  "github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
  screen_width = 1024
  screen_height = 720
  screen_scale = 1
  window_title = "uwuseca"
)

var (
  image *ebiten.Image
  image_options = &ebiten.DrawImageOptions{}
  debug_font font.Face
)

func init() {
  img, _, err := ebitenutil.NewImageFromFile("res/note.png", ebiten.FilterDefault)
  if err != nil {
    panic(err)
  }
  image = img

  font_file, err := ebitenutil.OpenFile("res/NotoSans-Regular.ttf")
  if err != nil {
    panic(err)
  }
  defer font_file.Close()
  font_contents, err := ioutil.ReadAll(font_file)
  if err != nil {
    panic(err)
  }
  tt, err := truetype.Parse(font_contents)
  if err != nil {
    panic(err)
  }
  debug_font = truetype.NewFace(tt, &truetype.Options{
    Size: 32,
    DPI: 72,
    Hinting: font.HintingFull,
  })
}

func update(screen *ebiten.Image) error {
  screen.DrawImage(image, image_options)
  ebitenutil.DebugPrint(screen, fmt.Sprintf("%0.2f", ebiten.CurrentFPS()))
  text.Draw(screen, "Hello, World!", debug_font, 20, 40, color.White)
  return nil
}

func main() {
  ebiten.Run(update, screen_width, screen_height, screen_scale, window_title)
}

