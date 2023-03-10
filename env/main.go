package main

import (
	"fmt"

	"github.com/caarlos0/env/v7"
)

type envs struct {
	// the all fields must be capitalized.
	T string `env:"SHELL"`
}

func main() {
	envs := &envs{}
	if err := env.Parse(envs); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", envs)
	fmt.Println(envs.T)
}