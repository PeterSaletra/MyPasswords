package cli

import (
	"errors"
	"fmt"
	"io"
	"mypasswords/commands/shell"
	"mypasswords/crypto"
	"mypasswords/store"
	"strings"

	"math/rand"

	"github.com/chzyer/readline"
	"github.com/mattn/go-shellwords"
)

type Cli struct {
	rl    *readline.Instance
	shell *shell.Shell
}

func NewCli(db *store.Database, keys *crypto.Keys) (*Cli, error) {
	config, err := GetConfig("main")
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}
	l, err := readline.NewEx(config)
	if err != nil {
		return nil, fmt.Errorf("New readline error: %v", err)
	}

	s := shell.NewShell(db, keys)
	s.PrepareCommands()
	return &Cli{
		rl:    l,
		shell: s,
	}, nil
}

func (c *Cli) Start() {
	fmt.Print("\033[H\033[2J")
	showBanner()

	for {
		line, err := c.rl.Readline()
		if errors.Is(err, readline.ErrInterrupt) {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue // Skip empty line
		}
		var cmd []string
		cmd, err = shellwords.Parse(line)
		if err != nil {
			fmt.Printf("shell parse error: %v\n", err)
			continue
		}

		c.shell.Execute(cmd)

	}
}

func showBanner() {
	colors := []string{
		"\033[31m", // Red
		"\033[32m", // Green
		"\033[33m", // Yellow
		"\033[34m", // Blue
		"\033[35m", // Magenta
		"\033[36m", // Cyan
		"\033[37m", // White
	}
	color := colors[rand.Intn(len(colors))]

	banner := []string{
		"  __  __       _____                                    _     ",
		" |  \\/  |     |  __ \\                                  | |    ",
		" | \\  / |_   _| |__) |_ _ ___ _____      _____  _ __ __| |___ ",
		" | |\\/| | | | |  ___/ _` / __/ __\\ \\ /\\ / / _ \\| '__/ _` / __|",
		" | |  | | |_| | |  | (_| \\__ \\__ \\\\ V  V / (_) | | | (_| \\__ \\",
		" |_|  |_|\\__, |_|   \\__,_|___/___/ \\_/\\_/ \\___/|_|  \\__,_|___/",
		"         __/ |                                                ",
		"        |___/                                                 ",
		"",
	}

	fmt.Print(color + strings.Join(banner, "\n") + "\033[0m")

	header := `
	Welcom to MyPassword,
	Your private password manager,

   	`

	fmt.Print(header)
}
