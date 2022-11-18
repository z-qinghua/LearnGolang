package main
import "fmt"

func main(){

	//尾数部分可能造成精度损失
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=",num3,"num4=",num4)
}