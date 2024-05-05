package explorerutil

import (
	"fmt"
	"strconv"

	g "github.com/AllenDang/giu"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
)

func DrawSuperTextWidget(text string) g.Widget {
	isHex := IsHex(text)

	var isNumber bool
	if _, err := strconv.Atoi(text); err != nil {
		isNumber = false
	} else {
		isNumber = true
	}

	var isAddress bool
	if _, err := cosmossdk.AccAddressFromBech32(text); err != nil {
		isAddress = false
	} else {
		isAddress = true
	}

	var widgets []g.Widget
	widgets = append(widgets,
		g.Selectable(text),
	)

	var contextMenuItems []g.Widget
	contextMenuItems = append(contextMenuItems,
		g.Selectable("Copy").OnClick(func() {
			g.Context.GetPlatform().SetClipboard(text)
		}),
	)

	if isAddress {
		contextMenuItems = append(contextMenuItems,
			g.Selectable("Find account by address").OnClick(func() {

				SetGlobalError(fmt.Errorf("not implemented"))
			}),
		)
	}

	if isHex {
		contextMenuItems = append(contextMenuItems,
			g.Selectable("Find block by hash").OnClick(func() {
				SetGlobalError(fmt.Errorf("not implemented"))
			}),
			g.Selectable("Find transaction by hash").OnClick(func() {
				SetGlobalError(fmt.Errorf("not implemented"))
			}),
		)
	}

	if isNumber {
		contextMenuItems = append(contextMenuItems,
			g.Selectable("Find block by height").OnClick(func() {
				SetGlobalError(fmt.Errorf("not implemented"))
			}),
		)
	}

	widgets = append(widgets, g.ContextMenu().MouseButton(g.MouseButtonLeft).Layout(contextMenuItems...))

	return g.Layout(widgets)
}
