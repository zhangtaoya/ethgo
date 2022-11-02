package testcases

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
	"github.com/umbracle/ethgo/testutil"
)

func TestContract_Signatures(t *testing.T) {
	var signatures []struct {
		Name      string `json:"name"`
		Signature string `json:"signature"`
		SigHash   string `json:"sigHash"`
		Abi       string `json:"abi"`
	}
	ReadTestCase(t, "contract-signatures", &signatures)

	for _, c := range signatures {
		t.Run(c.Name, func(t *testing.T) {
			m, err := abi.NewMethod(c.Signature)
			assert.NoError(t, err)

			sigHash := "0x" + hex.EncodeToString(m.ID())
			assert.Equal(t, sigHash, c.SigHash)
		})
	}
}

func TestContract_Interface(t *testing.T) {
	t.Skip()

	server := testutil.NewTestServer(t, nil)
	defer server.Close()

	var calls []struct {
		Name      string         `json:"name"`
		Interface string         `json:"interface"`
		Bytecode  ethgo.ArgBytes `json:"bytecode"`
		Result    ethgo.ArgBytes `json:"result"`
		Values    string         `json:"values"`
	}
	ReadTestCase(t, "contract-interface", &calls)

	for _, c := range calls {
		t.Run(c.Name, func(t *testing.T) {
			a, err := abi.NewABI(c.Interface)
			assert.NoError(t, err)

			method := a.GetMethod("test")

			receipt, err := server.SendTxn(&ethgo.Transaction{
				Input: c.Bytecode.Bytes(),
			})
			assert.NoError(t, err)

			outputRaw, err := server.Call(&ethgo.CallMsg{
				To:   &receipt.ContractAddress,
				Data: method.ID(),
			})
			assert.NoError(t, err)

			output, err := hex.DecodeString(outputRaw[2:])
			assert.NoError(t, err)

			_, err = method.Decode(output)
			assert.NoError(t, err)
		})
	}

}
