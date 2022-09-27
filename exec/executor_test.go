package exec

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestExe(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "TEST",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pool := NewExecutor(5)
			for x := 0; x < 10; x++ {
				index := x
				pool.Exec(func() {
					num := rand.Intn(10)
					head := "[" + strconv.Itoa(index) + "]" + strconv.Itoa(num)
					log.Println(head + " Start")
					time.Sleep(time.Duration(num) * time.Second)
					log.Println(head + " Done")
				})
			}
			pool.Await()
			//time.Sleep(30 * time.Second)
			assert.Equal(t, tt.name, "TEST")
		})
	}
}

func TestX(t *testing.T) {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)

	}

}
func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
