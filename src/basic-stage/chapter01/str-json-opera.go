package main

import (
	"fmt"
	"strings"
)

/**
 * (C) Copyright 2019 Dashuo. All Rights Reserved.
 * 功能描述:
 * @Date: 2020/7/29 上午8:41
 * @author: cnopens
 * @email: cnopens@gmail.com
 */

func main() {

}

//字符串操作
func strOpera()  {
	//str operation
	fmt.Println(strings.Contains("caixiaofeng","cai"))

	//
	aniamal := []string{"cat", "dog", "tiger"}
	fmt.Println(strings.Join(aniamal, "| "))


	fmt.Println(strings.Index("chicken", "ken"))

	fmt.Println("ba" + strings.Repeat("na", 2))

	fmt.Println(strings.Replace("ok ok ok", "k", "ky", 2))


	fmt.Printf("%q\n", strings.Split("a,b,c", ","))

	fmt.Println( strings.Fields("  a b  c   "))
}
