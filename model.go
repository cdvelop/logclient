package logclient

import "github.com/cdvelop/model"

type logClient struct {
	db model.DataBaseAdapter

	obj *model.Object

	log string
}
