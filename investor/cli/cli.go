package cli

import "log"

type Command interface {
	Execute()
}

type CLI struct {
	commands map[string]Command
}

func NewCLI() CLI {
	return CLI{make(map[string]Command)}
}

func (cli CLI) AddCommand(key string, command Command) {
	cli.commands[key] = command
}

func (cli CLI) Run(key string) {
	command, ok := cli.commands[key]
	if !ok {
		log.Fatal("unresolved command: " + key)
	}
	command.Execute()
}
