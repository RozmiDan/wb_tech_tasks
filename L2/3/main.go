package main

import (
	"fmt"
	"os"
)

/*
Внутренне интерфейс в Go хранится как указатель на тип и указатель на данные. Интерфейс считается nil, только если оба этих поля равны nil.
Функция Foo возвращает значение типа *os.PathError, которое само по себе nil, но информация о типе *os.PathError остаётся в интерфейсе.
Поэтому интерфейсное значение имеет форму (*os.PathError, nil) и не считается nil, хотя внутри данных нет.
Пустой интерфейс — это interface{}, который может хранить значение любого типа.
*/

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)        // nil
	fmt.Println(err == nil) // false
}
