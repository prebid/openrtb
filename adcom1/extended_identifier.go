package adcom1

import "encoding/json"

// ExtendedIdentifier support in the OpenRTB specification allows buyers to use audience data in real-time bidding.
// The exchange should ensure that business agreements allow for the sending of this data.
// Note, it is assumed that exchanges and DSPs will collaborate with the appropriate regulatory agencies and ID vendor(s) to ensure compliance.
type ExtendedIdentifier struct {
	// Attribute:
	//   source
	// Type:
	//   string
	// Definition:
	//   Source or technology provider responsible for the set of included IDs.
	//   Expressed as a top-level domain.
	Source string `json:"source,omitempty"`

	// Attribute:
	//   uids
	// Type:
	//   object array
	// Definition:
	//   Array of extended ID UID objects from the given source. Refer to Object: Extended Identifier UIDs.
	UIDs []ExtendedIdentifierUID `json:"uids,omitempty"`

	// Attribute:
	//   inserter
	// Type:
	//   string
	// Definition:
	//   The canonical domain name of the entity (publisher, publisher monetization company, SSP, Exchange, Header Wrapper, etc.) that caused the ID array element to be added.
	//   This should be the same value as used to identify sellers in an ads.txt file if one exists.
	//   For ad tech intermediaries, this would be the domain as used in ads.txt. For publishers, this would match the domain in the 'site' or 'app' object.
	Inserter string `json:"inserter,omitempty"`

	// Attribute:
	//   matcher
	// Type:
	//   string
	// Definition:
	//   Technology providing the match method as defined in 'mm'.
	//   In some cases, this may be the same value as inserter.
	//   When blank, it is assumed that the 'matcher' is equal to the 'source'.
	//   May be omitted when mm=0, 1, or 2.
	Matcher string `json:"matcher,omitempty"`

	// Attribute:
	//   mm
	// Type:
	//   integer
	// Definition:
	//   Match method used by the matcher. Refer to List: ID Match Methods
	MM MatchMethod `json:"mm,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
