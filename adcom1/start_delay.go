package adcom1

import (
	"github.com/pinpr/backend-shared-kit/vld"
	"golang.org/x/exp/slices"
)

// StartDelay represents video or audio start delay.
type StartDelay int8
type StartDelayEnum string

// Options for the video or audio start delay.
// If the start delay value is greater than 0, then the position is mid-roll and the value indicates the start delay.
const (
	StartPreRoll  StartDelayEnum = "preRoll"  // 0 Pre-Roll
	StartMidRoll  StartDelayEnum = "midRoll"  // -1 Generic Mid-Roll
	StartPostRoll StartDelayEnum = "postRoll" // -2 Generic Post-Roll
)

var allStartDelays = []StartDelayEnum{
	StartPreRoll,
	StartMidRoll,
	StartPostRoll,
}

func (e StartDelayEnum) IsValid() error {
	return vld.EnumIsValid(&e, allStartDelays)
}

func (e StartDelayEnum) ToValue() StartDelay {
	return StartDelay(-slices.Index(allStartDelays, e))
}

func (e StartDelay) ToValue() StartDelayEnum {
	return allStartDelays[-e]
}
