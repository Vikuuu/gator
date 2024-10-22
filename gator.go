package main

import (
	"errors"
	"fmt"

	"github.com/Vikuuu/gator/internal/config"
)

type state struct {
	Config *config.Config
}

type command struct {
	Name     string
	Argument []string
}

type commands struct {
	Command map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.Command[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handler, exist := c.Command[cmd.Name]
	if !exist {
		return errors.New("Command does not exists")
	}

	err := handler(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Argument) != 1 {
		return errors.New("you must enter a username")
	}

	err := s.Config.SetUser(cmd.Argument[0])
	if err != nil {
		return err
	}
	fmt.Println("User has been set!")

	return nil
}
