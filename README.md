# dynamicjson
Give the power to dynamically serialize json to Golang.

go 后端输出 json时，有些时候需要对结果 json 串进行修改，单纯只是通过 omitempty 来控制
不是万全之策，而一个比较好的方式就是做切面处理，在最终输出时进行拦截处理。

```shell script
go get -u github.com/guoruibiao/dynamicjson
```


## demo1
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/guoruibiao/dynamicjson/dynamicjson"
	"strings"
)

type UserInfo struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Address []string `json:"address"`
	Items []*UserInfoItem `json:"items"`
}

type UserInfoItem struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func removeField() {
	userInfo := UserInfo{
		Name: "泰戈尔🤩",
		Age: 25,
		Address: []string{"圣馨家园", "国展新座"},
		Items: make([]*UserInfoItem, 0),
	}
	userInfo.Items = append(userInfo.Items, &UserInfoItem{
		Name: "Marksinoberg",
		Description: "CSDN博客笔名",
	})
	userInfo.Items = append(userInfo.Items, &UserInfoItem{
		Name: "泰戈尔",
		Description: "泰戈尔=tiger",
	})
	jsonBytes, _ := json.Marshal(userInfo)
	newStruct := dynamicjson.New().MergeStructs(userInfo).RemoveField("Address").Build()
	_ = json.Unmarshal(jsonBytes, &newStruct)
	jsonBytes, _ = json.Marshal(newStruct)
	// 最终 json 输出串
	fmt.Println(string(jsonBytes))
}


func addField() {
	userInfo := UserInfo{
		Name: "泰戈尔🤩",
		Age: 25,
		Address: []string{"圣馨家园", "国展新座"},
		Items: make([]*UserInfoItem, 0),
	}
	userInfo.Items = append(userInfo.Items, &UserInfoItem{
		Name: "Marksinoberg",
		Description: "CSDN博客笔名",
	})
	userInfo.Items = append(userInfo.Items, &UserInfoItem{
		Name: "泰戈尔",
		Description: "泰戈尔=tiger",
	})
	jsonBytes, _ := json.Marshal(userInfo)
	fmt.Println(string(jsonBytes))
	newStruct := dynamicjson.New().MergeStructs(userInfo).AddField("School", []string{}, `json:"school,omitempty"`).Build()
	jsonBytes = []byte(strings.TrimRight(string(jsonBytes), `}`) + `, "school":["大连理工大学", "北京大学"]}`)
	fmt.Println(string(jsonBytes))
	_ = json.Unmarshal(jsonBytes, &newStruct)
	jsonBytes, _ = json.Marshal(newStruct)
	// 最终 json 输出串
	fmt.Println(string(jsonBytes))
}

func updateField() {
	userInfo := UserInfo{
		Name: "泰戈尔🤩",
		Age: 25,
		Address: []string{"圣馨家园", "国展新座"},
		Items: make([]*UserInfoItem, 0),
	}
	userInfo.Items = append(userInfo.Items, &UserInfoItem{
		Name: "Marksinoberg",
		Description: "CSDN博客笔名",
	})
	userInfo.Items = append(userInfo.Items, &UserInfoItem{
		Name: "泰戈尔",
		Description: "泰戈尔=tiger",
	})

	newStruct := dynamicjson.New().MergeStructs(userInfo).UpdateField("Address", "", `json:"address,omitempty"`).Build()
	jsonBytes := []byte(`{"name":"泰戈尔🤩","age":25,"address":"上海市陆家嘴服务中心","items":[{"name":"Marksinoberg","description":"CSDN博客笔名"},{"name":"泰戈尔","description":"泰戈尔=tiger"}]}`)
	_ = json.Unmarshal(jsonBytes, &newStruct)
	jsonBytes, _ = json.Marshal(newStruct)
	// 最终 json 输出串
	fmt.Println(string(jsonBytes))
}

func main() {
	removeField()
	addField()
	updateField()
}
```


## demo2

```go
package dynamicjson

import (
	"encoding/json"
	"testing"
)

var jsonTemplate string

func init() {
	jsonTemplate = `{"name":"泰戈尔🤩", "age":25, "address":["通州新建村", "圣馨家园", "国展新座", "曙光里社区", "树村丽景苑"]}`
}


func TestDynamicJson_AddField(t *testing.T) {
	s := New().
		AddField("Name", "", `json:"name,omitempty"`).
		AddField("Age", 1, `json:"age,omitempty"`).
		AddField("Address", []string{}, `json:"address,omitempty"`).
		Build()
	t.Log(s)
	if err := json.Unmarshal([]byte(jsonTemplate), &s); err != nil {
		t.Error(err)
	}else{
		t.Log(s)
	}
	t.Log("================================")
	if jsonBytes, err := json.Marshal(s); err != nil {
		t.Error(err)
	}else{
		t.Log("TestDynamicJson_AddField: ", string(jsonBytes))
	}
}


