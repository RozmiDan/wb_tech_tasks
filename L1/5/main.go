package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Generator(stopCh <-chan time.Time) <-chan int {
	resCh := make(chan int)
	tk := time.NewTicker(100 * time.Millisecond)

	go func() {
		defer func() {
			close(resCh)
			tk.Stop()
		}()
		for {
			select {
			case <-stopCh:
				fmt.Println("time deadline, stop generator")
				return
			case <-tk.C:
				select {
				case resCh <- rand.Intn(100):
				case <-stopCh:
					fmt.Println("time deadline, stop generator")
					return
				}
			}
		}
	}()

	return resCh
}

func main() {
	fmt.Print("Enter N seconds: ")

	sc := bufio.NewScanner(os.Stdin)
	if !sc.Scan() {
		if err := sc.Err(); err != nil {
			log.Fatal("scan error: ", err)
		}
		log.Fatal("no input")
	}

	N, err := strconv.Atoi(sc.Text())
	if err != nil || N < 0 {
		log.Fatal("scan error")
	}

	deadline := time.After(time.Second * time.Duration(N))

	resultCh := Generator(deadline)

	for val := range resultCh {
		fmt.Println("got value: ", val)
	}
}
