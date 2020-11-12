package elasticsearch

import (
	"elktools/cmd/utils"

	"github.com/desertbit/grumble"
)

func init() {
	Register("shards", initShards)
}

func initShards(name string) {
	shardsCommand := &grumble.Command{
		Name:      name,
		Help:      "GET _cat/shards?v",
		HelpGroup: defaultApp.App.Config().Name,
		Usage:     utils.Usage("shards"),
		Flags:     initFlags,
		AllowArgs: true,
		Run:       defaultApp.InitRun(shardsRun),
	}
	defaultApp.App.AddCommand(shardsCommand)
}

func shardsRun(c *grumble.Context) ([]byte, error) {
	return defaultApp.search.GetShards(utils.SortString(c), c.Args)
}
