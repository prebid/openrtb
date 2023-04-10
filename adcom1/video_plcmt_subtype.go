package adcom1

// VideoPlacementSubtype represents the the various types of video placements in accordance with updated IAB Digital Video Guidelines.
// To be sent using the plcmt attribute in Video object.
type VideoPlcmtSubtype int8

// Types of video placements derived largely from the IAB Digital Video Guidelines.
const (
	// VideoPlcmtInstream represents pre-roll, mid-roll, and post-roll ads that are played
	// before, during or after the streaming video content that the consumer has requested.
	// Instream video must be set to “sound on” by default at player start, or have
	// explicitly clear user intent to watch the video content. While there may be other
	// content surrounding the player, the video content must be the focus of the user’s
	// visit. It should remain the primary content on the page and the only video player
	// in-view capable of audio when playing. If the player converts to floating/sticky
	// subsequent ad calls should accurately convey the updated player size.
	VideoPlcmtInstream VideoPlcmtSubtype = 1

	// VideoPlcmtAccompanyingContent represents pre-roll, mid-roll, and post-roll ads that
	// are played before, during, or after streaming video content. The video player loads
	// and plays before, between, or after paragraphs of text or graphical content, and
	// starts playing only when it enters the viewport. Accompanying content should only
	// start playback upon entering the viewport. It may convert to a floating/sticky player
	// as it scrolls off the page.
	VideoPlcmtAccompanyingContent VideoPlcmtSubtype = 2

	// VideoPlcmtInterstitial represents video ads that are played without video content.
	// During playback, it must be the primary focus of the page and take up the majority
	// of the viewport and cannot be scrolled out of view. This can be in placements like
	// in-app video or slideshows.
	VideoPlcmtInterstitial VideoPlcmtSubtype = 3

	// VideoPlcmtNoContext represents no content / standalone video ads that are played
	// without streaming video content. This can be in placements like slideshows, native
	// feeds, in-content or sticky/floating.
	VideoPlcmtNoContent VideoPlcmtSubtype = 4
)
