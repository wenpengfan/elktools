package elasticsearch

import (
	"elktools/cmd/utils"

	"github.com/desertbit/grumble"
)

func init() {
	Register("tasks", initPendingTasks)
}

func initPendingTasks(name string) {
	pendingTasksCommand := &grumble.Command{
		Name:      name,
		Help:      "GET _cat/pending_tasks",
		HelpGroup: defaultApp.App.Config().Name,
		Usage:     utils.Usage("elastic tasks"),
		Flags:     initFlags,
		AllowArgs: true,
		Run:       defaultApp.InitRun(pendingTasksRun),
	}
	defaultApp.App.AddCommand(pendingTasksCommand)
}

func pendingTasksRun(c *grumble.Context) ([]byte, error) {
	return defaultApp.search.GetPendingTasks(utils.SortString(c), c.Args)
}
