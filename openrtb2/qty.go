package openrtb2

import "github.com/prebid/openrtb/v17/adcom1"

// Object: Qty
//
// A programmatic impression is often referred to as a ‘spot’ in digital out-of-home and CTV, with an impression being a unique member of the audience viewing it.
// Therefore, a standard means of passing a multiplier in the bid request, representing the total quantity of impressions, is required.
// This object includes the impression multiplier, and describes the source of the multiplier value.
type Qty struct {

	// Attribute:
	//   multiplier
	// Type:
	//   float; required
	// Description:
	//   The quantity of billable events which will be deemed to have occurred
	//   if this item is purchased. For example, a DOOH opportunity may be
	//   considered to be 14.2 impressions. Equivalent to qtyflt in OpenRTB 3.0.
	Multiplier float64 `json:"multiplier,omitempty"`

	// Attribute:
	//   sourcetype
	// Type:
	//   integer; recommended
	// Description:
	//   The source type of the quantity measurement, ie. publisher. Refer to
	//   List: DOOH Multiplier Measurement Source Types.
	SourceType adcom1.MultiplierMeasurementSourceType `json:"sourcetype,omitempty"`

	// Attribute:
	//   vendor
	// Type:
	//   string; required if sourcetype is present and type = 1
	// Description:
	//   The top level business domain name of the measurement vendor providing
	//   the quantity measurement.
	Vendor string `json:"vendor,omitempty"`
}
