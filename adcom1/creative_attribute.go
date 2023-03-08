package adcom1

import (
	"database/sql/driver"
	"github.com/pinpr/backend-shared-kit/sql"
	"github.com/pinpr/backend-shared-kit/vld"
	"golang.org/x/exp/slices"
)

// CreativeAttribute specifies a standard list of creative attributes that can describe an actual ad or restrictions relative to a given placement.
type (
	CreativeAttribute     int
	CreativeAttributeEnum string
	CreativeAttributes    []CreativeAttributeEnum
)

// Standard list of creative attributes that can describe an actual ad or restrictions relative to a given placement.
//
// Values of 500+ hold vendor-specific codes.
const (
	AttrAudioAuto              CreativeAttributeEnum = "audioAuto"              // 1 Audio Ad (Autoplay)
	AttrAudioUser              CreativeAttributeEnum = "audioUser"              // 2 Audio Ad (User Initiated)
	AttrExpandableAuto         CreativeAttributeEnum = "expandableAuto"         // 3 Expandable (Automatic)
	AttrExpandableUserClick    CreativeAttributeEnum = "expandableUserClick"    // 4 Expandable (User Initiated - Click)
	AttrExpandableUserRollover CreativeAttributeEnum = "expandableUserRollover" // 5 Expandable (User Initiated - Rollover)
	AttrVideoAuto              CreativeAttributeEnum = "videoAutoplay"          // 6 In-Banner Video Ad (Autoplay)
	AttrVideoUser              CreativeAttributeEnum = "videoUserInitiated"     // 7 In-Banner Video Ad (User Initiated)
	AttrPop                    CreativeAttributeEnum = "pop"                    // 8 Pop (e.g., Over, Under, or Upon Exit)
	AttrProvocative            CreativeAttributeEnum = "provocative"            // 9 Provocative or Suggestive Imagery
	AttrExtremeAnimation       CreativeAttributeEnum = "extremeAnimation"       // 10 Shaky, Flashing, Flickering, Extreme Animation, Smileys
	AttrSurvey                 CreativeAttributeEnum = "survey"                 // 11 Surveys
	AttrTextOnly               CreativeAttributeEnum = "textOnly"               // 12 Text Only
	AttrInteractive            CreativeAttributeEnum = "interactive"            // 13 User Interactive (e.g., Embedded Games)
	AttrWindowsDialog          CreativeAttributeEnum = "windowsDialog"          // 14 Windows Dialog or Alert Style
	AttrHasAudioToggleButton   CreativeAttributeEnum = "hasAudioToggleButton"   // 15 Has Audio On/Off Button
	AttrHasSkipButton          CreativeAttributeEnum = "hasSkipButton"          // 16 Ad Provides Skip Button (e.g. VPAID-rendered skip button on pre-roll video)
	AttrFlash                  CreativeAttributeEnum = "flash"                  // 17 Adobe Flash
	AttrResponsive             CreativeAttributeEnum = "responsive"             // 18 Responsive; Sizeless; Fluid (i.e., creatives that dynamically resize to environment)
)

var allCreativeAttrs = []CreativeAttributeEnum{
	AttrAudioAuto,
	AttrAudioUser,
	AttrExpandableAuto,
	AttrExpandableUserClick,
	AttrExpandableUserRollover,
	AttrVideoAuto,
	AttrVideoUser,
	AttrPop,
	AttrProvocative,
	AttrExtremeAnimation,
	AttrSurvey,
	AttrTextOnly,
	AttrInteractive,
	AttrWindowsDialog,
	AttrHasAudioToggleButton,
	AttrHasSkipButton,
	AttrFlash,
	AttrResponsive,
}

func (e CreativeAttributeEnum) IsValid() error {
	return vld.EnumIsValid(&e, allCreativeAttrs)
}

func (e CreativeAttributeEnum) ToValue() CreativeAttribute {
	return CreativeAttribute(slices.Index(allCreativeAttrs, e) + 1)
}

func (e CreativeAttribute) ToValue() CreativeAttributeEnum {
	return allCreativeAttrs[e-1]
}

func (c *CreativeAttributes) Scan(src interface{}) error {
	return sql.Scan(c, src)
}

func (c CreativeAttributes) Value() (driver.Value, error) {
	return sql.Value(c)
}
