package main

import (
	"github.com/Grandbusta/snapgo/cmd"
)

func main() {
	app := cmd.NewServer()
	app.Run()
}
