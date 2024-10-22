package main

import (
	"log"
	"os"

	"github.com/Vikuuu/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	st := &state{
		Config: &cfg,
	}
	cmds := &commands{
		Command: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("error, arguments is less than 2")
		return
	}

	cmd := command{Name: args[1], Argument: args[2:]}
	err = cmds.run(st, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
