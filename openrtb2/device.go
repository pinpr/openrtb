package openrtb2

import (
	"encoding/json"

	"github.com/prebid/openrtb/v18/adcom1"
)

// 3.2.18 Object: Device
//
// This object provides information pertaining to the device through which the user is interacting.
// Device information includes its hardware, platform, location, and carrier data.
// The device can refer to a mobile handset, a desktop computer, set-top box, or other digital device.
//
// BEST PRACTICE: There are currently no prominent open source lists for device makes, models, operating systems, or carriers.
// Exchanges typically use commercial products or other proprietary lists for these attributes.
// Until suitable open standards are available, exchanges are highly encouraged to publish lists of their device make, model, operating system, and carrier values to bidders.
//
// BEST PRACTICE: Proper device IP detection in mobile is not straightforward.
// Typically, it involves starting at the left of the x-forwarded-for header, skipping private carrier networks (e.g., 10.x.x.x or 192.x.x.x), and possibly scanning for known carrier IP ranges.
// Exchanges are urged to research and implement this feature carefully when presenting device IP values to bidders.

type Device struct {
	// Attribute:
	//   geo
	// Type:
	//   object; recommended
	// Description:
	//   Location of the device assumed to be the user’s current
	//   location defined by a Geo object (Section 3.2.19).
	Geo *Geo `json:"geo,omitempty"`

	// Attribute:
	//   dnt
	// Type:
	//   integer; recommended
	// Description:
	//   Standard “Do Not Track” flag as set in the header by the
	//   browser, where 0 = tracking is unrestricted, 1 = do not track.
	DNT *int8 `json:"dnt,omitempty" validate:"omitempty,oneof=0 1"`

	// Attribute:
	//   lmt
	// Type:
	//   integer; recommended
	// Description:
	//   “Limit Ad Tracking” signal commercially endorsed (e.g., iOS,
	//   Android), where 0 = tracking is unrestricted, 1 = tracking must
	//   be limited per commercial guidelines.
	Lmt *int8 `json:"lmt,omitempty" validate:"omitempty,oneof=0 1"`

	// Attribute:
	//   ua
	// Type:
	//   string
	// Description:
	//   Browser user agent string. This field represents a raw user
	//   agent string from the browser. For backwards compatibility,
	//   exchanges are recommended to always populate ‘ua’ with the
	//   User-Agent string, when available from the end user’s device,
	//   even if an alternative representation, such as the User-Agent
	//   Client-Hints, is available and is used to populate ‘sua’. No
	//   inferred or approximated user agents are expected in this field.
	//   If a client supports User-Agent Client Hints, and ‘sua’ field is
	//   present, bidders are recommended to rely on ‘sua’ for
	//   detecting device type, browser type and version and other
	//   purposes that rely on the user agent information, and ignore
	//   ‘ua’ field. This is because the ‘ua’ may contain a frozen or
	//   reduced user agent string.
	UA string `json:"ua,omitempty"`

	// Attribute:
	//   sua
	// Type:
	//   object
	// Description:
	//   Structured user agent information defined by a UserAgent
	//   object (see Section 3.2.29). If both ‘ua’ and ‘sua’ are present in
	//   the bid request, ‘sua’ should be considered the more accurate
	//   representation of the device attributes. This is because the ‘ua’
	//   may contain a frozen or reduced user agent string.
	SUA *UserAgent `json:"sua,omitempty"`

	// Attribute:
	//   ip
	// Type:
	//   string; recommended
	// Description:
	//   IPv4 address closest to device.
	IP string `json:"ip,omitempty" validate:"omitempty,ip"`

	// Attribute:
	//   ipv6
	// Type:
	//   string
	// Description:
	//   IP address closest to device as IPV6.
	IPV6 string `json:"ipv6,omitempty"`

	// Attribute:
	//   devicetype
	// Type:
	//   integer
	// Description:
	//   The general type of device. Refer to List: Device Types in
	//   AdCOM 1.0.
	// Note:
	//   OpenRTB <=2.5 defined only types 1..7.
	DeviceType adcom1.DeviceType `json:"devicetype,omitempty"`

	// Attribute:
	//   make
	// Type:
	//   string
	// Description:
	//   Device make (e.g., “Apple”).
	Make string `json:"make,omitempty"`

	// Attribute:
	//   model
	// Type:
	//   string
	// Description:
	//   Device model (e.g., “iPhone”).
	Model string `json:"model,omitempty"`

	// Attribute:
	//   os
	// Type:
	//   string
	// Description:
	//   Device operating system (e.g., “iOS”).
	OS string `json:"os,omitempty"`

	// Attribute:
	//   osv
	// Type:
	//   string
	// Description:
	//   Device operating system version (e.g., “3.1.2”).
	OSV string `json:"osv,omitempty"`

	// Attribute:
	//   hwv
	// Type:
	//   string
	// Description:
	//   Hardware version of the device (e.g., “5S” for iPhone 5S).
	HWV string `json:"hwv,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer
	// Description:
	//   Physical height of the screen in pixels.
	H int64 `json:"h,omitempty"`

	// Attribute:
	//   w
	// Type:
	//   integer
	// Description:
	//   Physical width of the screen in pixels.
	W int64 `json:"w,omitempty"`

	// Attribute:
	//   ppi
	// Type:
	//   integer
	// Description:
	//   Screen size as pixels per linear inch.
	PPI int64 `json:"ppi,omitempty"`

	// Attribute:
	//   pxratio
	// Type:
	//   float
	// Description:
	//   The ratio of physical pixels to device independent pixels.
	PxRatio float64 `json:"pxratio,omitempty"`

	// Attribute:
	//   js
	// Type:
	//   integer
	// Description:
	//   Support for JavaScript, where 0 = no, 1 = yes.
	JS *int8 `json:"js,omitempty"`

	// Attribute:
	//   geofetch
	// Type:
	//   integer
	// Description:
	//   Indicates if the geolocation API will be available to JavaScript
	//   code running in the banner, where 0 = no, 1 = yes.
	GeoFetch *int8 `json:"geofetch,omitempty"`

	// Attribute:
	//   flashver
	// Type:
	//   string
	// Description:
	//   Version of Flash supported by the browser.
	FlashVer string `json:"flashver,omitempty"`

	// Attribute:
	//   language
	// Type:
	//   string
	// Description:
	//   Browser language using ISO-639-1-alpha-2.
	//   e.g. for the locale “zh-Hant-HK”, returns “zh”.
	Language string `json:"language,omitempty"`

	// Attribute:
	//   langb
	// Type:
	//   string
	// Description:
	//   Content language using IETF BCP 47. Only one of language or
	//   langb should be present. (e.g. "zh-HK" or “zh-Hant-HK”)
	LangB string `json:"langb,omitempty" validate:"omitempty,bcp47_language_tag"`

	// Attribute:
	//   carrier
	// Type:
	//   string
	// Description:
	//   Carrier or ISP (e.g., “VERIZON”) using exchange curated string
	//   names which should be published to bidders a priori.
	Carrier string `json:"carrier,omitempty"`

	// Attribute:
	//   mccmnc
	// Type:
	//   string
	// Description:
	//   Mobile carrier as the concatenated MCC-MNC code
	//   (e.g., "310-005" identifies Verizon Wireless CDMA in the
	//   USA).
	//   Refer to https://en.wikipedia.org/wiki/Mobile_country_code
	//   for further examples. Note that the dash between the MCC
	//   and MNC parts is required to remove parsing ambiguity. The
	//   MCC-MNC values represent the SIM installed on the device
	//   and do not change when a device is roaming. Roaming may
	//   be inferred by a combination of the MCC-MNC, geo, IP and
	//   other data signals.
	MCCMNC string `json:"mccmnc,omitempty"`

	// Carrier region code using ISO-3166-2; 2-letter state code if USA
	CarrierRegion string `json:"carrierregion,omitempty" validate:"omitempty,iso3166_1_alpha2"`

	// Attribute:
	//   connectiontype
	// Type:
	//   integer
	// Description:
	//   Network connection type. Refer to List: Connection Types in
	//   AdCOM 1.0.
	// Note:
	//   OpenRTB <=2.5 defined only connection types 1..6.
	ConnectionType *adcom1.ConnectionType `json:"connectiontype,omitempty"`

	// Attribute:
	//   ifa
	// Type:
	//   string
	// Description:
	//   ID sanctioned for advertiser use in the clear (i.e., not hashed).
	IFA string `json:"ifa,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`

	DeviceExtension
}

