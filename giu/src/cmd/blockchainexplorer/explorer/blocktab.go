package explorer

import (
	g "github.com/AllenDang/giu"
)

var (
	blockNumberOrHash string
)

func DrawBlockWidgets() []g.Widget {
	return []g.Widget{
		g.Row(
			g.Label("Block Number or Hash:"),
			g.InputText(&blockNumberOrHash),
		),
		g.Row(
			g.Button("Get Block").OnClick(func() {
			}),
			g.Button("Get Latest Block").OnClick(func() {
			}),
		),
	}
}
