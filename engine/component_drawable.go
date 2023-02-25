package engine

import "github.com/hajimehoshi/ebiten/v2"

type Drawer interface {
    Draw() *ebiten.Image
}