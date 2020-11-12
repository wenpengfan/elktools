package elasticsearch

import (
	"elktools/cmd/utils"

	"github.com/desertbit/grumble"
)

func init() {
	Register("health", initHealth)
}

func initHealth(name string) {
	healthCommand := &grumble.Command{
		Name:      name,
		Help:      "GET _cat/health?v",
		HelpGroup: defaultApp.App.Config().Name,
		Usage:     utils.Usage("elastic health"),
		Flags:     initFlags,
		AllowArgs: true,
		Run:       defaultApp.InitRun(healthRun),
	}
	defaultApp.App.AddCommand(healthCommand)
}
func healthRun(c *grumble.Context) ([]byte, error) {
	return defaultApp.search.GetHealth(c.Args)
}
