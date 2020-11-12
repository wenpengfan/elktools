package elasticsearch

import (
	"elktools/cmd/utils"
	"elktools/esapi"
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"

	"github.com/fatih/color"

	"github.com/desertbit/grumble"
)

var (
	defaultApp        *App
	defaultElasticURL = "http://192.168.112.135:9200"
)

func init() {
	if runtime.GOOS == "windows" {
		defaultElasticURL = "http://192.168.112.135:9200"
	}
}

var (
	commandMu sync.RWMutex
	commands  = make(map[string]func(string))
)

func Register(name string, command func(string)) {
	commandMu.Lock()
	defer commandMu.Unlock()
	if command == nil {
		panic("elastic: Register command is nil")
	}
	if _, dup := commands[name]; dup {
		panic("elastic: Register called twice for command " + name)
	}
	commands[name] = command
}

type App struct {
	App    *grumble.App
	search *esapi.Search
}

func NewApp(c *grumble.Context) (*App, error) {
	defaultApp = &App{
		App: grumble.New(&grumble.Config{
			Name:              "ELKTools » elastic",
			Description:       "elastic tools",
			PromptColor:       color.New(color.FgGreen, color.Bold),
			HelpHeadlineColor: color.New(color.FgGreen),
		}),
	}
	for name, command := range commands {
		command(name)
	}

	if err := defaultApp.connectElasticSearch(c); err != nil {
		return nil, err
	}
	c.App.Printf("elastic search %s connect success\n", defaultApp.search.Req.URL)
	return defaultApp, nil
}

func (a *App) InitRun(run func(c *grumble.Context) ([]byte, error)) func(c *grumble.Context) error {
	return func(c *grumble.Context) error {
		if a.App.IsShell() {
			// 在交互模式下重置超时时间
			utils.TimeoutTimer.Reset()
		}
		if defaultApp.search == nil {
			if err := a.connectElasticSearch(c); err != nil {
				return err
			}
		}
		context := *c
		data, err := run(&context)
		if err != nil {
			return err
		}
		layout := "2006-01-02 15:04:05.000"
		c.App.Println(time.Now().Format(layout))
		if err := utils.PipelinePrintln(data, c); err != nil {
			return err
		}

		interval := c.Flags.Uint("interval")
		if interval == 0 {
			return err
		}

		signals := make(chan os.Signal, 1)
		signal.Notify(signals)
		ticker := time.NewTicker(time.Duration(interval) * time.Second)
		for range ticker.C {
			if a.App.IsShell() {
				// 在交互模式下重置超时时间
				utils.TimeoutTimer.Reset()
			}
			context := *c
			select {
			case sign := <-signals:
				if sign == os.Interrupt {
					signal.Stop(signals)
					return nil
				}
			default:
				data, err := run(&context)
				if err != nil {
					return err
				}
				c.App.Println(time.Now().Format(layout))
				if err := utils.PipelinePrintln(data, c); err != nil {
					return err
				}
			}
		}
		return nil
	}

}

func (a *App) connectElasticSearch(c *grumble.Context) error {
	elasticURL, err := parseRequestURI(c.Flags.String("address"))
	if err != nil {
		return err
	}
	username := c.Flags.String("username")
	password := c.Flags.String("password")
	timeout := c.Flags.Duration("timeout")
	debug := c.Flags.Bool("debug")
	utils.SetDebug(debug)

	a.search = esapi.NewSearch(elasticURL, username, password, timeout)
	return err
}

func parseRequestURI(rawURI string) (string, error) {
	u, err := url.ParseRequestURI(rawURI)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
