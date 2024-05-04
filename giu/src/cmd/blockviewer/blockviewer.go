package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"strconv"

	pp "github.com/k0kubun/pp/v3"

	g "github.com/AllenDang/giu"
	tendermintHTTP "github.com/tendermint/tendermint/rpc/client/http"
)

var (
	blockHeight         string
	rawBlockInfromation string

	tendermintHTTPClient *tendermintHTTP.HTTP
	globalError          error
)

func init() {
	var err error
	tendermintHTTPClient, err = tendermintHTTP.New("https://rpc-arctic-1.sei-apis.com/")
	if err != nil {
		log.Fatalf("Failed to create tendermint client: %v\n", err)
	}
}

func makeWidgets() []g.Widget {

	if globalError != nil {
		return []g.Widget{
			g.Label("Error: " + globalError.Error()),
			g.Button("Clear Error").OnClick(func() {
				globalError = nil
			}),
		}
	}

	widgets := []g.Widget{
		g.Label("Block Viewer"),

		g.Row(
			g.Label("Block Height"),
			g.InputText(&blockHeight),
		),

		g.Button("Get Latest Block Number").OnClick(onClickGetLatestBlockNumber),
		g.Button("Get Block").OnClick(onClickGetBlock),
		g.Button("Get Block Result").OnClick(onClickGetBlockResult),

		g.InputTextMultiline(&rawBlockInfromation).Size(-1, -1),
	}

	return widgets
}

func biggerFont(widgets []g.Widget) []g.Widget {
	return []g.Widget{
		g.Style().SetFontSize(20).To(widgets...),
	}
}

func loop() {
	widgets := biggerFont(makeWidgets())

	// g.setnext
	g.SingleWindow().Layout(
		widgets...,
	)
}

func main() {
	wnd := g.NewMasterWindow("Block viewer", 600, 400, 0)
	g.Context.FontAtlas.SetDefaultFontSize(3)

	wnd.Run(loop)
}

func onClickGetBlock() {
	blockHeightInt, err := strconv.ParseInt(blockHeight, 10, 64)
	if err != nil {
		globalError = fmt.Errorf("failed to parse block height: %w", err)
		return
	}
	result, err := tendermintHTTPClient.Block(context.Background(), &blockHeightInt)

	if err != nil {
		globalError = fmt.Errorf("failed to get block information: %w", err)
		return
	}

	rawBlockInfromation = fmt.Sprintf("Block: %v", result)
}

func onClickGetLatestBlockNumber() {
	result, err := tendermintHTTPClient.Status(context.Background())
	if err != nil {
		globalError = fmt.Errorf("failed to get latest block number: %w", err)
		return
	}

	blockHeight = strconv.FormatInt(result.SyncInfo.LatestBlockHeight, 10)
}

func onClickGetBlockResult() {
	blockHeightInt, err := strconv.ParseInt(blockHeight, 10, 64)
	if err != nil {
		globalError = fmt.Errorf("failed to parse block height: %w", err)
		return
	}
	result, err := tendermintHTTPClient.BlockResults(context.Background(), &blockHeightInt)

	if err != nil {
		globalError = fmt.Errorf("failed to get block result: %w", err)
		return
	}

	rawBlockInfromation = fmt.Sprintf("Block Result: %v", result)
	rawBlockInfromation = fmt.Sprintf("Block Result: %v", result)
	pp.Default.SetColoringEnabled(false)
	rawBlockInfromation = pp.Sprintln(result)
}