func TestDynamicJson_RemoveField(t *testing.T) {
	s := New().
		AddField("Name", "", `json:"name,omitempty"`).
		AddField("Age", 1, `json:"age,omitempty"`).
		AddField("Address", []string{}, `json:"address,omitempty"`).
		Build()
	t.Log(s)
	s2 := New().MergeStructs(s).RemoveField("Address").Build()
	if err := json.Unmarshal([]byte(jsonTemplate), &s2); err != nil {
		t.Error(err)
	}else{
		t.Log(s)
	}
	t.Log("================================")
	if jsonBytes, err := json.Marshal(s2); err != nil {
		t.Error(err)
	}else{
		t.Log("TestDynamicJson_RemoveField: ", string(jsonBytes))
	}
}

func TestDynamicJson_UpdateField(t *testing.T) {
	s := New().
		AddField("Name", "", `json:"name,omitempty"`).
		AddField("Age", 1, `json:"age,omitempty"`).
		AddField("Address", []string{}, `json:"address,omitempty"`).
		Build()
	t.Log(s)
	s2 := New().
		MergeStructs(s).
		RemoveField("Address").
		AddField("Address", "", `json:"address"`).
		Build()
	jsonTemplate = `{"name":"泰戈尔🤩", "age":25, "address":"北京市海淀区后厂村村草"}`
	if err := json.Unmarshal([]byte(jsonTemplate), &s2); err != nil {
		t.Error(err)
	}else{
		t.Log(s)
	}
	t.Log("================================")
	if jsonBytes, err := json.Marshal(s2); err != nil {
		t.Error(err)
	}else{
		t.Log("TestDynamicJson_UpdateField: ", string(jsonBytes))
	}
}

func TestDynamicJson_UpdateField2(t *testing.T) {
	type sStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Address []string `json:"address"`
	}
	s := sStruct{
		Name: "后厂村村草",
		Age: 25,
		Address: []string{"大连", "北京"},
	}
	s2 := New().
		MergeStructs(s).
		RemoveField("Address").
		AddField("Address", "", `json:"address"`).
		Build()
	jsonTemplate = `{"name":"泰戈尔🤩", "age":25, "address":"北京市海淀区后厂村村草"}`
	if err := json.Unmarshal([]byte(jsonTemplate), &s2); err != nil {
		t.Error(err)
	}else{
		t.Log(s)
	}
	t.Log("================================")
	if jsonBytes, err := json.Marshal(s2); err != nil {
		t.Error(err)
	}else{
		t.Log("TestDynamicJson_UpdateField: ", string(jsonBytes))
	}
}
```

```shell script
➜  dynamicjson git:(master) go test -v ./...
?   	github.com/guoruibiao/dynamicjson	[no test files]
=== RUN   TestDynamicJson_AddField
    TestDynamicJson_AddField: dynamicjson_test.go:21: &{[]  0}
    TestDynamicJson_AddField: dynamicjson_test.go:25: &{[通州新建村 圣馨家园 国展新座 曙光里社区 树村丽景苑] 泰戈尔🤩 25}
    TestDynamicJson_AddField: dynamicjson_test.go:27: ================================
    TestDynamicJson_AddField: dynamicjson_test.go:31: TestDynamicJson_AddField:  {"address":["通州新建村","圣馨家园","国展新座","曙光里社区","树村丽景苑"],"name":"泰戈尔🤩","age":25}
--- PASS: TestDynamicJson_AddField (0.00s)
=== RUN   TestDynamicJson_RemoveField
    TestDynamicJson_RemoveField: dynamicjson_test.go:42: &{[]  0}
    TestDynamicJson_RemoveField: dynamicjson_test.go:47: &{[]  0}
    TestDynamicJson_RemoveField: dynamicjson_test.go:49: ================================
    TestDynamicJson_RemoveField: dynamicjson_test.go:53: TestDynamicJson_RemoveField:  {"name":"泰戈尔🤩","age":25}
--- PASS: TestDynamicJson_RemoveField (0.00s)
=== RUN   TestDynamicJson_UpdateField
    TestDynamicJson_UpdateField: dynamicjson_test.go:63: &{ 0 []}
    TestDynamicJson_UpdateField: dynamicjson_test.go:73: &{ 0 []}
    TestDynamicJson_UpdateField: dynamicjson_test.go:75: ================================
    TestDynamicJson_UpdateField: dynamicjson_test.go:79: TestDynamicJson_UpdateField:  {"name":"泰戈尔🤩","age":25,"address":"北京市海淀区后厂村村草"}
--- PASS: TestDynamicJson_UpdateField (0.00s)
=== RUN   TestDynamicJson_UpdateField2
    TestDynamicJson_UpdateField2: dynamicjson_test.go:103: {后厂村村草 25 [大连 北京]}
    TestDynamicJson_UpdateField2: dynamicjson_test.go:105: ================================
    TestDynamicJson_UpdateField2: dynamicjson_test.go:109: TestDynamicJson_UpdateField:  {"name":"泰戈尔🤩","age":25,"address":"北京市海淀区后厂村村草"}
--- PASS: TestDynamicJson_UpdateField2 (0.00s)
PASS
ok  	github.com/guoruibiao/dynamicjson/dynamicjson	(cached)
```