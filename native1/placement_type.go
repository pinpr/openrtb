package native1

import (
	"github.com/pinpr/backend-shared-kit/vld"
	"golang.org/x/exp/slices"
)

// 7.3 Placement Type IDs
//
// The FORMAT of the ad you are purchasing, separate from the surrounding context

type PlacementType int
type PlacementTypeEnum string

const (
	PlacementTypeFeed                 PlacementTypeEnum = "feed"                 // 1 In the feed of content - for example as an item inside the organic feed/grid/listing/carousel.
	PlacementTypeAtomicContentUnit    PlacementTypeEnum = "atomicContentUnit"    // 2 In the atomic unit of the content - IE in the article page or single image page
	PlacementTypeOutsideCoreContent   PlacementTypeEnum = "outsideCoreContent"   // 3 Outside the core content - for example in the ads section on the right rail, as a banner-style placement near the content, etc.
	PlacementTypeRecommendationWidget PlacementTypeEnum = "recommendationWidget" // 4 Recommendation widget, most commonly presented below the article content.

	// 500+ To be defined by the exchange
)

var allPlcmtTypes = []PlacementTypeEnum{
	PlacementTypeFeed,
	PlacementTypeAtomicContentUnit,
	PlacementTypeOutsideCoreContent,
	PlacementTypeRecommendationWidget,
}

func (p PlacementTypeEnum) IsValid() error {
	return vld.EnumIsValid(&p, allPlcmtTypes)
}

func (p PlacementTypeEnum) ToValue() PlacementType {
	return PlacementType(slices.Index(allPlcmtTypes, p) + 1)
}

func (e PlacementType) ToValue() PlacementTypeEnum {
	return allPlcmtTypes[e-1]
}
