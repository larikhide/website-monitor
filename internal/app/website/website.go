package website

import "time"

type Website struct {
	URL       string
	Status    bool
	LastCheck time.Time
}
