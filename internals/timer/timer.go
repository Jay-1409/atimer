package timer

type Timer struct {
	Heaps []*TimerHeap
}

func NewTimer(heapCount int, queueSize int) *Timer {
	timer := &Timer{}
	for i := 0; i < heapCount; i = i + 1 {
		h := NewTimerHeap(i, queueSize)
		timer.Heaps = append(timer.Heaps, h)
	}
	return timer
}

func (t *Timer) Start() {
	for _, h := range t.Heaps {
		go h.Run()
	}
}
