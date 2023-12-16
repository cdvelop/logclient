package logclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (l *logClient) AddHandlerToRegisterLogsInDB(h *model.Handlers) {
	l.db = h.DataBaseAdapter
	l.time = h.TimeAdapter

	h.AddObjects(l.obj)

}

func (l *logClient) logError(t js.Value, p []js.Value) interface{} {

	if l.db != nil && l.time != nil {

		date, hour := l.time.DateToDayHour(true)

		err := l.db.CreateObjectsInDB(l.obj.Table, false, map[string]string{
			l.id:  date + "-" + hour,
			l.log: p[0].String(),
		})
		if err != "" {
			l.Log("error al registrar log en db:", p[0], err)
		}
	}

	return nil

}
