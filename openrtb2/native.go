package openrtb2

import (
	"encoding/json"

	"github.com/prebid/openrtb/v20/adcom1"
)

// 3.2.9 Object: Native
//
// This object represents a native type impression.
// Native ad units are intended to blend seamlessly into the surrounding content (e.g., a sponsored Twitter or Facebook post).
// As such, the response must be well-structured to afford the publisher fine-grained control over rendering.
//
// The Native Subcommittee has developed a companion specification to OpenRTB called the Dynamic Native Ads API.
// It defines the request parameters and response markup structure of native ad units.
// This object provides the means of transporting request parameters as an opaque string so that the specific parameters can evolve separately under the auspices of the Dynamic Native Ads API.
// Similarly, the ad markup served will be structured according to that specification.
//
// The presence of a Native as a subordinate of the Imp object indicates that this impression is offered as a native type impression.
// At the publisher’s discretion, that same impression may also be offered as banner, video, and/or audio by also including as Imp subordinates objects of those types.
// However, any given bid for the impression must conform to one of the offered types.
type Native struct {

	// Attribute:
	//   request
	// Type:
	//   string; required
	// Description:
	//   Request payload complying with the Native Ad Specification.
	//   The root node of the payload, “native”, was dropped in the
	//   Native Ad Specification 1.1.
	//
	//   For Native 1.0, this is a JSON-encoded string consisting of a
	//   unnamed root object with a single subordinate object named
	//   'native', which is the Native Markup Request object, section 4.1
	//   of OpenRTB Native 1.0 specification.
	//
	//   For Native 1.1 and higher, this is a JSON-encoded string
	//   consisting of an unnamed root object which is itself the Native
	//   Markup Request Object, section 4.1 of OpenRTB Native 1.1+.
	Request string `json:"request"`

	// RequestObj is used for the case when we receive not string-typed data in request field
	// In custom marshalers we cast object in string if it's possible
	RequestObj interface{} `json:"-"`

	// Attribute:
	//   ver
	// Type:
	//   string; recommended
	// Description:
	//   Version of the Dynamic Native Ads API to which request
	//   complies; highly recommended for efficient parsing.
	Ver string `json:"ver,omitempty"`

	// Attribute:
	//   api
	// Type:
	//   integer array
	// Description:
	//   List of supported API frameworks for this impression. Refer to
	//   List: API Frameworks in AdCOM. If an API is not explicitly listed,
	//   it is assumed not to be supported.
	// Note:
	//   OpenRTB <=2.5 defined only frameworks 1..6.
	API []adcom1.APIFramework `json:"api,omitempty"`

	// Attribute:
	//   sequence
	// Type:
	//   integer array
	// Description:
	//   Blocked creative attributes. Refer to List: Creative Attributes in
	//   AdCOM.
	// Note:
	//   OpenRTB <=2.5 defined only attributes with IDs 1..17.
	BAttr []adcom1.CreativeAttribute `json:"battr,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}

// Used to unmarshal not string Request
// Receive as interface and try to marshal into string then
func (native *Native) UnmarshalJSON(data []byte) error {
	type Alias Native
	aliasedNative := &struct {
		*Alias
		Request interface{} `json:"request"`
	}{
		Alias: (*Alias)(native),
	}

	err := json.Unmarshal(data, &aliasedNative)
	if err != nil {
		return err
	}

	receivedNativeRequest := aliasedNative.Request
	_, ok := receivedNativeRequest.(string)
	if ok {
		native.Request = receivedNativeRequest.(string)
		return nil
	}

	marshaledNativeRequest, err := JSONMarshal(receivedNativeRequest)
	if err != nil {
		return err
	}

	native.Request = string(marshaledNativeRequest)
	return nil
}

// If we need to send not string Request we fill RequestObj in advance and then marshal bid with aliased request field
func (native *Native) MarshalJSON() ([]byte, error) {
	if native.RequestObj == nil {
		return JSONMarshal(*native)
	}

	type Alias Native
	aliasedNative := &struct {
		*Alias
		Request interface{} `json:"request"`
	}{
		Alias:   (*Alias)(native),
		Request: native.RequestObj,
	}

	return JSONMarshal(aliasedNative)
}
