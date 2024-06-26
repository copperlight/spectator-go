package meter

import (
	"fmt"
	"github.com/Netflix/spectator-go/v2/spectator/writer"
	"time"
)

// Timer is used to measure how long (in seconds) some event is taking. This
// type is safe for concurrent use.
type Timer struct {
	id              *Id
	writer          writer.Writer
	meterTypeSymbol string
}

// NewTimer generates a new timer, using the provided meter identifier.
func NewTimer(id *Id, writer writer.Writer) *Timer {
	return &Timer{id, writer, "t"}
}

// MeterId returns the meter identifier.
func (t *Timer) MeterId() *Id {
	return t.id
}

// Record records the duration this specific event took.
func (t *Timer) Record(amount time.Duration) {
	if amount >= 0 {
		var line = fmt.Sprintf("%s:%s:%f", t.meterTypeSymbol, t.id.spectatordId, amount.Seconds())
		t.writer.Write(line)
	}
}
