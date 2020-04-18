package stopwatch

import (
	"fmt"
	"time"
)

type _stopwatchStruct struct {
	time time.Time
}

func New() *_stopwatchStruct {
	watch := _stopwatchStruct{time: time.Now()}
	return &watch
}

func (self *_stopwatchStruct) Reset() {
	self.time = time.Now()
}

func (self *_stopwatchStruct) ToString() string {
	currentDuration := time.Since(self.time)
	millisec := currentDuration.Milliseconds()
	hours := millisec / 1000 / 60 / 60
	minutes := millisec / 1000 / 60
	secs := millisec / 1000
	return fmt.Sprintf("%02d:%02d:%02d.%03d", hours,
		minutes,
		secs,
		millisec%1000)
}
