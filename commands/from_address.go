package commands

import (
	"gx/ipfs/QmVmDhyTTUcQXFD1rRQ64fGLMSAoaQvNH3hwuaCFAPq2hy/errors"
	cmdkit "gx/ipfs/QmceUdzxkimdYsgtX733uNgzf1DLHyBKN6ehGSp85ayppM/go-ipfs-cmdkit"

	"github.com/filecoin-project/go-filecoin/node"
	"github.com/filecoin-project/go-filecoin/types"
)

func fromAddress(opts cmdkit.OptMap, node *node.Node) (ret types.Address, err error) {
	o := opts["from"]
	if o != nil {
		ret, err = types.NewAddressFromString(o.(string))
		if err != nil {
			err = errors.Wrap(err, "invalid from address")
		}
	} else {
		ret, err = defaultWalletAddress(node)
		if err != nil || ret != (types.Address{}) {
			return
		}

		if len(node.Wallet.Addresses()) == 1 {
			ret = node.Wallet.Addresses()[0]
		} else {
			err = ErrCouldNotDefaultFromAddress
		}
	}
	return
}

func defaultWalletAddress(n *node.Node) (types.Address, error) {
	addr, err := n.Repo.Config().Get("wallet.defaultAddress")
	if err != nil {
		return types.Address{}, err
	}
	return addr.(types.Address), nil
}
