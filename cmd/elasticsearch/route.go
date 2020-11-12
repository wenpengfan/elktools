package elasticsearch

import (
	"elktools/cmd/utils"
	"encoding/json"
	"fmt"

	"github.com/desertbit/grumble"
)

func init() {
	Register("route", initRoot)
}

func initRoot(name string) {
	routeCommand := &grumble.Command{
		Name:      name,
		Help:      "POST /_cluster/reroute",
		HelpGroup: defaultApp.App.Config().Name,
		Usage:     utils.Usage("elastic route"),
	}
	defaultApp.App.AddCommand(routeCommand)

	routeCommand.AddCommand(&grumble.Command{
		Name:      "retry",
		Help:      "will attempt a single retry round for these shards",
		HelpGroup: routeCommand.Name,
		Usage:     utils.Usage("elastic route retry"),
		Flags:     initFlags,
		AllowArgs: false,
		Run:       defaultApp.InitRun(retryFailedRun),
	})
}

func retryFailedRun(c *grumble.Context) ([]byte, error) {
	data, err := defaultApp.search.RetryFailedRoute(c.Args)
	if err != nil {
		return nil, err
	}
	result := struct {
		Acknowledged bool        `json:"acknowledged"`
		Error        interface{} `json:"error"`
		Status       int         `json:"status"`
	}{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	if result.Status > 0 {
		return nil, fmt.Errorf("%s\n", data)
	}
	return []byte(fmt.Sprintf("%+v\n", result)), nil
}
