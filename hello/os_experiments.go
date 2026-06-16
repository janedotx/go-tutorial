package main

import (
	"fmt"
	"os"
)

func TestOsGetEnv() {
	e := os.Getenv("NONEXISTENT_ENV_VAR")
	fmt.Printf("the value of the env var is %s\n", e)
}