package utils

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

func isNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

func AppendOrSkip(parent bson.D, name string, data interface{}) bson.D {
	if data != nil && !isNilFixed(data) {
		parent = append(parent, bson.E{Key: name, Value: data})
		return parent
	}

	return parent
}
