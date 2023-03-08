package openrtb2

import (
	"github.com/pinpr/backend-shared-kit/vld"
	"golang.org/x/exp/slices"
)

// AuctionType defines an auction type.
type AuctionType int
type AuctionTypeEnum string

// AuctionType values.
const (
	AuctionFirstPrice      AuctionTypeEnum = "firstPrice"      // 1
	AuctionSecondPricePlus AuctionTypeEnum = "secondPricePlus" // 2
	AuctionDealPrice       AuctionTypeEnum = "dealPrice"       // 3
)

var allAuctionTypes = []AuctionTypeEnum{
	AuctionFirstPrice,
	AuctionSecondPricePlus,
	AuctionDealPrice,
}

func (e AuctionTypeEnum) IsValid() error {
	return vld.EnumIsValid(&e, allAuctionTypes)
}

func (e AuctionTypeEnum) ToValue() AuctionType {
	return AuctionType(slices.Index(allAuctionTypes, e) + 1)
}

func (e AuctionType) ToValue() AuctionTypeEnum {
	return allAuctionTypes[e-1]
}
