package main
import "fmt"
func main() {
	x:=5 //var x int = 5
	y:= "sachith"
	fmt.Println("hello world")
	if x>5{
		fmt.Println(x,"is larger than 5")
	} else{
		fmt.Println(x,"is not larger than 5")
	}
	fmt.Println("x:",x,"y:",y)

	
	for index := 0; index < 5; index++ {
		fmt.Print(index,",")
	}
	fmt.Println()

}