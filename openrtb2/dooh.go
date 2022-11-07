package openrtb2

import (
	"encoding/json"

	"github.com/prebid/openrtb/v17/adcom1"
)

// Object: DOOH
//
// This object should be included if the ad supported content is a Digital Out-Of-Home screen.
// A bid request with a DOOH object must not contain a site or app object.
// At a minimum, it is useful to provide id and/or venuetype, but this is not strictly required.
type DOOH struct {

	// Attribute:
	//   id
	// Type:
	//   string; recommended
	// Description:
	//   Exchange provided id for a placement or logical grouping of placements.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   App name (may be aliased at the publisher’s request).
	Name string `json:"name,omitempty"`

	// Attribute:
	//   venuetype
	// Type:
	//   string, array
	// Description:
	//   The type of out-of-home venue. The taxonomy to be used is defined by
	//   the venuetax field. If no venuetax field is supplied, The OpenOOH
	//   Venue Taxonomy is assumed.
	VenueType []string `json:"venuetype,omitempty"`

	// Attribute:
	//   venuetypetax
	// Type:
	//   integer; default 1
	// Description:
	//   The venue taxonomy in use. Refer to List: DOOH Venue Taxonomies
	VenueTypeTax *adcom1.DOOHVenueTaxonomy `json:"venuetypetax,omitempty"`

	// Attribute:
	//   publisher
	// Type:
	//   object
	// Description:
	//   Details about the publisher of the placement.
	Publisher *Publisher `json:"publisher,omitempty"`

	// Attribute:
	//   domain
	// Type:
	//   string
	// Description:
	//   Domain of the inventory (ads.txt) owner (e.g., “mysite.foo.com”)
	Domain string `json:"domain,omitempty"`

	// Attribute:
	//   keywords
	// Type:
	//   string
	// Description:
	//   Comma separated list of keywords about the DOOH placement.
	Keywords string `json:"keywords,omitempty"`

	// Attribute:
	//   content
	// Type:
	//   object
	// Description:
	//   Details about the Content within the DOOH placement.
	Content *Content `json:"content,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
