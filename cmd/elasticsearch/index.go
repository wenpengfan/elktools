package elasticsearch

import (
	"elktools/cmd/utils"
	"fmt"
	"strings"

	"github.com/desertbit/grumble"
)

var (
	HealthStatusIndex = []string{"green", "yellow", "red"}
)

func init() {
	Register("index", initIndex)
}

func initIndex(name string) {
	indexCommand := &grumble.Command{
		Name:      name,
		Help:      "GET _cat/indices/*-2020.01.01?v",
		HelpGroup: defaultApp.App.Config().Name,
		Usage:     utils.Usage("elastic index"),
		Flags: func(f *grumble.Flags) {
			indexFlags(f)
			f.StringL("day", "0", "index --day [-1 ort 2020.01.01]")
		},
		AllowArgs: true,
		Run:       defaultApp.InitRun(getIndexDayRun),
	}
	defaultApp.App.AddCommand(indexCommand)

	indexAllCommand := &grumble.Command{
		Name:      "all",
		Help:      "GET _cat/indices?v",
		HelpGroup: indexCommand.Name,
		Usage:     utils.Usage("elastic index all"),
		Flags:     indexFlags,
		AllowArgs: false,
		Run:       defaultApp.InitRun(getIndexAllRun),
	}
	indexCommand.AddCommand(indexAllCommand)
}

func getIndexDayRun(c *grumble.Context) ([]byte, error) {
	day := c.Flags.String("day")
	if err := parseHealth(c); err != nil {
		return nil, err
	}
	return defaultApp.search.GetIndexDay(day, utils.SortString(c), c.Args)
}

func getIndexAllRun(c *grumble.Context) ([]byte, error) {
	if err := parseHealth(c); err != nil {
		return nil, err
	}
	return defaultApp.search.GetIndexAll(utils.SortString(c), c.Args)
}

func parseHealth(c *grumble.Context) error {
	status := c.Flags.String("health")
	if !hasHealthStatusIndex(status) {
		return fmt.Errorf("elastic index [all] --health %v", HealthStatusIndex)
	}
	arg := fmt.Sprintf("health=%s", status)
	c.Args = append(c.Args, arg)
	return nil
}

func hasHealthStatusIndex(status string) bool {
	status = strings.TrimSpace(status)
	for _, s := range HealthStatusIndex {
		if strings.EqualFold(s, status) {
			return true
		}
	}
	return false
}
