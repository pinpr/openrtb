package openrtb

// 3.2.18 Object: Deal
//
// This object constitutes a specific deal that was struck a priori between a buyer and a seller. Its presence
// with the Pmp collection indicates that this impression is available under the terms of that deal. Refer to
// Section 7.2 for more details.
type Deal struct {

	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   A unique identifier for the direct deal.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   bidfloor
	// Type:
	//   float
	// Description:
	//   Minimum bid for this impression expressed in CPM.
	BidFloor float64 `json:"bidfloor,omitempty"`

	// Attribute:
	//   bidfloorcur
	// Type:
	//   string
	// Description:
	//   Currency specified using ISO-4217 alpha codes. This may be different from bid currency returned
	//   by bidder if this is allowed by the exchange.
	BidFloorCur string `json:"bidfloorcur,omitempty"`

	// Attribute:
	//   at
	// Type:
	//   integer
	// Description:
	//   Optional override of the overall auction type of the bid request, where 1 = First Price,
	//   2 = Second Price Plus, 3 = the value passed in bidfloor is the agreed upon deal price. Additional
	//   auction types can be defined by the exchange.
	AT int8 `json:"at,omitempty"`

	// Attribute:
	//   wseat
	// Type:
	//   string array
	// Description:
	//   Whitelist of buyer seats allowed to bid on this deal.  Seat IDs must be
	//   communicated between bidders and the exchange a priori. Omission implies no seat restrictions.
	WSeat []string `json:"wseat,omitempty"`

	// Attribute:
	//   wadomain
	// Type:
	//   string array
	// Description:
	//   Array of advertiser domains (e.g., advertiser.com) allowed to bid on this deal. Omission implies
	//   no advertiser restrictions.
	WADomain []string `json:"wadomain,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext RawJSON `json:"ext,omitempty"`
}
