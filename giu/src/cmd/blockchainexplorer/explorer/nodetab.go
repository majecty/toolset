package explorer

import (
	"context"
	"encoding/json"
	"fmt"
	explorerutil "toolset/giu/src/cmd/blockchainexplorer/explorerUtil"

	g "github.com/AllenDang/giu"
)

var (
	nodeStatus string
)

func DrawNodeWidgets() []g.Widget {
	return []g.Widget{
		g.Button("Get Node information").OnClick(func() {
			nodeStatus = ""
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

			nodeStatusJson, err := json.Marshal(nodeStatusResult)
			if err != nil {
				explorerutil.SetGlobalError(fmt.Errorf("failed to marshal node status: %w", err))
				return
			}

			nodeStatus = string(nodeStatusJson)
		}),
		g.Condition(nodeStatus != "", explorerutil.DrawJsonTreeWidget(nodeStatus), nil),
		// g.InputTextMultiline(&nodeStatus).Size(-1, -1),
	}
}
