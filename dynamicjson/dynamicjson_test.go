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

func TestDynamicJson_RemoveField2(t *testing.T) {
	type sStruct struct {
		Name string `json:"name"`
		Age int `json:"age"`
		Address []string `json:"address"`
	}
	s := sStruct{
		Name: "泰戈尔🤩",
		Age: 25,
		Address: []string{"曼彻斯特", "也门"},
	}
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