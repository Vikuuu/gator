package main

import "fmt"

func handlerHelp(c *commands, s *state, cmd command) error {
	fmt.Println("Welcome To Gator! A RSS feed in GO!")
	fmt.Println("Usage: gator <command> [...Args]")
	fmt.Println("Available commands:")

	for name, info := range c.registeredCommands {
		fmt.Printf(" %-10s %-30s %s\n", name, info.usage, info.description)
	}
	return nil
}
