package openrtb2

import "encoding/json"

// Object: Refresh
//
// The minimum exposure time, in seconds per view, that the (uninterrupted) ad creative will be rendered before refreshing to the next creative.
// If the field is absent, the exposure time is unknown.
// If 0, the exposure time is indefinite.
// This field does not account for viewability or external factors such as a user leaving a page.
type Refresh struct {
	// Attribute:
	//   count
	// Type:
	//   object array; recommended
	// Description:
	//   A RefreshSettings object (see Section 3.2.34) describing the
	//   mechanics of how an ad placement automatically refreshes.
	Settings []RefreshSettings `json:"settings,omitempty"`

	// Attribute:
	//   count
	// Type:
	//   integer; recommended
	// Description:
	//   The number of times this ad slot had been refreshed since last page load.
	Count int `json:"count,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Placeholder for vendor specific extensions to this object.
	Ext json.RawMessage `json:"ext,omitempty"`
}
