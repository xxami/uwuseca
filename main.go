
package main

import (
  "fmt"
  _ "image/png"

  "github.com/hajimehoshi/ebiten"
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
)

func init() {
  img, _, err := ebitenutil.NewImageFromFile("res/note.png", ebiten.FilterDefault)
  if err != nil {
    panic(err)
  }
  image = img
}

func update(screen *ebiten.Image) error {
  //ebitenutil.DebugPrint(screen, fmt.Sprintf("%0.2f", ebiten.CurrentFPS()))
  screen.DrawImage(image, image_options)
  return nil
}

func main() {
  ebiten.Run(update, screen_width, screen_height, screen_scale, window_title)
}

