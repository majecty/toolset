package main

import (
	"errors"
	"fmt"

	g "github.com/AllenDang/giu"
	"github.com/btcsuite/btcd/btcec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

const (
	HumanCoinUnit = "sei"
	BaseCoinUnit  = "usei"
	UseiExponent  = 6

	DefaultBondDenom = BaseCoinUnit

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address.
	Bech32PrefixAccAddr = "sei"
)

var (
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key.
	Bech32PrefixAccPub = Bech32PrefixAccAddr + "pub"
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address.
	Bech32PrefixValAddr = Bech32PrefixAccAddr + "valoper"
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key.
	Bech32PrefixValPub = Bech32PrefixAccAddr + "valoperpub"
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address.
	Bech32PrefixConsAddr = Bech32PrefixAccAddr + "valcons"
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key.
	Bech32PrefixConsPub = Bech32PrefixAccAddr + "valconspub"
)

type GeneratedMnemonic struct {
	mnemonic   string
	path       string
	publicKey  string
	privateKey string
	seiAddress string
	// evmAddress string
}

var globalError error
var globalGeneratedMnemonic *GeneratedMnemonic

func onClickGenerateMnemonic() {
	globalError = nil
	cstore := keyring.NewInMemory()
	sec := hd.Secp256k1
	// Add keys and see they return in alphabetical order
	bob, mnemonic, err := cstore.NewMnemonic("Bob", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, sec)
	if err != nil {
		// this should never happen
		fmt.Println("Error generating mnemonic: " + err.Error())
		globalError = err
		return
	}

	unsafecstore := keyring.NewUnsafe(cstore)
	privateKey, err := unsafecstore.UnsafeExportPrivKeyHex("Bob")
	// privatekeyRaw = hex.DecodeString(privateKey)
	if err != nil {
		fmt.Println("Error exporting private key: " + err.Error())
		globalError = err
		return
	}
	// _, pubkeyObject := btcsecp256k1.PrivKeyFromBytes(btcsecp256k1.S256(), privateKey)
	// pk := pubkeyObject.SerializeCompressed()
	// return &PubKey{Key: pk}

	// evmAddress, err := pubkeyToEVMAddress(bob.GetPubKey().)
	// if err != nil {
	// 	fmt.Println("Error converting public key to EVM address: " + err.Error())
	// 	globalError = err;
	// 	return;
	// }

	globalGeneratedMnemonic = &GeneratedMnemonic{
		mnemonic:   mnemonic,
		path:       sdk.FullFundraiserPath,
		publicKey:  bob.GetPubKey().String(),
		privateKey: privateKey,
		seiAddress: bob.GetAddress().String(),
		// evmAddress: evmAddress.Hex(),
	}

	fmt.Println(bob.GetName())
	fmt.Println(bob.GetAddress())
	fmt.Println("Menomic", mnemonic)
	fmt.Println("Public Key", bob.GetPubKey().String())
	fmt.Println("Private Key", privateKey)
}

func makeWidgets() []g.Widget {
	widgets := []g.Widget{
		g.Button("Generate Mnemonic").OnClick(onClickGenerateMnemonic),
	}
	if globalError != nil {
		errorWidgets := []g.Widget{
			g.Label("Error generating mnemonic: " + globalError.Error()),
		}
		widgets = append(widgets, errorWidgets...)
		return widgets
	}

	if globalGeneratedMnemonic != nil {
		mnemonicWidgets := []g.Widget{
			g.Row(
				g.Label("Mnemonic: "),
				g.InputText(&globalGeneratedMnemonic.mnemonic).Flags(g.InputTextFlagsReadOnly),
			),
			g.Spacing(),
			g.Label("Path: " + globalGeneratedMnemonic.path).Wrapped(true),
			g.Row(
				g.Label("Public Key: "),
				g.InputText(&globalGeneratedMnemonic.publicKey).Flags(g.InputTextFlagsReadOnly),
			),
			g.Row(g.Label("Private Key: "), g.InputText(&globalGeneratedMnemonic.privateKey).Flags(g.InputTextFlagsReadOnly)),
			g.Row(g.Label("SEI Address: "), g.InputText(&globalGeneratedMnemonic.seiAddress).Flags(g.InputTextFlagsReadOnly)),
			// g.Row(g.Label("EVM Address: "), g.InputText(&globalGeneratedMnemonic.evmAddress).Flags(g.InputTextFlagsReadOnly)),
		}
		widgets = append(widgets, mnemonicWidgets...)
	}
	return widgets
}

func loop() {
	widgets := makeWidgets()

	// g.setnext
	g.SingleWindow().Layout(
		widgets...,
	)
}

func main() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
	wnd := g.NewMasterWindow("Mnemonic Generator", 600, 400, 0)
	wnd.Run(loop)
}

// https://github.com/sei-protocol/sei-chain/blob/seiv2/x/evm/ante/preprocess.go#L288
// second half of go-ethereum/core/types/transaction_signing.go:recoverPlain
func pubkeyToEVMAddress(pub []byte) (common.Address, error) {
	if len(pub) == 0 || pub[0] != 4 {
		return common.Address{}, errors.New("invalid public key")
	}
	var addr common.Address
	copy(addr[:], crypto.Keccak256(pub[1:])[12:])
	return addr, nil
}

// https://github.com/sei-protocol/sei-chain/blob/seiv2/x/evm/ante/preprocess.go#L296C1-L301C2
func pubkeyBytesToSeiPubKey(pub []byte) secp256k1.PubKey {
	pubKey, _ := crypto.UnmarshalPubkey(pub)
	pubkeyObj := (*btcec.PublicKey)(pubKey)
	return secp256k1.PubKey{Key: pubkeyObj.SerializeCompressed()}
}
