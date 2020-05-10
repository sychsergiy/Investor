package cli

import (
	"log"
	"strings"
)

type Command interface {
	Execute()
}

type CLI struct {
	commands map[string]Command
}

func NewCLI() *CLI {
	return &CLI{make(map[string]Command)}
}

func (cli CLI) AddCommand(key string, command Command) {
	cli.commands[key] = command
}

func (cli CLI) AvailableCommands() string {
	var commands []string
	for name := range cli.commands {
		commands = append(commands, name)
	}

	return strings.Join(commands, ", ")
}

func (cli CLI) Run(key string) {
	command, ok := cli.commands[key]
	if !ok {
		log.Fatalf("unresolved command: %s, available commands: %s", key, cli.AvailableCommands())
	}
	command.Execute()
}
