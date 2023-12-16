package logclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (l *logClient) AddHandlerToRegisterLogsInDB(h *model.Handlers) {
	l.db = h.DataBaseAdapter

	h.AddObjects(l.obj)

}

func (l *logClient) logError(t js.Value, p []js.Value) interface{} {

	if l.db != nil {
		err := l.db.CreateObjectsInDB(l.obj.Table, true, map[string]string{
			l.log: p[0].String(),
		})
		if err != "" {
			l.Log("error al registrar log en db:", p[0], err)
		}
	}

	return nil

}
