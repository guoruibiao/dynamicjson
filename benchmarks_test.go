package main

import (
	"encoding/json"
	"fmt"
	"github.com/guoruibiao/dynamicjson/dynamicjson"
	"strings"
	"testing"
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

func Benchmark_New(b *testing.B) {
	for i:= 0; i<b.N; i++ {
		dynamicjson.New().Build()
	}
}

func Benchmark_UpdateField(b *testing.B) {
	for i:=0; i<b.N; i++ {
		updateField()
	}
}

func Benchmark_RemoveField(b *testing.B) {
	for i:=0; i<b.N; i++ {
		removeField()
	}
}

func Benchmark_AddField(b *testing.B) {
	for i:=0; i<b.N; i++ {
		addField()
	}
}