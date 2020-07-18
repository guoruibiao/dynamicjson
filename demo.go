package main

import (
"encoding/json"
"fmt"
"reflect"
	"strings"
)


type MyProfile struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Address []string `json:"address"`
	School []string `json:"school"`
}

func main() {
	fields := []reflect.StructField{
		{
			Name: "Name",
			Type: reflect.TypeOf(""),
			Tag: reflect.StructTag(`json:"name,omitempty"`),
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(1),
			Tag: reflect.StructTag(`json:"age,omitempty"`),
		},
		{
			Name: "Address",
			Type: reflect.TypeOf([]string{}),
			Tag: reflect.StructTag(`json:"address,omitempty"`),
		},
	}
	tmpStruct := reflect.StructOf(fields)
	fmt.Printf("%+v\n", tmpStruct)
	instance := reflect.New(tmpStruct).Interface()
	//mytype := reflect.New(tmpStruct).Interface()
	jsonTemplate := `{"name":"嘎嘣", "age":25, "address":["通州新建村", "圣馨家园", "国展新座", "曙光里社区", "树村丽景苑"]}`
	if err := json.Unmarshal([]byte(jsonTemplate), &instance); err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%+v\n", instance)
	}

	// 对 interface 进行配置读取
	// 静态语言无法直接进行反序列化的内容，归置到某一个结构体上
	// 将所有的 field 信息读取到一个 map 上，然后即可被反序列化的实时处理掉
	valueOf := reflect.Indirect(reflect.ValueOf(instance))
	typeOf  := valueOf.Type()
	if typeOf.Kind() == reflect.Struct {
		for i:=0; i < valueOf.NumField(); i++ {
			field := typeOf.Field(i)
			fmt.Println("Field=", field)
		}
	}

	fmt.Println("==========================")
	fmt.Println("==========================")

	// 假设现在 json 序列化的时候 删除某一个属性
	fields = fields[0:2]
	fmt.Println(fields)
	instance = reflect.New(reflect.StructOf(fields)).Interface()
	if err := json.Unmarshal([]byte(jsonTemplate), &instance); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(instance)
	}

	// 假设现在 json 序列化的时候 增加一个新的属性
	fields = append(fields, reflect.StructField{
		Name: "School",
		Type: reflect.TypeOf([]string{}),
		Tag:  reflect.StructTag(`json:"school,omitempty"`),
	})
	jsonTemplate = strings.TrimRight(jsonTemplate, "}")
	jsonTemplate += `, "school": ["大连理工大学", "北京大学"]}`
	fmt.Println(jsonTemplate)
	instance = reflect.New(reflect.StructOf(fields)).Interface()
	if err := json.Unmarshal([]byte(jsonTemplate), &instance); err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(instance)
	}

}
