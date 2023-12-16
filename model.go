package logclient

import "github.com/cdvelop/model"

type logClient struct {
	db   model.DataBaseAdapter
	time model.TimeAdapter

	obj *model.Object

	id  string
	log string
}
