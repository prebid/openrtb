package adcom1

import "encoding/json"

// Regs object contains any known legal, governmental, or industry regulations that are in effect.
type Regs struct {
	// Attribute:
	//   coppa
	// Type:
	//   integer
	// Definition:
	//   Flag indicating if COPPA regulations apply, where 0 = no, 1 = yes.
	//   The Children's Online Privacy Protection Act (COPPA) was established by the U.S. Federal Trade Commission.
	COPPA int8 `json:"coppa,omitempty"`

	// Attribute:
	//   gdpr
	// Type:
	//   integer
	// Definition:
	//   Flag indicating if GDPR regulations apply, where 0 = no, 1 = yes.
	//   The General Data Protection Regulation (GDPR) is a regulation of the European Union.
	GDPR int8 `json:"gdpr,omitempty"`

	// Attribute:
	//   gpp
	// Type:
	//   string
	// Definition:
	//   Contains the Global Privacy Platform’s consent string.
	//   See the Global Privacy Platform specification for more details.
	GPP string `json:"gpp,omitempty"`

	// Attribute:
	//   gpp_sid
	// Type:
	//   integer array
	// Definition:
	//   Array of the section(s) of the string which should be applied for this transaction.
	//   Generally will contain one and only one value, but there are edge cases where more than one may apply.
	//   GPP Section 3 (Header) and 4 (Signal Integrity) do not need to be included.
	//   See the GPP Section Information for more details.
	GPPSID []int8 `json:"gpp_sid,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
