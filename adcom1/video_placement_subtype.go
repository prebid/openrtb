package adcom1

// VideoPlacementSubtype represents types of video placements derived largely from the IAB Digital Video Guidelines.
// To be sent using the placement attribute in the Video object.
// DEPRECATED AS OF OPENRTB 2.6-202303.
// Proposed removal of this list and associated attribute in 2024.
type VideoPlacementSubtype int8

// Types of video placements derived largely from the IAB Digital Video Guidelines.
const (
	VideoPlacementInStream      VideoPlacementSubtype = 1 // In-Stream: Played before, during or after the streaming video content that the consumer has requested (e.g., Pre-roll, Mid-roll, Post-roll).
	VideoPlacementInBanner      VideoPlacementSubtype = 2 // In-Banner: Exists within a web banner that leverages the banner space to deliver a video experience as opposed to another static or rich media format. The format relies on the existence of display ad inventory on the page for its delivery.
	VideoPlacementInArticle     VideoPlacementSubtype = 3 // In-Article: Loads and plays dynamically between paragraphs of editorial content; existing as a standalone branded message.
	VideoPlacementInFeed        VideoPlacementSubtype = 4 // In-Feed: Found in content, social, or product feeds.
	VideoPlacementAlwaysVisible VideoPlacementSubtype = 5 // Interstitial/Slider/Floating: Covers the entire or a portion of screen area, but is always on screen while displayed (i.e. cannot be scrolled out of view).
)
