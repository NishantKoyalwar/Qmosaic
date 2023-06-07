package main

import (
	"time"

	"main.go/components"
)

func main() {

	time.Sleep(10 * time.Second)
	components.Api()
	components.Remove()
	components.Img()
	components.Set()
	components.Autostart()
}
