package openrtb2

import "encoding/json"

// 3.2.12 Object: Deal
//
// This object constitutes a specific deal that was struck a priori between a buyer and a seller.
// Its presence with the Pmp collection indicates that this impression is available under the terms of that deal.
// Refer to Section 7.3 for more details.
type Deal struct {

	// Attribute:
	//   id
	// Type:
	//   string; required
	// Description:
	//   A unique identifier for the direct deal.
	ID string `json:"id"`

	// Attribute:
	//   bidfloor
	// Type:
	//   float; default 0
	// Description:
	//   Minimum bid for this impression expressed in CPM.
	BidFloor float64 `json:"bidfloor,omitempty"`

	// Attribute:
	//   bidfloorcur
	// Type:
	//   string; default ”USD”
	// Description:
	//   Currency specified using ISO-4217 alpha codes. This may be
	//   different from bid currency returned by bidder if this is
	//   allowed by the exchange. This field does not inherit from
	//   imp.bidfloorcur; it is either explicitly specified or
	//   defaults to USD.
	BidFloorCur string `json:"bidfloorcur,omitempty"`

	// Attribute:
	//   at
	// Type:
	//   integer
	// Description:
	//   Optional override of the overall auction type of the bid
	//   request, where 1 = First Price, 2 = Second Price Plus, 3 = the
	//   value passed in bidfloor is the agreed upon deal price.
	//   Additional auction types can be defined by the exchange.
	AT int64 `json:"at,omitempty"`

	// Attribute:
	//   wseat
	// Type:
	//   string array
	// Description:
	//   Whitelist of buyer seats (e.g., advertisers, agencies) allowed to
	//   bid on this deal. IDs of seats and the buyer’s customers to
	//   which they refer must be coordinated between bidders and
	//   the exchange a priori. Omission implies no seat restrictions.
	WSeat []string `json:"wseat,omitempty"`

	// Attribute:
	//   wadomain
	// Type:
	//   string array
	// Description:
	//   Array of advertiser domains (e.g., advertiser.com) allowed to
	//   bid on this deal. Omission implies no advertiser restrictions.
	WADomain []string `json:"wadomain,omitempty"`

	// Attribute:
	//   guar
	// Type:
	//   integer, default 0
	// Description:
	//   Indicates that the deal is of type guaranteed and the bidder must
	//   bid on the deal, where 0 = not a guaranteed deal, 1 = guaranteed deal.
	Guar int8 `json:"guar,omitempty"`

	// Attribute:
	//   mincpmpersec
	// Type:
	//   float
	// Definition:
	//   Minimum CPM per second. This is a price floor for video or audio
	//   impression opportunities, relative to the duration of bids an
	//   advertiser may submit.
	MinCPMPerSec float64 `json:"mincpmpersec,omitempty"`

	// Attribute:
	//   durfloors
	// Type:
	//   object array
	// Description:
	//   Container for floor price by duration information, to be used if a
	//   given deal is eligible for video or audio demand. An array of DurFloors
	//   objects (see Section 3.2.35).
	DurFloors []DurFloors `json:"durfloors,omitempty"`

	// Attribute:
	//   wadomain
	// Type:
	//   string array
	// Description:
	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
