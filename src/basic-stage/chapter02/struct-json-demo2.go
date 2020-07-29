package main

import (
	"encoding/json"
	"fmt"
)

/**
 * (C) Copyright 2019 Dashuo. All Rights Reserved.
 * 功能描述:
 * @Date: 2020/7/27 上午9:20
 * @author: cnopens
 * @email: cnopens@gmail.com
 */

type Stu struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	HIgh  interface{}
	sex   interface{}
	Class interface{} `json:"class"`
}

type Class struct {
	Name  string
	Grade int
}

func main() {
	//实例化一个数据结构，用于生成json字符串
	stu := Stu{
		Name: "张三",
		Age:  18,
		HIgh: true,
		sex:  "男",
	}

	//指针变量
	cla := new(Class)
	cla.Name = "1班"
	cla.Grade = 3
	stu.Class = cla

	//Marshal失败时err!=nil
	jsonStu, err := json.Marshal(stu);
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	//jsonStu是[]byte类型，转化成string类型便于查看
	fmt.Println(string(jsonStu))
}
//打印效果
//{"name":"张三","Age":18,"HIgh":true,"class":{"Name":"1班","Grade":3}}
//从结果中可以看出，无论是string，int，bool，还是指针类型等，都可赋值给interface{}类型，且正常编码，效果与前面的例子一样。