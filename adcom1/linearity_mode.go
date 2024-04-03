package adcom1

// LinearityMode represents options for media linearity, typically for video.
// Options for video linearity.
// “In-stream” or “linear” video refers to preroll, post-roll, or mid-roll video ads where the user is forced to watch ad in order to see the video content.
// “Overlay” or “non-linear” refer to ads that are shown on top of the video content.
//
// This OpenRTB list has values derived from the Inventory Quality Guidelines (IQG).
// Practitioners should keep in sync with updates to the IQG values.
type LinearityMode int8

// Options for media linearity, typically for video.
const (
	LinearityLinear    LinearityMode = 1 // Linear (i.e., In-Stream such as Pre-Roll, Mid-Roll, Post-Roll)
	LinearityNonLinear LinearityMode = 2 // Non-Linear (i.e., Overlay)
)
