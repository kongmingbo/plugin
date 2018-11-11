package token

import (
	"gitlab.33.cn/chain33/chain33/pluginmgr"
	"gitlab.33.cn/chain33/plugin/plugin/dapp/token/commands"
	"gitlab.33.cn/chain33/plugin/plugin/dapp/token/executor"
	"gitlab.33.cn/chain33/plugin/plugin/dapp/token/rpc"
	"gitlab.33.cn/chain33/plugin/plugin/dapp/token/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.TokenX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.TokenCmd,
		RPC:      rpc.Init,
	})
}
