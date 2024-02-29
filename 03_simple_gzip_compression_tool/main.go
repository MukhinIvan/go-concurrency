// Ожидание завершения сопрограмм
package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	// Нет необходимости инициализировать WaitGroup
	var wg sync.WaitGroup
	// Так как переменная необходима за пределами цикла, она объявляется именно здесь
	var i int = -1
	var file string
	for i, file = range os.Args[1:] {
		// Для каждого файла сообщить группе, что ожидается выполнение еще одной операции сжатия
		wg.Add(1)
		// Эта функция вызывает функцию сжатия и уведомляет группу ожидания о ее завершении
		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)
	}
	// Внешняя сопрограмма (main) ожидает, пока все сопрограммы, выполняющие сжатие, вызовут w.Done
	wg.Wait()
	fmt.Printf("Compressed %d files\n", i+1)
}

func compress(filename string) error {
	// Открыть исходный файл для чтения
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	// Открыть файл архива с расширением .gz и именем исходного файла
	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	// Сжать данные и записать в соответствующий файл с помощью gzip.Writer
	gzout := gzip.NewWriter(out)
	// Функция io.Copy выполняет необходимое копирование
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}
