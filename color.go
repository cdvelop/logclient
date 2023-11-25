package logclient

import (
	"syscall/js"

	"github.com/cdvelop/strings"
)

// class ej: war,error
func colorMessage(msg string) {

	var class string

	switch {
	case strings.Contains(msg, "error") != 0:
		class = "white-red"
	case strings.Contains(msg, "tests") != 0:
		class = "white-blue"
	case strings.Contains(msg, "OK") != 0:
		class = "white-blue"
	case strings.Contains(msg, "test") != 0:
		class = "blue"
	}

	var color string

	switch class {
	case "black-yellow":
		color = "color: black; background: yellow;"
	case "white-red":
		color = "color: white; background-color: red;"
	case "white-blue":
		color = "color: white; background-color: blue;"
	case "blue":
		color = "color: blue"
	}

	js.Global().Get("console").Call("log", "%c"+msg, color)
}
