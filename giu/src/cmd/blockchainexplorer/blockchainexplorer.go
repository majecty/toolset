package main

import (
	_ "embed"

	explorer "toolset/giu/src/cmd/blockchainexplorer/explorer"
	explorerutil "toolset/giu/src/cmd/blockchainexplorer/explorerUtil"

	g "github.com/AllenDang/giu"
)

const (
	tabBlock = 1
	tabTx    = 2
)

func makeWidgets() []g.Widget {
	if explorerutil.GlobalError != nil {
		return explorerutil.DrawGlobalErrorWidgets()
	}

	widgets := []g.Widget{
		g.Label("Sei blockchain explorer"),
		g.Row(
			g.Label("Node URL:"),
			g.InputText(&explorer.NodeUrl),
		),
		g.Button("Reset to public node url").OnClick(func() {
			explorer.ResetNodeUrl()
		}),
		g.Spacing(),
		g.Separator(),
		g.TabBar().
			TabItems(
				g.TabItem("Node").
					Layout(
						explorer.DrawNodeWidgets()...,
					),
				g.TabItem("Block").
					Layout(
						explorer.DrawBlockWidgets()...,
					),
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

	homeWindow := g.Window("Home")
	homeWindow.Size(600, 400)

	homeWindow.Layout(
		widgets...,
	)
	explorerutil.DrawBlockWindows()
}

func main() {
	explorerutil.InitSeiCosmos()
	wnd := g.NewMasterWindow("Block explorer", 800, 800, 0)
	g.Context.FontAtlas.SetDefaultFontSize(3)

	explorerutil.DummyBlockWindows()

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
