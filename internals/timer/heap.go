package timer

import (
	"container/heap"
	"sync"
	"time"
)

type TimerTaskHeap []*TimerTask

type TimerHeap struct {
	ID          int
	Tasks       TimerTaskHeap
	MaxTaskSize int
	mu          sync.Mutex
}

func NewTimerHeap(id int, queueSize int) *TimerHeap {
	t := &TimerHeap{
		ID:          id,
		Tasks:       make(TimerTaskHeap, 0, queueSize),
		MaxTaskSize: queueSize,
	}

	heap.Init(&t.Tasks)
	return t
}

func (h TimerTaskHeap) Len() int {
	return len(h)
}

func (h TimerTaskHeap) Less(i, j int) bool {
	return h[i].FireAt.Before(h[j].FireAt)
}

func (h TimerTaskHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *TimerTaskHeap) Push(x any) {
	*h = append(*h, x.(*TimerTask))
}

func (h *TimerTaskHeap) Pop() any {
	old := *h
	n := len(old)

	task := old[n-1]
	*h = old[:n-1]

	return task
}

func (h TimerTaskHeap) Peek() *TimerTask {
	if len(h) == 0 {
		return nil
	}
	return h[0]
}

func (t *TimerHeap) AddTask(task *TimerTask) bool {
	if t.Tasks.Len() >= t.MaxTaskSize {
		return false
	}

	heap.Push(&t.Tasks, task)
	return true
}

func (t *TimerHeap) PeekTask() *TimerTask {
	return t.Tasks.Peek()
}

func (t *TimerHeap) PopTask() *TimerTask {
	if t.Tasks.Len() == 0 {
		return nil
	}

	return heap.Pop(&t.Tasks).(*TimerTask)
}

/*
*

	Wait until the top most task reaches its completion state
	TODO: Implement something better than just polling.

	The reason for this busy wait is, that if a timer task hits us, and it's to be fired before the next fire candidate in our queue
	in that case, if we are to go into sleep we would delay the firing of this new task.

	We need something smarter than just random sleep time.
*/
func (t *TimerHeap) Run() {
	for {
		next := t.PeekTask()

		if next == nil {
			// the heap is empty
			continue
		}
		for time.Until(next.FireAt) > 0 {
			// busy wait
		}
		t.FireExpired()
	}
}

/*
*

	While we start firing one task it is possible that there are multiple timers that exist at small timer intervals
	Lets call this as TASK FLOODING. The synchronous nature of this function allows us to fire all of such timers before returning to the heap.

	The reason for taking locks on popping one task is to avoid the blocking of adding new tasks to the heap during a TASK FLOOD.
*/
func (t *TimerHeap) FireExpired() {
	now := time.Now()
	for {
		t.mu.Lock()
		top := t.PeekTask()
		if top == nil || top.FireAt.After(now) {
			t.mu.Unlock()
			return
		}
		expired := t.PopTask()
		t.mu.Unlock()
		println("fired task", expired.ID, "from heap: ", t.ID)
	}
}
