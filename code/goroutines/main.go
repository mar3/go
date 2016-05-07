package main
 
import "sync"
import "fmt"
import "time"
 
func exec(tasks *sync.WaitGroup, id int) {
	defer tasks.Done()
	fmt.Printf("%v <task:%x> running...\n", time.Now(), id)
	time.Sleep(time.Second * 5)
	fmt.Printf("%v <task:%x> Ok.\n", time.Now(), id)
}
 
func main() {
	var tasks sync.WaitGroup
	for i := 0; i < 3; i++ {
		tasks.Add(1)
		go exec(&tasks, i)
	}
	tasks.Wait()
	fmt.Printf("%v <main> Ok.\n", time.Now())
}