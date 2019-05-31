package main //必须有一个main包

import "fmt" //fmt.Println()

type Person struct { //人
	name string
	sex  byte
	age  int
}

type Student struct { //学生
	*Person // 匿名字段，结构体指针类型
	id      int
	addr    string
}
func main() {
	//初始化
	s1 := Student{&Person{"mike", 'm', 18}, 1, "bj"}

	//{Person:0xc0420023e0 id:1 addr:bj}
	fmt.Printf("%+v\n", s1)
	//mike, m, 18
	fmt.Printf("%s, %c, %d\n", s1.name, s1.sex, s1.age)

	//声明变量
	var s2 Student
	s2.Person = new(Person) //分配空间
	s2.name = "yoyo"
	s2.sex = 'f'
	s2.age = 20

	s2.id = 2
	s2.addr = "sz"

	//yoyo 102 20 2 20
	fmt.Println(s2.name, s2.sex, s2.age, s2.id, s2.age)
}
