package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type TodoCounter struct {
	m *sync.Mutex
	counter int32
}

func (counter *TodoCounter) increment()  {
	counter.m.Lock()
	defer counter.m.Unlock()

	counter.counter++
}

func (counter *TodoCounter) get() int32 {
	counter.m.Lock()
	defer counter.m.Unlock()

	return counter.counter
}

func main() {
	todosToParse := 5

	//var wg sync.WaitGroup
	//wg.Add(todosToParse)

	counter := &TodoCounter{
		m:       &sync.Mutex{},
		counter: 0,
	}

	for i := 1; i <= todosToParse; i++ {
		go func(i int) {
			//defer wg.Done()
			url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", i)

			response, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}

			bodyBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			_ = response.Body.Close()

			fmt.Println("finished parsing todo #", i)
			fmt.Println(string(bodyBytes))
			counter.increment()

		}(i)
	}
	time.Sleep(time.Millisecond * 300)
	//wg.Wait()
	fmt.Println("finish", "counter value:", counter.get())
}
