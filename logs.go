package logclient

import (
	"strconv"
	"syscall/js"

	"github.com/cdvelop/model"
)

// support: error, map[string]string
func (logClient) Log(message ...interface{}) {
	var out_any []interface{}
	var out_string string
	var space string

	for _, msg := range message {

		switch msg := msg.(type) {

		case string:
			out_string += space + msg

		case error:
			if msg != nil {
				out_string += space + msg.Error()
			}

		case bool:
			out_any = append(out_any, msg)

		case int:
			out_string += space + strconv.Itoa(msg)

		case int64:
			out_string += space + strconv.FormatInt(msg, 10)

		case *js.Value:
			out_any = append(out_any, *msg)

		case js.Value:
			out_any = append(out_any, msg)

		case map[string]string:
			new := make(map[string]interface{})
			for k, v := range msg {
				new[k] = v
			}
			out_any = append(out_any, new)
		case []map[string]string:
			for _, items := range msg {
				new := make(map[string]interface{})
				for k, v := range items {
					new[k] = v
				}
				out_any = append(out_any, new)
			}

		case map[string]interface{}:
			out_any = append(out_any, msg)

		case []map[string]interface{}:
			for _, item := range msg {
				out_any = append(out_any, item)
			}

		case []string:
			// Convert []string to JavaScript array
			jsArray := js.Global().Get("Array").New(len(msg))
			for i, item := range msg {
				jsArray.SetIndex(i, item)
			}
			out_any = append(out_any, jsArray)

		case model.Response:
			out_any = append(out_any, map[string]interface{}{
				"Action":   msg.Action,
				"Object  ": msg.Object,
				"Message":  msg.Message,
				"Data":     msg.Data,
			})

		}
		space = " "
	}

	if len(out_string) != 0 {
		colorMessage(out_string)
	}

	if len(out_any) != 0 {
		js.Global().Get("console").Call("log", out_any...)
	}

}
