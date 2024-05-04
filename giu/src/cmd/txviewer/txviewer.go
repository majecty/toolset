package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"

	g "github.com/AllenDang/giu"
	pp "github.com/k0kubun/pp/v3"
	"github.com/tendermint/tendermint/libs/bytes"
	tendermintHTTP "github.com/tendermint/tendermint/rpc/client/http"
)

var (
	txHash               string
	rawTxInformation     string
	tendermintHTTPClient *tendermintHTTP.HTTP
	//go:embed txviewer.css
	cssStyle []byte
)

func init() {
	var err error
	tendermintHTTPClient, err = tendermintHTTP.New("https://rpc-arctic-1.sei-apis.com/")
	if err != nil {
		log.Fatalf("Failed to create tendermint client: %v\n", err)
	}
}

func makeWidgets() []g.Widget {
	widgets := []g.Widget{
		g.Label("Tx Viewer"),
		g.InputText(&txHash),
		g.Button("Get Random Recent Tx").OnClick(onClickGetRandomRecentTx),
		g.Button("Get Tx information").OnClick(onClickGetTx),

		g.InputTextMultiline(&rawTxInformation).Size(-1, -1),
	}

	return []g.Widget{
		g.Style().SetFontSize(20).To(widgets...),
	}
}

func loop() {
	widgets := makeWidgets()

	// g.setnext
	g.SingleWindow().Layout(
		widgets...,
	)
}

func main() {
	wnd := g.NewMasterWindow("Tx viewer", 600, 400, 0)
	g.Context.FontAtlas.SetDefaultFontSize(3)

	// if err := g.ParseCSSStyleSheet(cssStyle); err != nil {
	// 	panic(err)
	// }

	wnd.Run(loop)
}

func onClickGetTx() {
	result, err := tendermintHTTPClient.Tx(context.Background(), bytes.HexBytes(txHash), false)
	if err != nil {
		rawTxInformation = fmt.Sprintf("Failed to get tx: %v", err)
		fmt.Printf("Failed to get tx: %v\n", err)
		return
	}
	rawTxInformation = fmt.Sprintf("Tx: %v", result)
	fmt.Printf("Tx: %v\n", result)
	fmt.Println("hi")
}

func onClickGetRandomRecentTx() {
	blockheight, err := tendermintHTTPClient.Status(context.Background())
	if err != nil {
		fmt.Printf("failed to get latest block number: %w\n", err)
		return
	}

	query := fmt.Sprintf("tx.height >= %d", blockheight.SyncInfo.LatestBlockHeight)
	page := 1
	perPage := 1
	result, err := tendermintHTTPClient.TxSearch(context.Background(), query, false, &page, &perPage, "asc")
	if err != nil {
		rawTxInformation = fmt.Sprintf("Failed to get tx: %v", err)
		fmt.Printf("Failed to get tx: %v\n", err)
		return
	}
	txHash = string(result.Txs[0].Tx.String())
	rawTxInformation = fmt.Sprintf("Tx: %v", result)

	pp.Default.SetColoringEnabled(false)
	rawTxInformation = pp.Sprint(result)
	fmt.Printf("Tx: %v\n", result)
}
