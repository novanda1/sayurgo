package utils

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

func AppendOrSkip(parent bson.D, name string, data interface{}) bson.D {
	if data != nil && !reflect.ValueOf(data).IsNil() {
		parent = append(parent, bson.E{Key: name, Value: data})
		return parent
	}

	return parent
}
