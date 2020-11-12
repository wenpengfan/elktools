package cmd

import (
	"elktools/cmd/elasticsearch"
	"elktools/cmd/utils"
	"errors"
	"os"
	"time"

	"github.com/desertbit/grumble"
	"github.com/fatih/color"
)

const version = "0.0.2"

var App = app

var app = grumble.New(&grumble.Config{
	Name:              "ELKTools",
	Description:       "ELK tools",
	PromptColor:       color.New(color.FgGreen, color.Bold),
	HelpHeadlineColor: color.New(color.FgGreen),
})

var (
	ErrTimeout     = errors.New("timeout")
	defaultTimeout = 30 * time.Minute
)

func init() {
	var err error
	app.OnInit(func(a *grumble.App, flags grumble.FlagMap) error {
		if !a.IsShell() {
			return nil
		}
		utils.NewTimeout(defaultTimeout, func() {
			err = ErrTimeout
			a.Close()
		})
		return nil
	})

	app.OnClose(func() error {
		if app.IsShell() {
			if err == ErrTimeout {
				app.Printf("\nBye, error: timeout %v\n", defaultTimeout)
				os.Exit(2)
			} else {
				app.Println("Bye")
			}
		}
		os.Exit(0)
		return nil
	})
	app.SetPrintASCIILogo(func(a *grumble.App) {
		a.Println("  ______ _      _  __  _______          _      ")
		a.Println(" |  ____| |    | |/ / |__   __|        | |    ")
		a.Println(" | |__  | |    | ' /     | | ___   ___ | |___ ")
		a.Println(" |  __| | |    |  <      | |/ _ \\ / _ \\| / __|")
		a.Println(" | |____| |____| . \\     | | (_) | (_) | \\__ \\")
		a.Println(" |______|______|_|\\_\\    |_|\\___/ \\___/|_|___/")
		a.Println("                                              ")
	})

	app.AddCommand(&grumble.Command{
		Name:      "version",
		Help:      "show version number and quit",
		HelpGroup: app.Config().Name,
		Usage:     utils.Usage("version", "v"),
		Run: func(c *grumble.Context) error {
			c.App.Println(version)
			utils.TimeoutTimer.Reset()
			return nil
		},
	})

	app.AddCommand(&grumble.Command{
		Name:      "elastic",
		Aliases:   []string{"es"},
		Help:      "elastic search command",
		HelpGroup: app.Config().Name,
		Usage:     utils.Usage("elastic"),
		Flags:     elasticsearch.AppFlags,
		AllowArgs: true,
		Run: func(c *grumble.Context) error {
			esApp, err := elasticsearch.NewApp(c)
			if err != nil {
				return err
			}
			if c.App.IsShell() {
				// 在交互模式下重置超时时间
				utils.TimeoutTimer.Reset()
				return esApp.App.Run()
			}
			return esApp.App.RunCommand(c.Args)
		},
	})
}
