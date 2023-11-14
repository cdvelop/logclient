package logclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

// support: error, map[string]string
func (Log) Log(message ...interface{}) interface{} {
	var out []interface{}

	for _, msg := range message {

		switch msg := msg.(type) {

		case string, int:
			out = append(out, msg)

		case error:
			out = append(out, msg.Error())

		case *js.Value:
			out = append(out, *msg)

		case js.Value:
			out = append(out, msg)

		case map[string]string:
			new := make(map[string]interface{})
			for k, v := range msg {
				new[k] = v
			}
			out = append(out, new)
		case []map[string]string:
			for _, items := range msg {
				new := make(map[string]interface{})
				for k, v := range items {
					new[k] = v
				}
				out = append(out, new)
			}

		case map[string]interface{}:
			out = append(out, msg)

		case []map[string]interface{}:
			for _, item := range msg {
				out = append(out, item)
			}

		case []string:
			// Convert []string to JavaScript array
			jsArray := js.Global().Get("Array").New(len(msg))
			for i, item := range msg {
				jsArray.SetIndex(i, item)
			}
			out = append(out, jsArray)

		case model.Response:
			out = append(out, map[string]interface{}{
				"Action":   msg.Action,
				"Object  ": msg.Object,
				"Message":  msg.Message,
				"Data":     msg.Data,
			})

		}
	}

	js.Global().Get("console").Call("log", out...)
	return nil
}
