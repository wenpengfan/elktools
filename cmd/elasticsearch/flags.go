package elasticsearch

import (
	"elktools/cmd/utils"
	"fmt"
	"time"

	"github.com/desertbit/grumble"
)

func AppFlags(f *grumble.Flags) {
	f.String("a", "address", defaultElasticURL, "set elastic search address")
	f.String("u", "username", "elastic", "set elastic search username")
	f.String("p", "password", "changeme", "set elastic search password")
	f.Duration("t", "timeout", 60*time.Second, "set connect elastic search timeout")
	f.Bool("d", "debug", false, "set debug")
}

func initFlags(f *grumble.Flags) {
	utils.PipelineFlags(f)
	utils.SortFlags(f)
	f.UintL("interval", 0, "the interval parameter specifies the amount of time in seconds between each report")
}

func indexFlags(f *grumble.Flags) {
	initFlags(f)
	f.StringL("health", "green", fmt.Sprintf("index --health %s", HealthStatusIndex))
}
