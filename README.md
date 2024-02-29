# Go-concurrency

## [Использование go-подпрограмм для выполнения заданий](https://github.com/MukhinIvan/go-concurrency/tree/main/01_using_a_goroutine/main.go)

**Сопрограммой** может быть любая функция, которая вызывается с помощью ключевого слова `go`.

В программе используются следующие библиотеки:

* `fmt` - для вывода строк в консоль (`fmt.Println()`);
* `io` - для интерфейсов чтения и записи (`io.Reader`, `io.Writer`), а также для копирования данных из одного потока в другой (`io.Copy()`);
* `os` - для стандартных потоков ввода-вывода (`os.Stdout` и `os.Stdin`);
* `time` - для задержки на 30 секунд (`time.Sleep()`).
