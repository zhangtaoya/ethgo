package ens

import "github.com/umbracle/ethgo"

var defaultEnsAddr = ethgo.HexToAddress("0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e")

var builtinEnsAddr = map[uint64]ethgo.Address{
	1: defaultEnsAddr,
	3: defaultEnsAddr,
	4: defaultEnsAddr,
	5: defaultEnsAddr,
}
