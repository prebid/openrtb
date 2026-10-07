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

	// VideoPlcmtPause represents an ad present in streaming video content that the consumer
	// has requested. While there may be other content surrounding the player, the video
	// content must be the focus of the user’s visit. It should remain the primary content
	// on the page and the only video player in-view.
	VideoPlcmtPause VideoPlcmtSubtype = 5

	// VideoPlcmtScreensaver represents an ad present when OS/App Screen Saver is initiated.
	VideoPlcmtScreensaver VideoPlcmtSubtype = 6

	// VideoPlcmtOverlay represents ads occurring during program content and outside of the
	// traditional ad break. The ad format places ad creatives over top of the program
	// content, often as a banner or picture in picture execution. While there may be other
	// content surrounding the player with the overlay, the player with the overlay must be
	// the focus of the user’s visit. It should remain the primary content on the page and
	// the primary video player in-view.
	VideoPlcmtOverlay VideoPlcmtSubtype = 7

	// VideoPlcmtSqueezeback represents ads Alongside or Adjacent to Content, also known as
	// L-Shape Ads & Double Box, which are ads that occur during program content and outside
	// of the traditional ad break. The program content is resized to allow for placement of
	// an ad on the screen. In a squeezeback, none of the content is covered by the ad,
	// instead it shares screen space with the ad. This is the major distinction between
	// squeezeback and overlay. While there may be other content surrounding the player with
	// the squeezeback, the player with the squeezeback must be the focus of the user visit.
	// It should remain the primary content on the page and the primary video player in-view
	VideoPlcmtSqueezeback VideoPlcmtSubtype = 8

	// VideoPlcmtInScene represents a form of advertising that integrates branded elements
	// directly within the content itself, rather than appearing as separate pre-roll,
	// mid-roll, overlay, or display formats. In Scene Advertising is composited into the
	// video content and blends the brand naturally into the scene, environment, or
	// storyline, ensuring contextual relevance and viewer immersion without interrupting
	// the user experience. In-scene ads typically consist of product placements or virtual
	// out of A insertions. At this time, the standardization focuses on the out of home
	// insertions in non-live content. While there may be other content surrounding the
	// player with the in-scene ad, the player with the in-scene ad must be the focus of the
	// user’s visit. It should remain the primary content on the page and the primary video
	// player in-view.
	VideoPlcmtInScene VideoPlcmtSubtype = 9
)
