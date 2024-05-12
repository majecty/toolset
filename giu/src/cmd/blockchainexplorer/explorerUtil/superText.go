package explorerutil

import (
	"encoding/base64"
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

	var isBase64 bool
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(text)))
	n, err := base64.StdEncoding.Decode(base64Text, []byte(text))
	base64Text = base64Text[:n]
	if err != nil {
		isBase64 = false
	} else {
		isBase64 = true
		for _, c := range base64Text {
			if c < 32 || c > 126 {
				isBase64 = false
				break
			}
		}
	}

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

	if isHex && !isNumber {
		contextMenuItems = append(contextMenuItems,
			// TODO: 블록 hash로 조회가 잘 안됨
			// g.Selectable("Find block by hash").OnClick(func() {
			// 	AddBlockWindow(text)
			// }),
			g.Selectable("Find transaction by hash").OnClick(func() {
				SetGlobalError(fmt.Errorf("not implemented"))
			}),
		)
	}

	if isNumber {
		contextMenuItems = append(contextMenuItems,
			g.Selectable("Find block by height").OnClick(func() {
				AddBlockWindow(text)
			}),
		)
	}

	var widgets []g.Widget
	widgets = append(widgets,
		g.Selectable(text),
		g.ContextMenu().MouseButton(g.MouseButtonLeft).Layout(contextMenuItems...),
		// g.Label(fmt.Sprintf("isBase64 %v", isBase64)),
	)
	if isBase64 {
		widgets = append(widgets,
			g.Label(fmt.Sprintf("Base64: %s", base64Text)),
		)
	}

	return g.Layout(widgets)
}
