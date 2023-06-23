package main

import (
	"fmt"
)

type person struct{
	name,city string
	age int8
}
// 构造函数
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
func (p person) eat(){
	fmt.Printf("%v在吃东西",p.name)
}
// 指针类型的接受者
func (p *person) SetAge(newAge int8){
	p.age = newAge
}
// 值类型的接受者
func (p person) SetAge2(newAge int8){
	p.age = newAge
}
func main() {		
	// p1 := newPerson("张三", "沙河", 90)
	// fmt.Printf("%#v\n", p1) //&main.person{name:"张三", city:"沙河", age:90}
	// p2 := newPerson("李四", "北京", 99)
	// fmt.Printf("%#v\n", p2) //&main.person{name:"李四", city:"北京", age:99}

	// Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）
	// 函数不需要任何类型 方法属于特定的类型
	// func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	// 		函数体
	// }
	p1 := newPerson("张三", "沙河", 90)
	p1.eat() //张三在吃东西
  fmt.Println(p1.age) //90

	//  指针类型的接受者
	// p1.SetAge(30)
	// fmt.Println(p1.age) //30

	// 值类型的接受者 无法修改接收者变量本身
  p1.SetAge2(30) // (*p1).SetAge2(30)
	fmt.Println(p1.age) // 90




}