package token

import (
	"github.com/33cn/chain33/pluginmgr"
	"github.com/33cn/plugin/plugin/dapp/token/commands"
	"github.com/33cn/plugin/plugin/dapp/token/executor"
	"github.com/33cn/plugin/plugin/dapp/token/rpc"
	"github.com/33cn/plugin/plugin/dapp/token/types"
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