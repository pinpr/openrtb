package adcom1

import (
	"github.com/pinpr/backend-shared-kit/vld"
	"golang.org/x/exp/slices"
)

// MediaRating represents media ratings used in describing content based on the TAG Inventory Quality Guidelines (IQG) v2.1 categorization.
// Refer to www.iab.com/guidelines/digital-video-suite for more information.
type MediaRating int8
type MediaRatingEnum string

// Media ratings used in describing content based on the TAG Inventory Quality Guidelines (IQG) v2.1 categorization.
const (
	MediaRatingAll    MediaRatingEnum = "all"    // 1 All Audiences
	MediaRatingOver12 MediaRatingEnum = "over12" // 2 Everyone Over Age 12
	MediaRatingMature MediaRatingEnum = "mature" // 3 Mature Audiences
)

var allMediaRatings = []MediaRatingEnum{
	MediaRatingAll,
	MediaRatingOver12,
	MediaRatingMature,
}

func (m MediaRatingEnum) IsValid() error {
	return vld.EnumIsValid(&m, allMediaRatings)
}

func (m MediaRatingEnum) ToValue() MediaRating {
	return MediaRating(slices.Index(allMediaRatings, m) + 1)
}

func (e MediaRating) ToValue() MediaRatingEnum {
	return allMediaRatings[e-1]
}
