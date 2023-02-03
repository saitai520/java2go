## 数据类型

数据类型是入门任何一门语言都要优先掌握的，虽然它很简单。我们仍然是通过与Java数据类型对比的方式来学习。

#### 字符串

`string`

只能用一对双引号（""）或反引号（``）括起来定义，不能用单引号（''）定义！

#### 布尔

`bool`

只有 true 和 false，默认为 false。

#### 数字

**整型**

`int8` `uint8`

`int16` `uint16`

`int32` `uint32`

`int64` `uint64`

`int` `uint`，具体长度取决于 CPU 位数。

**浮点型**

`float32` `float64`

## 常量声明

**常量**，在程序编译阶段就确定下来的值，而程序在运行时无法改变该值。

**单个常量声明**

第一种：const 变量名称 数据类型 = 变量值

如果不赋值，使用的是该数据类型的默认值。

第二种：const 变量名称 = 变量值

根据变量值，自行判断数据类型。

**多个常量声明**

第一种：const 变量名称,变量名称 ... ,数据类型 = 变量值,变量值 ...

第二种：const 变量名称,变量名称 ...  = 变量值,变量值 ...

**测试代码**

```
//demo_1.go
package main

import (
	"fmt"
)

func main() {
	const name string = "Tom"
	fmt.Println(name)

	const age = 30
	fmt.Println(age)

	const name_1, name_2 string = "Tom", "Jay"
	fmt.Println(name_1, name_2)

	const name_3, age_1 = "Tom", 30
	fmt.Println(name_3, age_1)
}
```

运行结果：

![](https://github.com/xinliangnote/Go/blob/master/00-基础语法/images/02-变量声明/2_go_1.png)

## 变量声明

**单个变量声明**

第一种：var 变量名称 数据类型 = 变量值

如果不赋值，使用的是该数据类型的默认值。

第二种：var 变量名称 = 变量值

根据变量值，自行判断数据类型。

第三种：变量名称 := 变量值

省略了 var 和数据类型，变量名称一定要是未声明过的。

**多个变量声明**

第一种：var 变量名称,变量名称 ... ,数据类型 = 变量值,变量值 ...

第二种：var 变量名称,变量名称 ...  = 变量值,变量值 ...

第三种：变量名称,变量名称 ... := 变量值,变量值 ...