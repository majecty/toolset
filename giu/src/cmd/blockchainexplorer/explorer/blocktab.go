package explorer

import (
	"context"
	"fmt"
	explorerutil "toolset/giu/src/cmd/blockchainexplorer/explorerUtil"

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
			g.Button("Get Latest Block").OnClick(getLatestBlock),
		),
	}
}

func getLatestBlock() {
	httpClient, err := getTendermintHTTPClient()
	if err != nil {
		explorerutil.SetGlobalError(fmt.Errorf("failed to create tendermint client: %w", err))
		return
	}

	nodeStatusResult, err := httpClient.Status(context.Background())
	if err != nil {
		explorerutil.SetGlobalError(fmt.Errorf("failed to get node status: %w", err))
		return
	}

	blockNumberOrHash = nodeStatusResult.SyncInfo.LatestBlockHash.String()
}
