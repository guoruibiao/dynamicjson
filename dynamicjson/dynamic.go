package dynamicjson

import "reflect"

type Builder interface {
	AddField(name string, tp interface{}, tag string) Builder
	RemoveField(name string) Builder
	UpdateField(name string, tp interface{}, tag string) Builder
	GetField(name string) (*Field, bool)
	Build() interface{}
}


type Field struct {
	Type reflect.Type
	Tag  reflect.StructTag
}