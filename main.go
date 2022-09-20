package main

import (
	"fmt"
)
	

//golang中的指针类型
func main(){
	var num int = 10
	var ptr *int = &num
	fmt.Printf("num变量的地址为%v\n,指针指向的地址为%v\n",&num,ptr)
	*ptr = 20
	fmt.Printf("num变量的值为%d\n",num)


}