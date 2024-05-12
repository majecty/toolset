package explorerutil

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	g "github.com/AllenDang/giu"
	"github.com/tendermint/tendermint/libs/bytes"
	"github.com/tendermint/tendermint/rpc/coretypes"
)

var windows []BlockWindow

type BlockWindow struct {
	blockNumberOrHash string
	BlockResult       string
}

func AddBlockWindow(blockNumberOrHash string) {
	for i := range windows {
		if windows[i].blockNumberOrHash == blockNumberOrHash {
			return
		}
	}

	windows = append(windows, BlockWindow{
		blockNumberOrHash: blockNumberOrHash,
	})
}

func DummyBlockWindows() {
	windows = append(windows, BlockWindow{
		blockNumberOrHash: "18937999",
		BlockResult:       "",
	})
	windows = append(windows, BlockWindow{
		blockNumberOrHash: "18938999",
		BlockResult:       "",
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
		gWindow.Size(400, 400)
		gWindow.Layout(
			g.Style().SetFontSize(20).To(
				window.drawBlockindow()...,
			),
		)
	}
}

func (bw *BlockWindow) drawBlockindow() []g.Widget {
	return []g.Widget{
		g.Label("Block Number or Hash"),
		g.InputText(&bw.blockNumberOrHash),
		g.Button("Close").OnClick(func() {
			var myWindowIndex int
			for i := range windows {
				if windows[i].blockNumberOrHash == bw.blockNumberOrHash {
					myWindowIndex = i
					break
				}
			}

			windows = append(windows[:myWindowIndex], windows[myWindowIndex+1:]...)
		}),
		g.Button("Get Block").OnClick(func() {
			bw.getBlock()
		}),
		g.Condition(bw.BlockResult != "", DrawJsonTreeWidget(bw.BlockResult), nil),
	}
}

func (bw *BlockWindow) getBlock() {
	blockHeightInt, err := strconv.ParseInt(bw.blockNumberOrHash, 10, 64)
	useBlockHeight := err == nil

	tendermintClient, err := GetTendermintHTTPClient()
	if err != nil {
		SetGlobalError(fmt.Errorf("failed to create tendermint client: %w", err))
		return
	}

	var result *coretypes.ResultBlock
	if useBlockHeight {
		result, err = tendermintClient.Block(context.Background(), &blockHeightInt)
	} else {
		result, err = tendermintClient.BlockByHash(context.Background(), bytes.HexBytes(bw.blockNumberOrHash))
	}

	if err != nil {
		SetGlobalError(fmt.Errorf("failed to get block information: %w", err))
		return
	}

	blockJson, err := json.Marshal(result)
	if err != nil {
		SetGlobalError(fmt.Errorf("failed to marshal block information: %w", err))
		return
	}

	bw.BlockResult = string(blockJson)
}
