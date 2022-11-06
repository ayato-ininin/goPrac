package mylib

import "fmt"

//これ全部大文字じゃないと、main.goから読み込めない。public
type Person struct{
	// Name
	Name string
	// Age
	Age int
}

//小文字にするとprivate。ほかから呼べない。
type person struct{
	name string
	age int
}

//ここも大文字にしないとほかから呼べない。public
func Say()  {
	fmt.Println("I'm a perfect human")
}
