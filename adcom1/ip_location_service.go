package adcom1

// IPLocationService represents services and/or vendors used for resolving IP addresses to geolocations.
type IPLocationService int16

// Services and/or vendors used for resolving IP addresses to geolocations.
const (
	LocationServiceIP2Location               IPLocationService = 1   // ip2location
	LocationServiceNeustar                   IPLocationService = 2   // Neustar (Quova)
	LocationServiceMaxMind                   IPLocationService = 3   // MaxMind
	LocationServiceNetAcuity                 IPLocationService = 4   // NetAcuity (Digital Element)
	LocationService51DegreesHighConfidence   IPLocationService = 511 // 51Degrees (High Confidence)
	LocationService51DegreesMediumConfidence IPLocationService = 512 // 51Degrees (Medium Confidence)
)
