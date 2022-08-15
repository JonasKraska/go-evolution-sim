package engine

import (
	"image"
)

type Drawer interface {
	Draw() image.Image
}