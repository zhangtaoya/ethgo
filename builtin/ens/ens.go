// Code generated by ethgo/abigen. DO NOT EDIT.
// Hash: bfee2618a5908e1a24f19dcce873d3b8e797374138dd7604f7b593db3cca5c17
// Version: 0.1.1
package ens

import (
	"fmt"
	"math/big"

	"github.com/umbracle/ethgo"
	"github.com/zhangtaoya/ethgo/contract"
	"github.com/zhangtaoya/ethgo/jsonrpc"
)

var (
	_ = big.NewInt
	_ = jsonrpc.NewClient
)

// ENS is a solidity contract
type ENS struct {
	c *contract.Contract
}

// DeployENS deploys a new ENS contract
func DeployENS(provider *jsonrpc.Client, from ethgo.Address, args []interface{}, opts ...contract.ContractOption) (contract.Txn, error) {
	return contract.DeployContract(abiENS, binENS, args, opts...)
}

// NewENS creates a new instance of the contract at a specific address
func NewENS(addr ethgo.Address, opts ...contract.ContractOption) *ENS {
	return &ENS{c: contract.NewContract(addr, abiENS, opts...)}
}

// calls

// Owner calls the owner method in the solidity contract
func (e *ENS) Owner(node [32]byte, block ...ethgo.BlockNumber) (retval0 ethgo.Address, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = e.c.Call("owner", ethgo.EncodeBlock(block...), node)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(ethgo.Address)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// Resolver calls the resolver method in the solidity contract
func (e *ENS) Resolver(node [32]byte, block ...ethgo.BlockNumber) (retval0 ethgo.Address, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = e.c.Call("resolver", ethgo.EncodeBlock(block...), node)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(ethgo.Address)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// Ttl calls the ttl method in the solidity contract
func (e *ENS) Ttl(node [32]byte, block ...ethgo.BlockNumber) (retval0 uint64, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = e.c.Call("ttl", ethgo.EncodeBlock(block...), node)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(uint64)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// txns

// SetOwner sends a setOwner transaction in the solidity contract
func (e *ENS) SetOwner(node [32]byte, owner ethgo.Address) (contract.Txn, error) {
	return e.c.Txn("setOwner", node, owner)
}

// SetResolver sends a setResolver transaction in the solidity contract
func (e *ENS) SetResolver(node [32]byte, resolver ethgo.Address) (contract.Txn, error) {
	return e.c.Txn("setResolver", node, resolver)
}

// SetSubnodeOwner sends a setSubnodeOwner transaction in the solidity contract
func (e *ENS) SetSubnodeOwner(node [32]byte, label [32]byte, owner ethgo.Address) (contract.Txn, error) {
	return e.c.Txn("setSubnodeOwner", node, label, owner)
}

// SetTTL sends a setTTL transaction in the solidity contract
func (e *ENS) SetTTL(node [32]byte, ttl uint64) (contract.Txn, error) {
	return e.c.Txn("setTTL", node, ttl)
}

// events

func (e *ENS) NewOwnerEventSig() ethgo.Hash {
	return e.c.GetABI().Events["NewOwner"].ID()
}

func (e *ENS) NewResolverEventSig() ethgo.Hash {
	return e.c.GetABI().Events["NewResolver"].ID()
}

func (e *ENS) NewTTLEventSig() ethgo.Hash {
	return e.c.GetABI().Events["NewTTL"].ID()
}

func (e *ENS) TransferEventSig() ethgo.Hash {
	return e.c.GetABI().Events["Transfer"].ID()
}
