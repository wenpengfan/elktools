package utils

import "github.com/desertbit/grumble"

func SortFlags(f *grumble.Flags) {
	f.BoolL("desc", false, "desc text")
}

func SortString(c *grumble.Context) string {
	if c.Flags.Bool("desc") {
		return "desc"
	}
	return "asc"
}
