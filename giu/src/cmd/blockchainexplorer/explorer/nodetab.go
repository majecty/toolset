package explorer

import (
	"context"
	"encoding/json"
	explorerutil "toolset/giu/src/cmd/blockchainexplorer/explorerUtil"

	g "github.com/AllenDang/giu"
)

var (
	nodeStatus string
)

func DrawNodeWidgets() []g.Widget {
	return []g.Widget{
		g.Label("Node information"),
		g.Button("Get Node information").OnClick(func() {
			nodeStatus = ""
			httpClient, err := getTendermintHTTPClient()
			if err != nil {
				explorerutil.SetGlobalError(err)
				return
			}

			nodeStatusResult, err := httpClient.Status(context.Background())
			if err != nil {
				explorerutil.SetGlobalError(err)
				return
			}

			nodeStatusJson, err := json.Marshal(nodeStatusResult)
			if err != nil {
				explorerutil.SetGlobalError(err)
				return
			}

			nodeStatus = string(nodeStatusJson)
		}),
		g.InputTextMultiline(&nodeStatus).Size(-1, -1),
	}
}
