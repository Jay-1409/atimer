package timer

type Timer struct {
	Heaps        []*TimerHeap
	EventHandler *TimerEventHandler
}

func NewTimer(heapCount int, queueSize int, eventHandler *TimerEventHandler) *Timer {
	timer := &Timer{
		EventHandler: eventHandler,
	}
	for i := 0; i < heapCount; i = i + 1 {
		h := NewTimerHeap(i, queueSize)
		h.EventHandler = eventHandler
		timer.Heaps = append(timer.Heaps, h)
	}
	return timer
}
