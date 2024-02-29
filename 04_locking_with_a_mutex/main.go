// Счетчик слов
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	// Используется группа ожидания для мониторинга группы сопрограмм
	var wg sync.WaitGroup
	w := newWords()
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	// Перед завершением программы вывести результаты поиска
	fmt.Println("Words that appear more than once:")
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
}

// Извлекаемые слова помещаются в структуру
type words struct {
	sync.Mutex // наследовать мьютекс
	found      map[string]int
}

// Создание нового экземляра класса
func newWords() *words {
	return &words{found: map[string]int{}}
}

// Фиксирует количество вхождений этого слова
func (w *words) add(word string, n int) {
	// Заблокировать объект, изменить и разблокировать
	w.Lock()
	defer w.Unlock()

	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

// Открытие файла, анализ содержимого и подсчет найденных в нем слов
func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}
