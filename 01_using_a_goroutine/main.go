// Использование go-подпрограмм для выполнения заданий
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	go echo(os.Stdin, os.Stdout) // вызов функции echo как go-подпрограммы
	time.Sleep(30 * time.Second) // 30-секундная пауза
	fmt.Println("Timed out.")    // вывод сообщения о завершении ожидания
	os.Exit(0)                   // выход из программы; при этом сопрограмма будет остановлена
}

func echo(in io.Reader, out io.Writer) { // функция echo является обычной функцией
	io.Copy(out, in) // io.Copy скопирует данные из os.Reader в os.Writer
}
