package main

import "errors"

type commands struct {
	cmdFuncMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	function, ok := c.cmdFuncMap[cmd.name]
	if !ok {
		return errors.New("this command is not yet registered")
	}
	err := function(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdFuncMap[name] = f
}
