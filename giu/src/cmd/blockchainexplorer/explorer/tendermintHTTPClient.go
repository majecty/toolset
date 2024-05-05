package explorer

import (
	tendermintHTTP "github.com/tendermint/tendermint/rpc/client/http"
)

const (
	publicNodeUrl = "https://rpc-arctic-1.sei-apis.com/"
)

var (
	NodeUrl string = publicNodeUrl
)

func getTendermintHTTPClient() (*tendermintHTTP.HTTP, error) {
	tendermintHTTPClient, err := tendermintHTTP.New(NodeUrl)
	if err != nil {
		return nil, err
	}

	return tendermintHTTPClient, nil
}

func ResetNodeUrl() {
	NodeUrl = publicNodeUrl
}
