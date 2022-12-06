package adcom1

// DOOHMultiplierMeasurementSourceType identifies the types of entities that provide quantity measurement for
// impression multipliers, which are common in DOOH (Digital Out of Home) advertising.
type DOOHMultiplierMeasurementSourceType int8

// MultiplierMeasurementSourceType options.
const (
	MultiplierUnknown                   DOOHMultiplierMeasurementSourceType = 0
	MultiplierMeasurementVendorProvided DOOHMultiplierMeasurementSourceType = 1
	MultiplierPublisherProvided         DOOHMultiplierMeasurementSourceType = 2
	MultiplierExchangeProvided          DOOHMultiplierMeasurementSourceType = 3
)
