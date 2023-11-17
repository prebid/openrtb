package openrtb2

import (
	"encoding/json"
)

// Object: DurFloors
//
// This object allows sellers to specify price floors for video and audio creatives, whose price varies based on time.
// For example: 1-15 seconds at a floor of $5; 16-30 seconds at a floor of $10, > 31 seconds at a floor of $20.
// There are no explicit constraints on the defined ranges, nor guarantees that they don't overlap.
// In cases where multiple ranges may apply, it is up to the buyer and seller to coordinate on which floor is applicable.
type DurFloors struct {

	// Attribute:
	//   mindur
	// Type:
	//   integer
	// Description:
	//   An integer indicating the low end of a duration range. If this
	//   value is missing, the low end is unbounded. Either mindur or maxdur
	//   is required, but not both.
	MinDur int64 `json:"mindur,omitempty"`

	// Attribute:
	//   maxdur
	// Type:
	//   integer
	// Description:
	//   An integer indicating the high end of a duration range. If this
	//   value is missing, the high end is unbounded. Either mindur or maxdur
	//   is required, but not both.
	MaxDur int64 `json:"maxdur,omitempty"`

	// Attribute:
	//   bidfloor
	// Type:
	//   float; default 0
	// Description:
	//   Minimum bid for a given impression opportunity, if bidding with a
	//   creative in this duration range, expressed in CPM. For any creatives
	//   whose durations are outside of the defined min/max, the bidfloor at
	//   the Imp level will serve as the default floor.
	BidFloor float64 `json:"bidfloor,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Placeholder for vendor specific extensions to this object.
	Ext json.RawMessage `json:"ext,omitempty"`
}
