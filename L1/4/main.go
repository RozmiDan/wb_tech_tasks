package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
Я использовал context.Context, чтобы при получении SIGINT через cancel() все горутины получили \
сигнал завершения. Это удобно, так как контекст — стандартный механизм Go для управления временем
жизни горутин, позволяющий централизованно и безопасно останавливать их работу.
*/

func Generator(ctx context.Context, cap int) chan int {
	outpChan := make(chan int, cap)

	go func() {
		defer close(outpChan)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Millisecond * 200)
				select {
				case <-ctx.Done():
					return
				default:
					outpChan <- rand.Intn(1000)
				}

			}
		}
	}()

	return outpChan
}

func startWorkers(ctx context.Context, wg *sync.WaitGroup, count int, inputCh <-chan int) {

	for i := 0; i < count; i++ {

		go func(num int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("worker %d stoping\n", num)
					return
				case val, ok := <-inputCh:
					select {
					case <-ctx.Done():
						fmt.Printf("worker %d stoping\n", num)
						return
					default:
						if !ok {
							return
						}
						fmt.Printf("worker %d got some number: %d\n", num, val)
					}
				}
			}
		}(i)
	}
}

func main() {
	countWorkers := 10

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	outCh := Generator(ctx, countWorkers)

	wg.Add(countWorkers)

	startWorkers(ctx, wg, countWorkers, outCh)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT)

	<-stop

	cancel()

	wg.Wait()

	fmt.Println("All goroutine stoped")
}
