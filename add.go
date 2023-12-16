package logclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func AddLoggerAdapter() *logClient {

	l := &logClient{
		log: "log",
	}

	l.obj = &model.Object{
		ObjectName: "logger",
		Table:      l.log,
		Fields: []model.Field{
			{Name: "id_log", Legend: "id"},
			{Name: l.log, Legend: l.log},
		},
		Module: &model.Module{ModuleName: "logger"},
	}

	js.Global().Set("logError", js.FuncOf(l.logError))

	return l
}