type DeviceExtension struct {
	Identifier
	DeviceInfo
}

type DeviceInfo struct {
	// The device name (e.g. somebody's iPhone), hashed by MD5
	Name string `json:"name,omitempty" validate:"omitempty,md5"`

	// Hardware of the device (e.g., “iPhone13,1” for iPhone 12 mini)
	HW string `json:"hw,omitempty"`

	// Hardware model of the device (e.g., “D52gAP” for iPhone 12 mini)
	HWM string `json:"hwm,omitempty"`

	// The total amount of physical memory(RAM) in bytes
	RAM uint64 `json:"ram,omitempty"`

	// The total amount of disk space in bytes
	Disk uint64 `json:"disk,omitempty"`

	// The available amount of disk space in bytes
	FreeDisk uint64 `json:"freedisk,omitempty"`

	// System boot time of the device in seconds since last reboot (use `string` instead of `double` type for precision)
	SysBootTime string `json:"sysboottime,omitempty"`

	// System update time of the device in seconds (use `string` instead of `double` type for precision)
	SysUptime string `json:"sysuptime,omitempty"`

	// Region code using ISO-3166-1-alpha-2 (e.g. for the locale “zh-Hant-HK”, returns “HK”)
	Region string `json:"region,omitempty" validate:"omitempty,iso3166_1_alpha2"`

	// The current difference in seconds between the receiver and GMT (Greenwich Mean Time)
	Timezone int32 `json:"timezone,omitempty"`
}
