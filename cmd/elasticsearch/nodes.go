package elasticsearch

import (
	"elktools/cmd/utils"

	"github.com/desertbit/grumble"
)

func init() {
	Register("nodes", initNodes)
}

func initNodes(name string) {
	nodesCommand := &grumble.Command{
		Name:      name,
		Help:      "GET _cat/nodes?v",
		HelpGroup: defaultApp.App.Config().Name,
		Usage:     utils.Usage("nodes"),
		Flags:     initFlags,
		AllowArgs: true,
		Run:       defaultApp.InitRun(nodesRun),
	}
	defaultApp.App.AddCommand(nodesCommand)
}

func nodesRun(c *grumble.Context) ([]byte, error) {
	return defaultApp.search.GetNodes(utils.SortString(c), c.Args)
}
