package openrtb2

import (
	"encoding/json"

	"github.com/prebid/openrtb/v20/adcom1"
)

// Object: RefSettings
//
// Information on how often and what triggers an ad slot being refreshed.
type RefSettings struct {

	// Attribute:
	//   reftype
	// Type:
	//   integer; default 0; recommended
	// Description:
	//   The type of the declared auto refresh.
	//   Refer to List: Auto Refresh Triggers in AdCOM 1.0
	RefType adcom1.AutoRefreshTrigger `json:"reftype,omitempty"`

	// Attribute:
	//   minint
	// Type:
	//   integer; recommended
	// Description:
	//   The minimum refresh interval in seconds. This applies to all
	//   refresh types. This is the (uninterrupted) time the ad creative
	//   will be rendered before refreshing to the next creative. If
	//   the field is absent, the exposure time is unknown. This field
	//   does not account for viewability or external factors such as a
	//   user leaving a page.
	MinInt int `json:"minint,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Placeholder for vendor specific extensions to this object.
	Ext json.RawMessage `json:"ext,omitempty"`
}
