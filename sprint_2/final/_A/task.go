package main

// ID посылки: 66612281

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Deque двусторонняя очередь
type Deque struct {
	stack []int
	head,
	tail,
	maxSize,
	size int
}

// command команда действия
type command struct {
	action string
	num    int
}

const (
	ErrStackEmpty = "error"
	ErrStackMax   = "error"

	ActionPushBack  = "push_back"
	ActionPushFront = "push_front"
	ActionPopFront  = "pop_front"
	ActionPopBack   = "pop_back"
)

func main() {

	// сбор данных из input
	q, commands, err := getInputData()
	if err != nil {
		showError(err)
	}

	// решение
	result := q.Run(commands)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}

// Run выполнение алгоритма
func (q *Deque) Run(commands []command) []string {
	var (
		err    error
		result []string
	)
	for i := 0; i < len(commands); i++ {
		switch commands[i].action {
		case ActionPushBack:
			if err = q.pushBack(commands[i].num); err != nil {
				result = append(result, err.Error())
			}
			break
		case ActionPushFront:
			if err = q.pushFront(commands[i].num); err != nil {
				result = append(result, err.Error())
			}
		case ActionPopFront:
			result = append(result, q.popFront())
		case ActionPopBack:
			result = append(result, q.popBack())
		default:
			break
		}
	}

	return result
}

// pushBack добавить элемент в конец
func (q *Deque) pushBack(value int) error {
	if q.isMax() {
		return fmt.Errorf(ErrStackMax)
	}

	q.stack[q.getIndex(q.tail)] = value
	q.tail = (q.tail + 1) % q.maxSize
	q.size++

	return nil
}

// pushFront добавить элемент в начало
func (q *Deque) pushFront(value int) error {
	if q.isMax() {
		return fmt.Errorf(ErrStackMax)
	}

	q.stack[q.getIndex(q.head-1)] = value
	q.head = (q.head - 1) % q.maxSize
	q.size++

	return nil
}

// popBack вывести последний элемент и удалить его
func (q *Deque) popBack() string {
	if q.isEmpty() {
		return ErrStackEmpty
	}

	x := q.stack[q.getIndex(q.tail-1)]
	q.stack[q.getIndex(q.tail-1)] = 0

	q.tail = (q.tail - 1) % q.maxSize
	q.size--

	return fmt.Sprint(x)
}

// popFront вывести последний элемент и удалить его
func (q *Deque) popFront() string {
	if q.isEmpty() {
		return ErrStackEmpty
	}

	x := q.stack[q.getIndex(q.head)]
	q.stack[q.getIndex(q.head)] = 0

	q.head = (q.head + 1) % q.maxSize
	q.size--

	return fmt.Sprint(x)
}

// getIndex перевод "обратных" индексов
func (q *Deque) getIndex(index int) int {
	if index < 0 {
		return len(q.stack) + index
	}

	return index
}

// isEmpty определяет, пуст ли дек
func (q *Deque) isEmpty() bool {
	return q.size == 0
}

// isMax определяет, переполнен ли дек
func (q *Deque) isMax() bool {
	return q.size >= q.maxSize
}

// getInputData парсинг входных данных
func getInputData() (q *Deque, commands []command, err error) {

	input, err := getInputFromFile()
	if err != nil {
		showError(err)
	}
	// close file
	defer func(input *os.File) {
		_ = input.Close()
	}(input)

	reader := bufio.NewReader(input)

	var n, m int
	strNum, _, _ := reader.ReadLine()
	n, err = strconv.Atoi(string(strNum))
	if err != nil {
		return
	}

	strNum, _, _ = reader.ReadLine()
	m, err = strconv.Atoi(string(strNum))
	if err != nil {
		return
	}

	q = &Deque{
		stack:   make([]int, m),
		maxSize: m,
	}

	var com command
	commands = make([]command, 0, n)

	for i := 0; i < n; i++ {
		strNums, _ := reader.ReadString('\n')
		comStr := strings.Split(strings.TrimSpace(strNums), " ")

		com = command{
			action: comStr[0],
		}
		if len(comStr) == 2 {
			com.num, _ = strconv.Atoi(comStr[1])
		}

		commands = append(commands, com)
	}

	// clear bufio
	defer reader.Reset(reader)
	return
}

// getInputFromFile получение input из файла
func getInputFromFile() (*os.File, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	return file, nil
}

// showError обработка ошибки
func showError(err interface{}) {
	panic(err)
}
