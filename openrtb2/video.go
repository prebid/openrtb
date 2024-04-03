package openrtb2

import (
	"github.com/onfocusio/openrtb-adagio/common"
)

// 3.2.7 Object: Video
//
// This object represents an in-stream video impression.
// Many of the fields are non-essential for minimally viable transactions, but are included to offer fine control when needed.
// Video in OpenRTB generally assumes compliance with the VAST standard.
// As such, the notion of companion ads is supported by optionally including an array of Banner objects (refer to the Banner object in Section 3.2.6) that define these companion ads.
//
// The presence of a Video as a subordinate of the Imp object indicates that this impression is offered as a video type impression.
// At the publisher’s discretion, that same impression may also be offered as banner, audio, and/or native by also including as Imp subordinates objects of those types.
// However, any given bid for the impression must conform to one of the offered types.

type Video = common.Video
