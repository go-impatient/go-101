package main
import (
	"fmt"
)

func sayHello(name string) string {
	return "Hello," + name + "!"
}

func main() {
	message := sayHello("world")
	fmt.Println(message)
}
