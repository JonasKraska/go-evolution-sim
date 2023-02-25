package engine

import "time"

type Updater interface {
	Update(delta time.Duration)
}
