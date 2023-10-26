package logclient

import "syscall/js"

func (Log) Log(message ...any) interface{} {

	for i, msg := range message {
		// Comprueba si el mensaje es de tipo error
		if err, isError := msg.(error); isError {
			// Edita el mensaje y conviértelo a string
			message[i] = err.Error()
		}
	}

	js.Global().Get("console").Call("log", message...)

	return nil
}
