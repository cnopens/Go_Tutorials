#Basic Stage

go func 闭包




所谓闭包是指内层函数引用了外层函数中的变量或称为引用了自由变量的函数，其返回值也是一个函数，了解过的语言中有闭包概念的像 js，python，golang 都类似这样。

python 中的闭包可以嵌套函数，像下面这样：

def make_adder(addend):
    def adder(augend):
        return augend + addend
    return adder

转化成 golang 代码则像下面这样：

func outer(x int) func(int) int{
    func inner(y int) int{
        return x + y
    }
    return inner
}

当然这是错误的，golang 中是不能嵌套函数的，但是可以在一个函数中包含匿名函数，完整示例像下面这样：

package main
 
import (
    "fmt"
)
 
func outer(x int) func(int) int {
    return func(y int) int {
        return x + y 
    }   
}
 
func main() {
    f := outer(10)
    fmt.Println(f(100))
} 

看了一段时间 golang 后，对于 golang 中的闭包可能出现的坑大致有下面几个。
1，for range 中使用闭包

一个示例：

func main() {                
    s := []string{"a", "b", "c"}                             
    for _, v := range s { 
        go func() {
            fmt.Println(v)
        }()                 
    }                        
    select {}    // 阻塞模式                                                         
}                          

// 嗯，结果应该是 a,b,c 吧

来看看结果：
 
运行结果

输出的结果不期而然，大家的结果也不一定和我相同。

对比下面的改进：

func main() {                
    s := []string{"a", "b", "c"}                             
    for _, v := range s { 
        go func(v string) {
            fmt.Println(v)
        }(v)   //每次将变量 v 的拷贝传进函数                 
    }                        
    select {}                                                      
}  

所以结果当然是:

"a"
"b"
"c"

由于使用了 go 协程，并非顺序输出。

    解释：也不用多解释了吧，在没有将变量 v 的拷贝值传进匿名函数之前，只能获取最后一次循环的值,这是新手最容易遇到的坑。

2，函数列表使用不当

package main

import (
    "fmt"
)

func test() []func() {
    var s []func()

    for i := 0; i < 3; i++ {
        s = append(s, func() {  //将多个匿名函数添加到列表
            fmt.Println(&i, i)
        })
    }

    return s    //返回匿名函数列表
}
func main() {
    for _, f := range test() {  //执行所有匿名函数
        f()   
    }
}

运行结果：
 
运行结果

解决方法：

package main

import (
    "fmt"
)

func test() []func() {
    var s []func()
    
    for i := 0; i < 3; i++ {
        x := i                  //复制变量
        s = append(s, func() {
            fmt.Println(&x, x)
        })
    }

    return s
}
func main() {
    for _, f := range test() {
        f()
    }
}

    解释：每次 append 操作仅将匿名函数放入到列表中，但并未执行，并且引用的变量都是 i，随着 i 的改变匿名函数中的 i 也在改变，所以当执行这些函数时，他们读取的都是环境变量 i 最后一次的值。解决的方法就是每次复制变量 i 然后传到匿名函数中，让闭包的环境变量不相同。

若是你对闭包理解了，也可以利用闭包来修改全局变量：

package main                   
                               
import (                       
    "fmt"                      
)                              
                               
var x int = 1                  
                               
func main() {                  
    y := func() int {          
        x += 1                 
        return x               
    }()                        
    fmt.Println("main:", x, y)                                                            
} 

// 结果：    main: 2 2

3，延迟调用

defer 调用会在当前函数执行结束前才被执行，这些调用被称为延迟调用，
defer 中使用匿名函数依然是一个闭包。

package main

import "fmt"

func main() {
    x, y := 1, 2

    defer func(a int) { 
        fmt.Printf("x:%d,y:%d\n", a, y)  // y 为闭包引用
    }(x)      // 复制 x 的值

    x += 100
    y += 100
    fmt.Println(x, y)
}

输出结果：

101 102
x:1,y:102

总结：从形式上看，匿名函数都是闭包。闭包的使用非常灵活，上面仅是几个比较简单的示例，不当的使用容易产生难以发现的 bug，当出现意外情况时，首先检查函数的参数，声明可以接收参数的匿名函数，这些类型的闭包问题也就引刃而解了。
