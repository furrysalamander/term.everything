package main

import (
	"github.com/furrysalamander/term.everything/termeverything"
)

//go:generate go generate ./wayland

func main() {
	termeverything.MainLoop()
}
