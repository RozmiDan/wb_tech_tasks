package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Обе горутины asChan отдают свои последовательности (1,3,5,7) и (2,4,6,8) с рандомными задержками, а merge с select неблокирующе выбирает готовый кейс.
Порядок значений внутри каждой горутины сохраняется, но уже в итоговом канале значения из двух каналов перемешаются случайно между собой
Когда один из входных каналов закрывается, мы присваиваем ему nil, тем самым отключая соответствующий кейс select; когда оба стали nil,
закрываем выходной канал c и выходим — поэтому чтение в main корректно завершается.
*/

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v, ok := <-a:
				if ok {
					c <- v
				} else {
					a = nil
				}
			case v, ok := <-b:
				if ok {
					c <- v
				} else {
					b = nil
				}
			}
			if a == nil && b == nil {
				close(c)
				return
			}
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Print(v)
	}
}
