package openrtb2

type Identifier struct {

	// -------------------------------------------------------------------------------------------------------------------
	// Below IDs are for iOS platform

	// The Identifier for Advertisers (IDFA) is a random device identifier assigned by Apple to a user’s device.
	// Advertisers use this to track data, so they can deliver customized advertising.
	// The IDFA is used for tracking and identifying a user (without revealing personal information).
	IDFA string `json:"idfa,omitempty" validate:"omitempty,uppercase"`

	// CAID is a surveillance technology developed by the China Advertising Association to circumvent web tracking
	// restrictions set by Apple. CAID was developed by the state-supported.
	// CAID is believed to be an open framework based on device fingerprinting that uses a common API-based service,
	// initially operated by the China Advertising Association. to coordinate activities by multiple actors.
	// Data that can be used to create the fingerprint includes model, IP address, language and device start-up time.
	CAID string `json:"caid,omitempty" validate:"omitempty,uppercase"`

	// The identifier for vendors (IDFV) is a code assigned to all apps by one developer and
	// is shared across all apps by that developer on your device.
	// The value of the IDFV is the same for apps from the same developer running on the same device.
	IDFV string `json:"idfv,omitempty" validate:"omitempty,uppercase"`

	// -------------------------------------------------------------------------------------------------------------------
	// Below IDs are for Android platform

	// The Google advertising ID is a device identifier for advertisers that allows them to
	// anonymously track user ad activity on Android devices. It has often also been called the Android advertising ID
	GAID string `json:"gaid,omitempty" validate:"omitempty,uuid"`

	// An OAID is a non-permanent device identifier.
	// You can use the OAID to serve personalized ads to users while protecting their privacy.
	OAID string `json:"oaid,omitempty" validate:"omitempty,uuid"`

	// The Android unique device ID is called the Android Advertising ID (AAID - Android ID).
	// It’s an anonymous string of numbers and letters generated for the device upon initial setup.
	// The Android unique device ID is a combination of 8 digits followed by a dash and then three sets of 4 digits.
	// Here is an example of what an advertising ID on Android looks like: `9b9ba2ce17104d0`
	AAID string `json:"aaid,omitempty" validate:"omitempty,hexadecimal"`
}
