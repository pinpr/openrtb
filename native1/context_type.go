package native1

import (
	"github.com/pinpr/backend-shared-kit/vld"
	"golang.org/x/exp/slices"
)

// 7.1 Context Type IDs
//
// The context in which the ad appears - what type of content is surrounding the ad on the
// page at a high level.
// This maps directly to the new Deep Dive on In-Feed Ad Units. This
// denotes the primary context, but does not imply other content may not exist on the
// page - for example it's expected that most content platforms have some social
// components, etc.

type ContextType int
type ContextTypeEnum string

const (
	ContextTypeContent ContextTypeEnum = "content" // 1 Content-centric context such as newsfeed, article, image gallery, video gallery, or similar.
	ContextTypeSocial  ContextTypeEnum = "social"  // 2 Social-centric context such as social network feed, email, chat, or similar.
	ContextTypeProduct ContextTypeEnum = "product" // 3 Product context such as product listings, details, recommendations, reviews, or similar.

	// 500+ To be defined by the exchange.
)

var allContextTypes = []ContextTypeEnum{
	ContextTypeContent,
	ContextTypeSocial,
	ContextTypeProduct,
}

func (c ContextTypeEnum) IsValid() error {
	return vld.EnumIsValid(&c, allContextTypes)
}

func (c ContextTypeEnum) ToValue() ContextType {
	return ContextType(slices.Index(allContextTypes, c) + 1)
}

func (e ContextType) ToValue() ContextTypeEnum {
	return allContextTypes[e-1]
}
