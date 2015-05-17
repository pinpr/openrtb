package rtb

// 5.16 Location Type
//   The following table lists the options to indicate how the geographic information was determined.
const (
	GeoTypeGPS          uint8 = 1 // 1 GPS/Location Services
	GeoTypeIP           uint8 = 2 // 2 IP Address
	GeoTypeUserProvided uint8 = 3 // 3 User provided (e.g., registration data)
)

// 3.2.12 Object: Geo
//
// This object encapsulates various methods for specifying a geographic location. When subordinate to a
// Device object, it indicates the location of the device which can also be interpreted as the user’s current
// location. When subordinate to a User object, it indicates the location of the user’s home base (i.e., not
// necessarily their current location).
// The lat/lon attributes should only be passed if they conform to the accuracy depicted in the type
// attribute. For example, the centroid of a geographic region such as postal code should not be passed.
type Geo struct {

	// Attribute:
	//   lat
	// Type:
	//   float
	// Description:
	//   Latitude from -90.0 to +90.0, where negative is south.
	Lat float64 `json:"lat"`

	// Attribute:
	//   lon
	// Type:
	//   float
	// Description:
	//   Longitude from -180.0 to +180.0, where negative is west.
	Lon float64 `json:"lon"`

	// Attribute:
	//   type
	// Type:
	//   integer
	// Description:
	//   Source of location data; recommended when passing
	//   lat/lon. Refer to List 5.16.
	Type uint8 `json:"type"`

	// Attribute:
	//   country
	// Type:
	//   string
	// Description:
	//   Country code using ISO-3166-1-alpha-3.
	Country string `json:"country"`

	// Attribute:
	//   region
	// Type:
	//   string
	// Description:
	//   Region code using ISO-3166-2; 2-letter state code if USA.
	Region string `json:"region"`

	// Attribute:
	//   regionfips104
	// Type:
	//   string
	// Description:
	//   Region of a country using FIPS 10-4 notation. While OpenRTB
	//   supports this attribute, it has been withdrawn by NIST in 2008.
	RegionFIPS104 string `json:"regionfips104"`

	// Attribute:
	//   metro
	// Type:
	//   string
	// Description:
	//   Google metro code; similar to but not exactly Nielsen DMAs.
	//   See Appendix A for a link to the codes.
	Metro string `json:"metro"`

	// Attribute:
	//   city
	// Type:
	//   string
	// Description:
	//   City using United Nations Code for Trade & Transport
	//   Locations. See Appendix A for a link to the codes.
	City string `json:"city"`

	// Attribute:
	//   zip
	// Type:
	//   string
	// Description:
	//   Zip or postal code.
	ZIP string `json:"zip"`

	// Attribute:
	//   utcoffset
	// Type:
	//   integer
	// Description:
	//   Local time as the number +/- of minutes from UTC.
	UTCOffset int8 `json:"utcoffset"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext"`
}