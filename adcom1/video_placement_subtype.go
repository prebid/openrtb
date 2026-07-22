package adcom1

// VideoPlacementSubtype represents types of video placements derived largely from the IAB Digital Video Guidelines.
// To be sent using the placement attribute in the Video object.
// DEPRECATED AS OF OPENRTB 2.6-202303.
// Proposed removal of this list and associated attribute in 2024.
type VideoPlacementSubtype int8

// Types of video placements derived largely from the IAB Digital Video Guidelines.
const (
	VideoPlacementInStream            VideoPlacementSubtype = 1 // In-Stream: Played before, during or after the streaming video content that the consumer has requested (e.g., Pre-roll, Mid-roll, Post-roll).
	VideoPlacementAccompanyingContent VideoPlacementSubtype = 2 // Accompanying Content: Pre-roll, mid-roll, and post-roll ads that are played before, during, or after streaming video content.
	VideoPlacementInterstitial        VideoPlacementSubtype = 3 // Interstitial: Video ads that are played without video content.
	VideoPlacementNoContent           VideoPlacementSubtype = 4 // No Content/Standalone: Video ads that are played without streaming video content.
	VideoPlacementPause               VideoPlacementSubtype = 5 // Pause: An ad present in streaming video content that the consumer has requested.
	VideoPlacementScreensaver         VideoPlacementSubtype = 6 // Screensaver: An ad present, when OS/App Screen Saver are initiated.
	VideoPlacementOverlay             VideoPlacementSubtype = 7 // Overlay: Ads occurring during program content and outside of the traditional ad break.
	VideoPlacementSqueezeback         VideoPlacementSubtype = 8 // Squeezeback: Ads Alongside or Adjacent to Content, also known as L-Shape Ads & Double Box, are ads that occur during program content and outside of the traditional ad break.
	VideoPlacementInScene             VideoPlacementSubtype = 9 // In-scene: A form of advertising that integrates branded elements directly within the content itself, rather than appearing as separate pre-roll, mid-roll, overlay, or display formats.
)
