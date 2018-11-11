package executor

import (
	"gitlab.33.cn/chain33/chain33/types"
	pty "gitlab.33.cn/chain33/plugin/plugin/dapp/norm/types"
)

func (n *Norm) Query_NormGet(in *pty.NormGetKey) (types.Message, error) {
	value, err := n.GetStateDB().Get(Key(in.Key))
	if err != nil {
		return nil, types.ErrNotFound
	}
	return &types.ReplyString{string(value)}, nil
}
