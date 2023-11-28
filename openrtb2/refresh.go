package openrtb2

import "encoding/json"

// Object: Refresh
type Refresh struct {

	// Attribute:
	//   refsettings
	// Type:
	//   object array; recommended
	// Description:
	//   A RefSettings object (see Section 3.2.34) describing the mechanics
	//   of how an ad placement automatically refreshes.
	RefSettings []RefSettings `json:"refsettings,omitempty"`

	// Attribute:
	//   count
	// Type:
	//   integer; recommended
	// Description:
	//   The number of times this ad slot had been refreshed since last page load.
	Count *int `json:"count,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Placeholder for vendor specific extensions to this object.
	Ext json.RawMessage `json:"ext,omitempty"`
}
