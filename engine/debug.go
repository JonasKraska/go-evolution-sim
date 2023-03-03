package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var debugOutputs []string

func DebugReset() {
	debugOutputs = make([]string, 0)
}

func DebugPrintln(output string) {
	debugOutputs = append(debugOutputs, output)
}

func DebugPrint(screen *ebiten.Image) {
	var output string

	for _, o := range debugOutputs {
		output = output + o + "\n"
	}

	ebitenutil.DebugPrint(screen, output)
}
