package ticket

import (
	"github.com/33cn/chain33/pluginmgr"
	"github.com/33cn/plugin/plugin/dapp/ticket/commands"
	"github.com/33cn/plugin/plugin/dapp/ticket/executor"
	"github.com/33cn/plugin/plugin/dapp/ticket/rpc"
	"github.com/33cn/plugin/plugin/dapp/ticket/types"
	_ "github.com/33cn/plugin/plugin/dapp/ticket/wallet"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.TicketX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.TicketCmd,
		RPC:      rpc.Init,
	})
}