package openrtb2

import (
	"encoding/json"

	"github.com/prebid/openrtb/v18/adcom1"
)

// Object: RefreshSettings
//
// Information on how often and what triggers an ad slot being refreshed.
type RefreshSettings struct {
	// Attribute:
	//   ref_type
	// Type:
	//   integer; default 0; recommended
	// Description:
	//   The type of the declared auto refresh.
	//   Refer to List: Auto Refresh Triggers in AdCOM 1.0
	RefType adcom1.AutoRefreshTrigger `json:"ref_type,omitempty"`

	// Attribute:
	//   count
	// Type:
	//   integer; recommended
	// Description:
	//   The minimum refresh interval in seconds.
	//   This applies to all refresh types.
	MinInt int `json:"min_int,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Placeholder for vendor specific extensions to this object.
	Ext json.RawMessage `json:"ext,omitempty"`
}
