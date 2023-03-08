package adcom1

import (
	"github.com/pinpr/backend-shared-kit/vld"
	"golang.org/x/exp/slices"
)

// PlaybackCessationMode represents modes for when media playback terminates.
type PlaybackCessationMode int8
type PlaybackCessationModeEnum string

// Modes for when media playback terminates.
const (
	PlaybackCompletion      PlaybackCessationModeEnum = "completion"      // 1 On Video Completion or when Terminated by User
	PlaybackLeavingViewport PlaybackCessationModeEnum = "leavingViewport" // 2 On Leaving Viewport or when Terminated by User
	PlaybackFloating        PlaybackCessationModeEnum = "floating"        // 3 On Leaving Viewport Continues as a Floating/Slider Unit until Video Completion or when Terminated by User
)

var allPlaybackModes = []PlaybackCessationModeEnum{
	PlaybackCompletion,
	PlaybackLeavingViewport,
	PlaybackFloating,
}

func (e PlaybackCessationModeEnum) IsValid() error {
	return vld.EnumIsValid(&e, allPlaybackModes)
}

func (e PlaybackCessationModeEnum) ToValue() PlaybackCessationMode {
	return PlaybackCessationMode(slices.Index(allPlaybackModes, e) + 1)
}

func (e PlaybackCessationMode) ToValue() PlaybackCessationModeEnum {
	return allPlaybackModes[e-1]
}
