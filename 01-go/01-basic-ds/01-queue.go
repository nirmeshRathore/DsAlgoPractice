// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

type Pair struct {
    first int
    second int
}

type Queue[T any] struct {
    items []T
}

func (q* Queue[T])Enqueue(item T) {
    q.items = append(q.items, item)
}

func (q* Queue[T])Dequeue() (T, bool) {
    var zero T
    if (len(q.items) == 0) {
        return zero, false
    }
    
    item := q.items[0]
    q.items = q.items[1:]
    
    return item, true
}

func (q* Queue[T]) Front() (T) {
    return q.items[0]
}


func main() {
  fmt.Println("Try programiz.pro")
  
  q := Queue[Pair]{}
  
  q.Enqueue(Pair{first: 1, second: 2})
  q.Enqueue(Pair{first: 3, second: 4})
  
  fmt.Println(q);
  
  fmt.Println(q.Dequeue());
  fmt.Println(q.Front());
}
