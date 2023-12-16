package logclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func AddLoggerAdapter() *logClient {

	l := &logClient{
		id:  "id_log",
		log: "log",
	}

	l.obj = &model.Object{
		ObjectName: "logger",
		Table:      l.log,
		Fields: []model.Field{
			{Name: l.id, Legend: "id"},
			{Name: l.log, Legend: l.log},
		},
		Module: &model.Module{ModuleName: "logger"},
	}

	js.Global().Set("logError", js.FuncOf(l.logError))

	return l
}
