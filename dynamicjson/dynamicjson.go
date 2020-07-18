package dynamicjson

import (
	"reflect"
)

type DynamicJson struct {
	fields map[string]*Field
}

func New() *DynamicJson {
	return &DynamicJson{
		fields: make(map[string]*Field),
	}
}

func (dj DynamicJson) AddField(name string, tp interface{}, tag string) Builder {
	dj.fields[name] = &Field{
		Type: reflect.TypeOf(tp),
		Tag : reflect.StructTag(tag),
	}
	return dj
}

func (dj DynamicJson)MergeStructs(values ...interface{}) Builder {
	builder := New()

	for _, value := range values {
		valueOf := reflect.Indirect(reflect.ValueOf(value))
		typeOf := valueOf.Type()

		for i := 0; i < valueOf.NumField(); i++ {
			fval := valueOf.Field(i)
			ftyp := typeOf.Field(i)
			builder.AddField(ftyp.Name, fval.Interface(), string(ftyp.Tag))
		}
	}

	return builder
}

func (dj DynamicJson)RemoveField(name string) Builder {
	delete(dj.fields, name)
	return dj
}

func (dj DynamicJson)UpdateField(name string, tp interface{}, tag string) Builder {
	dj.fields[name] = &Field{
		Type: reflect.TypeOf(tp),
		Tag : reflect.StructTag(tag),
	}
	return dj
}

func (dj DynamicJson) GetField(name string)(*Field, bool) {
	field, exists := dj.fields[name]
	return field, exists
}

func (dj DynamicJson) Build() (newStruct interface{}) {
	if len(dj.fields) <= 0{
		return
	}

	fields := make([]reflect.StructField, 0)
	for name, config := range dj.fields {
		fields = append(fields, reflect.StructField{
			Name: name,
			Type: config.Type,
			Tag : config.Tag,
		})
	}
	newStruct = reflect.New(reflect.StructOf(fields)).Interface()
	return
}