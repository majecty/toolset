package explorerutil

import (
	"fmt"

	g "github.com/AllenDang/giu"
)

var windows []BlockWindow

type BlockWindow struct {
	blockNumberOrHash string
	Layout            g.Layout
	BlockResult       string
}

func DummyBlockWindows() {
	windows = append(windows, BlockWindow{
		blockNumberOrHash: "0",
		Layout: g.Layout{
			g.Label("Block Number or Hash:"),
		},
		BlockResult: "",
	})
	windows = append(windows, BlockWindow{
		blockNumberOrHash: "1",
		Layout: g.Layout{
			g.Label("Block Number or Hash:"),
		},
		BlockResult: "",
	})
}

func (bw *BlockWindow) Title() string {
	return fmt.Sprintf("Block %s", bw.blockNumberOrHash)
}

func DrawBlockWindows() {
	for i := range windows {
		window := &windows[i]
		gWindow := g.Window(window.Title())
		gWindow.Pos(float32(500+30*(i+1)), float32(30*(i+1)))
		gWindow.Layout(window.Layout)
	}
}
