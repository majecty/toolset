package explorer

import (
	g "github.com/AllenDang/giu"
)

var (
	nodeInformation string
)

func DrawNodeWidgets() []g.Widget {
	return []g.Widget{
		g.Label("Node information"),
		g.InputTextMultiline(&nodeInformation).Size(-1, -1),
	}
}
