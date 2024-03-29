# Go-concurrency

По книге Go in Practice, Matt Butcher, Matt Farina

## [Использование go-подпрограмм для выполнения заданий](https://github.com/MukhinIvan/go-concurrency/tree/main/01_using_a_goroutine/main.go)

**Сопрограммой** может быть любая функция, которая вызывается с помощью ключевого слова `go`.

В программе используются следующие библиотеки:

* `fmt` - для вывода строк в консоль (`fmt.Println()`);
* `io` - для интерфейсов чтения и записи (`io.Reader`, `io.Writer`), а также для копирования данных из одного потока в другой (`io.Copy()`);
* `os` - для стандартных потоков ввода-вывода (`os.Stdout` и `os.Stdin`);
* `time` - для задержки на 30 секунд (`time.Sleep()`).

## [Использование замыканий сопрограмм](https://github.com/MukhinIvan/go-concurrency/tree/main/02_an_anonymous_goroutine/main.go)

Объявляем анонимную функцию и вызываем ее как сопрограмму.

В программе используются следующие библиотеки:

* `fmt` - для вывода строк в консоль (`fmt.Println()`);
* `runtime` - для `runtime.Gosched()`.

Чтобы у планировщика была возможность запустить сопрограмму до того, как функция `main` завершит программу, мы вызываем `runtime.Gosched()`, давая среде выполнения шанс запустить сопрограмму до завершения программы.

Если сопрограмма выполняет запрос к базе данных, простого вызова `runtime.Gosched` может быть недостаточно, чтобы обеспечить завершение запросов другими сопрограммами. Они могут находиться в состоянии ожидания ответа от базы данных, что вынудит планировщик продолжить выполнение текущей функции.

## [Ожидание завершения сопрограмм](https://github.com/MukhinIvan/go-different-programs/tree/main/03_simple_gzip_compression_tool/main.go)

Реализация простого инструмента сжатия произвольного количества отдельных файлов с использованием встроенной библиотеки `Gzip` (`compress/gzip`).

```
go run main.go file1.txt file2.txt
```

Иногда требуется вызвать несколько сопрограмм и дождаться, когда они все завершать работу. Проще всего достичь нужного результата позволяет формирование **группы ожидания**. Группа ожидания - это механизм передачи сообщений, оповещающих ожидающую сопрограмму, что она может продолжить работу.

Используем функцию `sync.WaitGroup`, чтобы известить внешний процесс, что все сопрограммы завершились и можно продолжить выполнение.

В программе используются следующие библиотеки:

* `compress/gzip` - для сжатия данных и записи их в файл (`gzip.Writer`);
* `fmt` - для вывода строк в консоль (`fmt.Printf()`);
* `io` - для копирования данных из одного потока в другой (`io.Copy()`);
* `os` - для открытия и создания файлов (`os.Open()`, `os.Create()`);
* `sync` - для создания группы ожидания (`sync.WaitGroup`).

## [Блокировка с помощью мьютексов](https://github.com/MukhinIvan/go-different-programs/tree/main/04_locking_with_a_mutex/main.go)

Эта программа читает файлы, указанные в аргументах командной строки, и подсчитывает число вхождений каждого найденного в них слова. В завершение она выводит список слов, встречающихся более одного раза.

```
go run main.go file1.txt file2.txt
```

```
Words that appear more than once:
go: 2
am: 6
brother: 2
every: 3
friends: 2
i: 18
a: 5
to: 4
twelve: 2
is: 6
always: 2
my: 10
love: 2
and: 8
at: 2
as: 2
have: 3
favourite: 2
school: 3
```

Всякий раз, когда две или более сопрограмм работают с одним и тем же фрагментом данных и эти данные могут изменяться, возникает вероятность появления **состояния гонки**.

Один из способов избежать состояния гонки - дать каждой сопрограмме возможность "заблокировать" ресурс перед использованием и разблокировать после выполнения операций. Для блокировки и разблокировки объекта используется функция `sync.Mutex`.

Метод `add` блокирует объект `map`, изменяет его и затем разблокирует. Это предотвращает одновременное изменение `map`-объекта несколькими сопрограммами.

В программе используются следующие библиотеки:

* `bufio` - для `NewScanner`;
* `fmt` - для `fmt.Printf()` и `fmt.Println()`;
* `os` - для `os.Args` и `os.Open`;
* `strings` - для `strings.ToLower()`;
* `sync` - `sync.WaitGroup` и `sync.Mutex`.
