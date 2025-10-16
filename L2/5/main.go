package main

/*
структура customError имплементирует интерфейс error
поэтому когда мы из ф-ии test() возвращаем nil-указатель типа customError, то присваиваем полю itab
структуры интерфейса тип customError, после чего переменная err типа интерфейса error становится не
nil. Поэтому в терминале будет вывод "error"
*/

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
