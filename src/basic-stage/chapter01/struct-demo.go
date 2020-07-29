package main

/**
 * (C) Copyright 2019 Dashuo. All Rights Reserved.
 * 功能描述:
 * @Date: 2020/7/27 上午9:17
 * @author: cnopens
 * @email: cnopens@gmail.com
 */

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	Name  string `json:"name"`
	Age   int
	HIgh  bool
	sex   string
	Class *Class `json:"class"`
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
	cla := new(Class)   //这个new方法，就相当于  cla := &Class{}，是一个取地址的操作。
	cla.Name = "1班"
	cla.Grade = 3
	stu.Class = cla

	//Marshal失败时err!=nil
	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	//jsonStu是[]byte类型，转化成string类型便于查看
	fmt.Println(string(jsonStu))
}

/***
Note:
只要是可导出成员（变量首字母大写），都可以转成json。因成员变量sex是不可导出的，故无法转成json。

如果变量打上了json标签，如Name旁边的 `json:"name"` ，那么转化成的json key就用该标签“name”，否则取变量名作为key，如“Age”，“HIgh”。

bool类型也是可以直接转换为json的value值。Channel， complex 以及函数不能被编码json字符串。当然，循环的数据结构也不行，它会导致marshal陷入死循环。

指针变量，编码时自动转换为它所指向的值，如cla变量。
（当然，不传指针，Stu struct的成员Class如果换成Class struct类型，效果也是一模一样的。只不过指针更快，且能节省内存空间。）

最后，强调一句：json编码成字符串后就是纯粹的字符串了。

上面的成员变量都是已知的类型，只能接收指定的类型，比如string类型的Name只能赋值string类型的数据。
但有时为了通用性，或使代码简洁，我们希望有一种类型可以接受各种类型的数据，并进行json编码。这就用到了interface{}类型。
 */