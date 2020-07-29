package main

import (
	"encoding/json"
	"fmt"
)

/**
 * (C) Copyright 2019 Dashuo. All Rights Reserved.
 * 功能描述:
 * @Date: 2020/7/27 上午9:27
 * @author: cnopens
 * @email: cnopens@gmail.com
 */

func main() {
	type StuRead struct {
		Name  interface{} `json:"name"`
		Age   interface{}
		HIgh  interface{}
		sex   interface{}
		Class interface{} `json:"class"`
		Test  interface{}
	}


	//方式1：只声明，不分配内存
	var stus1 []*StuRead

	//方式2：分配初始值为0的内存
	stus2 := make([]*StuRead,0)

	//错误示范
	//new()只能实例化一个struct对象，而[]StuRead是切片，不是对象
	//stus := new([]StuRead)

	stu1 := &StuRead{"asd1",1,1,1,1,1}
	stu2 := &StuRead{"asd2",2,2,2,2,2}

	//由方式1和2创建的切片，都能成功追加数据
	//方式2最好分配0长度，append时会自动增长。反之指定初始长度，长度不够时不会自动增长，导致数据丢失
	stus1 = append(stus1,stu1)  //因为上面stus1是切片类型的结构体指针类型，所以append的类型也必须是取的地址。
	stus2 = append(stus2,stu2)  //因为上面stus2是切片类型的结构体指针类型，所以append的类型也必须是取的地址。

	//成功编码
	json1,_ := json.Marshal(stus1)
	json2,_ := json.Marshal(stus2)
	fmt.Println("===============01"+string(json1))
	fmt.Println("===============02"+string(json2))
}
//打印效果
//[{"name":"asd1","Age":1,"HIgh":1,"class":1,"Test":1}]
//[{"name":"asd2","Age":2,"HIgh":2,"class":2,"Test":2}]