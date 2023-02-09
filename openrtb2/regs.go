package openrtb2

import "encoding/json"

// Object: Regs
//
// This object contains any legal, governmental, or industry regulations that the sender deems applicable to the request.
// See Section 7.5 for more details on the flags supporting Coppa, GDPR and others.
type Regs struct {

	// Attribute:
	//   coppa
	// Type:
	//   integer
	// Description:
	//   Flag indicating if this request is subject to the COPPA
	//   regulations established by the USA FTC, where 0 = no, 1 = yes.
	//   Refer to Section 7.5 for more information.
	COPPA int8 `json:"coppa,omitempty"`

	// Attribute:
	//   gdpr
	// Type:
	//   integer
	// Description:
	//   Flag that indicates whether or not the request is subject to
	//   GDPR regulations 0 = No, 1 = Yes, omission indicates
	//   Unknown. Refer to Section 7.5 for more information.
	GDPR *int8 `json:"gdpr,omitempty"`

	// Attribute:
	//   us_privacy
	// Type:
	//   string
	// Description:
	//   Communicates signals regarding consumer privacy under US
	//   privacy regulation. See US Privacy String specifications. Refer
	//   to Section 7.5 for more information.
	USPrivacy string `json:"us_privacy,omitempty"`

	// Attribute:
	//   gpp
	// Type:
	//   string
	// Description:
	//   Contains the Global Privacy Platform’s consent string. See the
	//   Global Privacy Platform specification for more details.
	GPP string `json:"gpp,omitempty"`

	// Attribute:
	//   gpp_sid
	// Type:
	//   integer array
	// Description:
	//   Array of the section(s) of the string which should be applied for this
	//   transaction. Generally will contain one and only one value, but there
	//   are edge cases where more than one may apply. GPP Section 3 (Header)
	//   and 4 (Signal Integrity) do not need to be included. See the
	//   GPP Section Information for more details.
	GPPSID []int8 `json:"gpp_sid,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
