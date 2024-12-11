package openrtb2

import (
	"encoding/json"

	"github.com/prebid/openrtb/v20/adcom1"
)

// 3.2.27 Object: EID
//
// Extended identifiers support in the OpenRTB specification allows buyers to use audience data in real-time bidding.
// This object can contain one or more UIDs from a single source or a technology provider.
// The exchange should ensure that business agreements allow for the sending of this data.
type EID struct {

	// Attribute:
	//   inserter
	// Type:
	//   string
	// Description:
	//   The canonical domain name of the entity (publisher, publisher monetization
	//   company, SSP, Exchange, Header Wrapper, etc.) that caused the ID array element
	//   to be added. This may be the operational domain of the system, if that is
	//   different from the parent corporate domain, to facilitate WHOIS and reverse IP
	//   lookups to establish clear ownership of the delegate system.
	//
	//   This should be the same value as used to identify sellers in an ads.txt file if
	//   one exists.
	//
	//   For ad tech intermediaries, this would be the domain as used in ads.txt. For
	//   publishers, this would match the domain in the `site` or `app` object.
	Inserter string `json:"inserter,omitempty"`

	// Attribute:
	//   matcher
	// Type:
	//   string
	// Description:
	//   Technology providing the match method as defined in mm.
	//
	//   In some cases, this may be the same value as inserter.
	//
	// When blank, it is assumed that the matcher is equal to the source
	//
	// May be omitted when mm=0, 1, or 2.
	Matcher string `json:"matcher,omitempty"`

	// Attribute:
	//   mm
	// Type:
	//   integer
	// Description:
	//   Match method used by the matcher. Refer to List: ID Match Methods
	//   in AdCOM 1.0.
	MM adcom1.MatchMethod `json:"mm,omitempty"`

	// Attribute:
	//   source
	// Type:
	//   string
	// Description:
	//   Canonical domain of the ID.
	Source string `json:"source,omitempty"`

	// Attribute:
	//   uids
	// Type:
	//   object array
	// Description:
	//   Array of extended ID UID objects from the given source. Refer
	//   to 3.2.28 Extended Identifier UIDs.
	UIDs []UID `json:"uids,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for advertising-system specific extensions to this object.
	Ext json.RawMessage `json:"ext,omitempty"`
}
