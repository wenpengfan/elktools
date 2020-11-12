package elasticsearch

import (
	"elktools/cmd/utils"

	"github.com/desertbit/grumble"
)

func init() {
	Register("info", initInfo)
}

func initInfo(name string) {
	infoCommand := &grumble.Command{
		Name:      name,
		Help:      "show elastic info",
		HelpGroup: defaultApp.App.Config().Name,
		Usage:     utils.Usage("elastic info"),
		AllowArgs: false,
		Run: func(c *grumble.Context) error {
			c.App.Printf("URL: %s\nUsername: %s\nPassword: %s\nTimeout: %v\n",
				defaultApp.search.Req.URL,
				defaultApp.search.Req.Username,
				defaultApp.search.Req.Password,
				defaultApp.search.Req.Timeout,
			)
			return nil
		},
	}
	defaultApp.App.AddCommand(infoCommand)
}
