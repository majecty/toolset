package main

import (
	_ "embed"

	g "github.com/AllenDang/giu"
	tendermintHTTP "github.com/tendermint/tendermint/rpc/client/http"
)

const (
	tabBlock      = 1
	tabTx         = 2
	publicNodeUrl = "https://rpc-arctic-1.sei-apis.com/"
)

var (
	nodeUrl     string = publicNodeUrl
	globalError error
)

func makeWidgets() []g.Widget {
	if globalError != nil {
		return []g.Widget{
			g.Label("Error: " + globalError.Error()),
			g.Button("Reset").OnClick(func() {
				globalError = nil
			}),
		}
	}

	widgets := []g.Widget{
		g.Label("Sei blockchain explorer"),
		g.Row(
			g.Label("Node URL:"),
			g.InputText(&nodeUrl),
		),
		g.Button("Reset to public node url").OnClick(func() {
			nodeUrl = publicNodeUrl
		}),
		g.Spacing(),
		g.TabBar().
			TabItems(
				g.TabItem("Node").
					Layout(g.Label("In the Node tab")),
				g.TabItem("Block").
					Layout(g.Label("In the Block tab")),
				g.TabItem("Tx").
					Layout(g.Label("In the Tx tab")),
				g.TabItem("Address").
					Layout(g.Label("In the Address tab")),
			),
		// g.InputText(&txHash),
		// g.Button("Get Random Recent Tx").OnClick(onClickGetRandomRecentTx),
		// g.Button("Get Tx information").OnClick(onClickGetTx),
		// g.Condition(rawTxInformation != "", makeJsonTreeWidget(), nil),
		// g.InputTextMultiline(&rawTxInformation).Size(-1, -1),
	}

	return []g.Widget{
		g.Style().SetFontSize(20).To(widgets...),
	}
}

func loop() {
	widgets := makeWidgets()

	g.SingleWindow().Layout(
		widgets...,
	)
}

func main() {
	wnd := g.NewMasterWindow("Block explorer", 600, 400, 0)
	g.Context.FontAtlas.SetDefaultFontSize(3)

	wnd.Run(loop)
}

func onClickGetTx() {
	// result, err := tendermintHTTPClient.Tx(context.Background(), bytes.HexBytes(txHash), false)
	// if err != nil {
	// 	rawTxInformation = fmt.Sprintf("Failed to get tx: %v", err)
	// 	fmt.Printf("Failed to get tx: %v\n", err)
	// 	return
	// }
	// rawTxInformation = fmt.Sprintf("Tx: %v", result)
	// fmt.Printf("Tx: %v\n", result)
	// fmt.Println("hi")
}

func onClickGetRandomRecentTx() {
	// blockheight, err := tendermintHTTPClient.Status(context.Background())
	// if err != nil {
	// 	fmt.Printf("failed to get latest block number: %w\n", err)
	// 	return
	// }

	// query := fmt.Sprintf("tx.height >= %d", blockheight.SyncInfo.LatestBlockHeight)
	// page := 1
	// perPage := 1
	// result, err := tendermintHTTPClient.TxSearch(context.Background(), query, false, &page, &perPage, "asc")
	// if err != nil {
	// 	rawTxInformation = fmt.Sprintf("Failed to get tx: %v", err)
	// 	fmt.Printf("Failed to get tx: %v\n", err)
	// 	return
	// }
	// txHash = string(result.Txs[0].Tx.String())
	// rawTxInformation = fmt.Sprintf("Tx: %v", result)

	// pp.Default.SetColoringEnabled(false)
	// // marsharl result to json and save it to rawTxInformation
	// jsonBytes, _ := json.MarshalIndent(result, "", "  ")
	// rawTxInformation = string(jsonBytes)

	// fmt.Printf("Tx: %v\n", result)
}

func getTendermintHTTPClient() (*tendermintHTTP.HTTP, error) {
	tendermintHTTPClient, err := tendermintHTTP.New(nodeUrl)
	if err != nil {
		return nil, err
	}

	return tendermintHTTPClient, nil
}
