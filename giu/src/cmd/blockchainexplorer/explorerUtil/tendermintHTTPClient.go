package explorerutil

import (
	"fmt"

	tendermintHTTP "github.com/tendermint/tendermint/rpc/client/http"
)

const (
	publicNodeUrl = "https://rpc-arctic-1.sei-apis.com/"
)

var (
	NodeUrl string = publicNodeUrl
)

func GetTendermintHTTPClient() (*tendermintHTTP.HTTP, error) {
	tendermintHTTPClient, err := tendermintHTTP.New(NodeUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create tendermint client: %w", err)
	}

	return tendermintHTTPClient, nil
}

func ResetNodeUrl() {
	NodeUrl = publicNodeUrl
}
