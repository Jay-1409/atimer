package timer

import (
	"time"
)

type TimerTask struct {
	ID          string
	FireAt      time.Time
	CallBackURL string
}
