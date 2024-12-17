package openrtb2

import (
	"github.com/prebid/openrtb/v20/common"
)

// 3.2.6 Object: Banner
//
// This object represents the most general type of impression.
// Although the term “banner” may have very specific meaning in other contexts, here it can be many things including a simple static image, an expandable ad unit, or even in-banner video (refer to the Video object in Section 3.2.7 for the more generalized and full featured video ad units).
// An array of Banner objects can also appear within the Video to describe optional companion ads defined in the VAST specification.
//
// The presence of a Banner as a subordinate of the Imp object indicates that this impression is offered as a banner type impression.
// At the publisher’s discretion, that same impression may also be offered as video, audio, and/or native by also including as Imp subordinates objects of those types.
// However, any given bid for the impression must conform to one of the offered types.
type Banner = common.Banner
