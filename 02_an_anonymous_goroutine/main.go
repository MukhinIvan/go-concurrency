// Использование замыканий сопрограмм
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside a goroutine.")
	go func() { // объявление анонимной функции и вызов ее как сопрограммы
		fmt.Println("Inside a goroutine.")
	}()
	fmt.Println("Outside again.")

	runtime.Gosched() // обращение к планировщику
}
