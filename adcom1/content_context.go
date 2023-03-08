package adcom1

import (
	"github.com/pinpr/backend-shared-kit/vld"
	"golang.org/x/exp/slices"
)

// ContentContext represents options for indicating the type of content being used or consumed by the user in which ads may appear.
// This table has values derived from the TAG Inventory Quality Guidelines (IQG).
type ContentContext int8
type ContentContextEnum string

// Options for indicating the type of content being used or consumed by the user in which ads may appear.
// This table has values derived from the TAG Inventory Quality Guidelines (IQG).
const (
	ContentVideo   ContentContextEnum = "video"   // 1 Video (i.e., video file or stream such as Internet TV broadcasts)
	ContentGame    ContentContextEnum = "game"    // 2 Game (i.e., an interactive software game)
	ContentMusic   ContentContextEnum = "music"   // 3 Music (i.e., audio file or stream such as Internet radio broadcasts)
	ContentApp     ContentContextEnum = "app"     // 4 Application (i.e., an interactive software application)
	ContentText    ContentContextEnum = "text"    // 5 Text (i.e., primarily textual document such as a web page, eBook, or news article)
	ContentOther   ContentContextEnum = "other"   // 6 Other (i.e., none of the other categories applies)
	ContentUnknown ContentContextEnum = "unknown" // 7 Unknown
)

var allContentContexts = []ContentContextEnum{
	ContentVideo,
	ContentGame,
	ContentMusic,
	ContentApp,
	ContentText,
	ContentOther,
	ContentUnknown,
}

func (c ContentContextEnum) IsValid() error {
	return vld.EnumIsValid(&c, allContentContexts)
}

func (c ContentContextEnum) ToValue() ContentContext {
	return ContentContext(slices.Index(allContentContexts, c) + 1)
}

func (e ContentContext) ToValue() ContentContextEnum {
	return allContentContexts[e-1]
}
