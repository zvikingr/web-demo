package uuid

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// NewUUID create uuid, eg: b152a6f0-0bfb-11ed-b06f-acde48001122
// If an error occurs, use the timestamp instead
func NewUUID() string {
	x, err := uuid.NewUUID()
	if err != nil {
		return fmt.Sprintf("%d", time.Now().Nanosecond())
	}
	return x.String()
}
