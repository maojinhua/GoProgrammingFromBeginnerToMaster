package main

import "fmt"

type BinaryAdder interface {
	Add(int, int) int
}

type MyAdderFunc func(int, int) int

func (f MyAdderFunc) Add(x, y int) int {
	return f(x, y)
}

func MyAdd(x, y int) int {
	return x + y
}

func main() {
	var i BinaryAdder = MyAdderFunc(MyAdd)
	fmt.Println(i.Add(5, 6))

	var d BinaryDeleter = MyDeleterFunc(MyDelete)
	fmt.Println(d.Delete(5,6))

	var d2 BinaryDeleter = MyDeleterFunc(MyDelete2)
	fmt.Println(d2.Delete(5,6))
}

type BinaryDeleter interface{
	Delete(int,int) int
}

type MyDeleterFunc func(int,int)int

func (f MyDeleterFunc) Delete(x,y int)int{
	return f(x,y)
}

func MyDelete(x,y int)int{
	return x-y
}

func MyDelete2(x,y int)int{
	return 2*x-y
}