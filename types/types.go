package types

import "time"

type Subtitle struct {
	Episode string
	Lines2  string
	Lines   []string
	StartAt time.Duration
	EndAt   time.Duration
}
