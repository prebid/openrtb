package request

import (
	"encoding/json"
	"github.com/prebid/openrtb/v20/adcom1"
)

// 4.5 Video Object
//
// The video object to be used for all video elements supported in the Native Ad.
// This corresponds to the Video object of OpenRTB.
// Exchange implementers can impose their own specific restrictions.
// Here are the required attributes of the Video Object.
// For optional attributes please refer to OpenRTB.
type Video struct {

	// Attribute:
	//   mimes
	// Type:
	//   string array; required
	// Description:
	//   Content MIME types supported (e.g., “video/x-ms-wmv”,
	//   “video/mp4”).
	MIMEs []string `json:"mimes"`

	// Attribute:
	//   minduration
	// Type:
	//   integer; default 0; recommended
	// Definition:
	//   Minimum video ad duration in seconds. This field is mutually
	//   exclusive with rqddurs; only one of minduration and rqddurs
	//   may be in a bid request.
	MinDuration int64 `json:"minduration,omitempty"`

	// Attribute:
	//   maxduration
	// Type:
	//   integer; recommended
	// Definition:
	//  Maximum video ad duration in seconds. This field is mutually
	//  exclusive with rqddurs; only one of maxduration and rqddurs
	//  may be in a bid request.
	MaxDuration int64 `json:"maxduration,omitempty"`

	// Attribute:
	//   startdelay
	// Type:
	//   integer; recommended
	// Definition:
	//   Indicates the start delay in seconds for pre-roll, mid-roll, or
	//   post-roll ad placements. Refer to List: Start Delay Modes
	//   in AdCOM 1.0.
	StartDelay *adcom1.StartDelay `json:"startdelay,omitempty"`

	// Attribute:
	//   maxseq
	// Type:
	//   integer; recommended
	// Definition:
	//   Indicates the maximum number of ads that may be served into
	//   a “dynamic” video ad pod (where the precise number of ads is
	//   not predetermined by the seller). See Section 7.6 for more
	//   details.
	MaxSeq int64 `json:"maxseq,omitempty"`

	// Attribute:
	//   poddur
	// Type:
	//   integer; recommended
	// Definition:
	//   Indicates the total amount of time in seconds that advertisers
	//   may fill for a “dynamic” video ad pod (See Section 7.6 for more
	//   details), or the dynamic portion of a “hybrid” ad pod. This field
	//   is required only for the dynamic portion(s) of video ad pods.
	//   This field refers to the length of the entire ad break, whereas
	//   minduration/maxduration/rqddurs are constraints relating to
	//   the slots that make up the pod.
	PodDur int64 `json:"poddur,omitempty"`

	// Attribute:
	//   protocols
	// Type:
	//   integer array; recommended
	// Definition:
	//   Array of supported video protocols. Refer to List: Creative
	//   Subtypes - Audio/Video in AdCOM 1.0.
	// Note:
	//   OpenRTB <=2.5 defined only protocols 1..10.
	Protocols []adcom1.MediaCreativeSubtype `json:"protocols,omitempty"`

	// Attribute:
	//   protocol
	// Type:
	//   integer; DEPRECATED; REMOVED in OpenRTB 2.6
	// Description:
	//   NOTE: Deprecated in favor of protocols.
	//   Supported video protocol. Refer to List 5.8. At least one
	//   supported protocol must be specified in either the protocol
	//   or protocols attribute.
	// Note:
	//   OpenRTB <=2.5 defined only protocols 1..10.
	Protocol adcom1.MediaCreativeSubtype `json:"protocol,omitempty"`

	// Attribute:
	//   w
	// Type:
	//   integer; recommended
	// Description:
	//   Width of the video player in device independent pixels (DIPS).
	W *int64 `json:"w,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer; recommended
	// Description:
	//   Height of the video player in device independent pixels (DIPS).
	H *int64 `json:"h,omitempty"`

	// Attribute:
	//   podid
	// Type:
	//   integer
	// Definition:
	//   Unique identifier indicating that an impression opportunity
	//   belongs to a video ad pod. If multiple impression opportunities
	//   within a bid request share the same podid, this indicates that
	//   those impression opportunities belong to the same video ad
	//   pod.
	PodID string `json:"podid,omitempty"`

	// Attribute:
	//   podseq
	// Type:
	//   integer; default 0
	// Definition:
	//   The sequence (position) of the video ad pod within a
	//  content stream. Refer to List: Pod Sequence in AdCOM 1.0
	//  for guidance on the use of this field.
	PodSeq adcom1.PodSequence `json:"podseq,omitempty"`

	// Attribute:
	//   rqddurs
	// Type:
	//   integer array
	// Definition:
	//   Precise acceptable durations for video creatives in
	//   seconds. This field specifically targets the Live TV use case
	//   where non-exact ad durations would result in undesirable
	//   ‘dead air’. This field is mutually exclusive with minduration
	//   and maxduration; if rqddurs is specified, minduration and
	//   maxduration must not be specified and vice versa.
	RqdDurs []int64 `json:"rqddurs,omitempty"`

	// Attribute:
	//   placement
	// Type:
	//   integer; DEPRECATED
	// Description:
	//   Video placement type for the impression. Refer to List:
	//   Placement Subtypes - Video in AdCOM 1.0.
	Placement adcom1.VideoPlacementSubtype `json:"placement,omitempty"`

	// Attribute:
	//   plcmt
	// Type:
	//   integer
	// Description:
	//   Video placement type for the impression. Refer to List: Plcmt Subtypes - Video in AdCOM 1.0.
	Plcmt adcom1.VideoPlcmtSubtype `json:"plcmt,omitempty"`

	// Attribute:
	//   linearity
	// Type:
	//   integer
	// Description:
	//   Indicates if the impression must be linear, nonlinear, etc. If
	//   none specified, assume all are allowed. Refer to List: Linearity
	//   Modes in AdCOM 1.0. Note that this field describes the
	//   expected VAST response and not whether a placement is in-
	//   stream, out-stream, etc. For that, see placement.
	Linearity adcom1.LinearityMode `json:"linearity,omitempty"`

	// Attribute:
	//   skip
	// Type:
	//   integer
	// Description:
	//   Indicates if the player will allow the video to be skipped,
	//   where 0 = no, 1 = yes.
	//   If a bidder sends markup/creative that is itself skippable, the
	//   Bid object should include the attr array with an element of
	//   16 indicating skippable video. Refer to List: Creative
	//   Attributes in AdCOM 1.0.
	Skip *int8 `json:"skip,omitempty"`

	// Attribute:
	//   skipmin
	// Type:
	//   integer; default 0
	// Description:
	//   Videos of total duration greater than this number of seconds
	//   can be skippable; only applicable if the ad is skippable.
	SkipMin int64 `json:"skipmin,omitempty"`

	// Attribute:
	//   skipafter
	// Type:
	//   integer; default 0
	// Description:
	//   Number of seconds a video must play before skipping is
	//   enabled; only applicable if the ad is skippable
	SkipAfter int64 `json:"skipafter,omitempty"`

	// Attribute:
	//   sequence
	// Type:
	//   integer; default 0; DEPRECATED
	// Description:
	//   If multiple ad impressions are offered in the same bid request,
	//   the sequence number will allow for the coordinated delivery
	//   of multiple creatives.
	Sequence int8 `json:"sequence,omitempty"`

	// Attribute:
	//   slotinpod
	// Type:
	//   integer; default 0
	// Description:
	//   For video ad pods, this value indicates that the seller can
	//   guarantee delivery against the indicated slot position in the
	//   pod. Refer to  List: Slot Position in Pod in AdCOM 1.0 guidance
	//   on the use of this field.
	SlotInPod adcom1.SlotPositionInPod `json:"slotinpod,omitempty"`

	// Attribute:
	//   mincpmpersec
	// Type:
	//   float
	// Description:
	//   Minimum CPM per second. This is a price floor for the
	//   "dynamic" portion of a video ad pod, relative to the duration
	//   of bids an advertiser may submit.
	MinCPMPerSec float64 `json:"mincpmpersec,omitempty"`

	// Attribute:
	//   battr
	// Type:
	//   integer array
	// Description:
	//   Blocked creative attributes. Refer to List: Creative Attributes in
	//   AdCOM 1.0
	// Note:
	//   OpenRTB <=2.5 defined only attributes 1..17.
	BAttr []adcom1.CreativeAttribute `json:"battr,omitempty"`

	// Attribute:
	//   maxextended
	// Type:
	//   integer
	// Description:
	//   Maximum extended ad duration if extension is allowed. If
	//   blank or 0, extension is not allowed. If -1, extension is
	//   allowed, and there is no time limit imposed. If greater than 0,
	//   then the value represents the number of seconds of extended
	//   play supported beyond the maxduration value.
	MaxExtended int64 `json:"maxextended,omitempty"`

	// Attribute:
	//   minbitrate
	// Type:
	//   integer
	// Description:
	//   Minimum bit rate in Kbps (kilobits per second).
	MinBitRate int64 `json:"minbitrate,omitempty"`

	// Attribute:
	//   maxbitrate
	// Type:
	//   integer
	// Description:
	//   Maximum bit rate in Kbps (kilobits per second).
	MaxBitRate int64 `json:"maxbitrate,omitempty"`

	// Attribute:
	//   boxingallowed
	// Type:
	//   integer; default 1
	// Description:
	//   Indicates if letter-boxing of 4:3 content into a 16:9 window is
	//   allowed, where 0 = no, 1 = yes.
	BoxingAllowed *int8 `json:"boxingallowed,omitempty"`

	// Attribute:
	//   playbackmethod
	// Type:
	//   integer array
	// Description:
	//   Playback methods that may be in use. If none are specified,
	//   any method may be used. Refer to List: Playback Methods
	//   in AdCOM 1.0. Only one method is typically used in practice.
	//   As a result, this array may be converted to an integer in a
	//   future version of the specification. It is strongly advised to use
	//   only the first element of this array in preparation for this
	//   change.
	// Note:
	//   OpenRTB <=2.5 defined only methods 1..6.
	PlaybackMethod []adcom1.PlaybackMethod `json:"playbackmethod,omitempty"`

	// Attribute:
	//   playbackend
	// Type:
	//   integer
	// Description:
	//   The event that causes playback to end. Refer to List: Playback
	//   Cessation Modes in AdCOM 1.0.
	PlaybackEnd adcom1.PlaybackCessationMode `json:"playbackend,omitempty"`

	// Attribute:
	//   delivery
	// Type:
	//   integer array
	// Description:
	//   Supported delivery methods (e.g., streaming, progressive). If
	//   none specified, assume all are supported. Refer to List:
	//   Delivery Methods in AdCOM 1.0.
	Delivery []adcom1.DeliveryMethod `json:"delivery,omitempty"`

	// Attribute:
	//   pos
	// Type:
	//   integer
	// Description:
	//   Ad position on screen. Refer to List: Placement Positions in
	//   AdCOM 1.0.
	Pos *adcom1.PlacementPosition `json:"pos,omitempty"`

	// Attribute:
	//   companionad
	// Type:
	//   object array
	// Description:
	//   Array of Banner objects (Section 3.2.6) if companion ads are
	//   available.
	CompanionAd []Banner `json:"companionad,omitempty"`

	// Attribute:
	//   api
	// Type:
	//   integer array
	// Description:
	//   List of supported API frameworks for this impression. Refer to
	//   List: API Frameworks in AdCOM 1.0. If an API is not explicitly
	//   listed, it is assumed not to be supported.
	// Note:
	//   OpenRTB <=2.5 defined only frameworks 1..6.
	API []adcom1.APIFramework `json:"api,omitempty"`

	// Attribute:
	//   companiontype
	// Type:
	//   integer array
	// Description:
	//   Supported VAST companion ad types. Refer to List:
	//   Companion Types in AdCOM 1.0. Recommended if
	//   companion Banner objects are included via the
	//   companionad array. If one of these banners will be
	//   rendered as an end-card, this can be specified using the vcm
	//   attribute with the particular banner (Section 3.2.6).
	CompanionType []adcom1.CompanionType `json:"companiontype,omitempty"`

	// Attribute:
	//   durfloors
	// Type:
	//   object array
	// Description:
	//   An array of DurFloors objects (Section 3.2.35) indicating the floor
	//   prices for video creatives of various durations that the buyer may bid with.
	DurFloors []DurFloors `json:"durfloors,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}

// 3.2.6 Object: Banner
//
// This object represents the most general type of impression.
// Although the term “banner” may have very specific meaning in other contexts, here it can be many things including a simple static image, an expandable ad unit, or even in-banner video (refer to the Video object in Section 3.2.7 for the more generalized and full featured video ad units).
// An array of Banner objects can also appear within the Video to describe optional companion ads defined in the VAST specification.
//
// The presence of a Banner as a subordinate of the Imp object indicates that this impression is offered as a banner type impression.
// At the publisher’s discretion, that same impression may also be offered as video, audio, and/or native by also including as Imp subordinates objects of those types.
// However, any given bid for the impression must conform to one of the offered types.
type Banner struct {

	// Attribute:
	//   format
	// Type:
	//   object array; recommended
	// Description:
	//   Array of format objects (Section 3.2.10) representing the
	//   banner sizes permitted. If none are specified, then use of the
	//   h and w attributes is highly recommended.
	Format []Format `json:"format,omitempty"`

	// Attribute:
	//   w
	// Type:
	//   integer; recommended
	// Description:
	//   Exact width in device independent pixels (DIPS);
	//   recommended if no format objects are specified.
	W *int64 `json:"w,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer; recommended
	// Description:
	//   Exact height in device independent pixels (DIPS);
	//   recommended if no format objects are specified.
	H *int64 `json:"h,omitempty"`

	// Attribute:
	//   wmax
	// Type:
	//   integer; DEPRECATED; REMOVED in OpenRTB 2.6
	// Description:
	//   NOTE: Deprecated in favor of the format array.
	//   Maximum width in device independent pixels (DIPS).
	WMax int64 `json:"wmax,omitempty"`

	// Attribute:
	//   hmax
	// Type:
	//   integer; DEPRECATED; REMOVED in OpenRTB 2.6
	// Description:
	//   NOTE: Deprecated in favor of the format array.
	//   Maximum height in device independent pixels (DIPS).
	HMax int64 `json:"hmax,omitempty"`

	// Attribute:
	//   wmin
	// Type:
	//   integer; DEPRECATED; REMOVED in OpenRTB 2.6
	// Description:
	//   NOTE: Deprecated in favor of the format array.
	//   Minimum width in device independent pixels (DIPS).
	WMin int64 `json:"wmin,omitempty"`

	// Attribute:
	//   hmin
	// Type:
	//   integer; DEPRECATED; REMOVED in OpenRTB 2.6
	// Description:
	//   NOTE: Deprecated in favor of the format array.
	//   Minimum height in device independent pixels (DIPS).
	HMin int64 `json:"hmin,omitempty"`

	// Attribute:
	//   btype
	// Type:
	//   integer array
	// Description:
	//   Blocked banner ad types.
	//   Refer to BannerAdType constants.
	BType []BannerAdType `json:"btype,omitempty"`

	// Attribute:
	//   battr
	// Type:
	//   integer array
	// Description:
	//   Blocked creative attributes. Refer to List: Creative Attributes in AdCOM 1.0.
	// Note:
	//   OpenRTB <=2.5 defined only attributes with IDs 1..17.
	BAttr []adcom1.CreativeAttribute `json:"battr,omitempty"`

	// Attribute:
	//   pos
	// Type:
	//   integer
	// Description:
	//   Ad position on screen. Refer to List: Placement Positions in AdCOM 1.0.
	Pos *adcom1.PlacementPosition `json:"pos,omitempty"`

	// Attribute:
	//   mimes
	// Type:
	//   string array
	// Description:
	//   Content MIME types supported. Popular MIME types may include,
	//   "image/jpeg" and "image/gif".
	MIMEs []string `json:"mimes,omitempty"`

	// Attribute:
	//   topframe
	// Type:
	//   integer
	// Description:
	//   Indicates if the banner is in the top frame as opposed to an
	//   iframe, where 0 = no, 1 = yes.
	TopFrame int8 `json:"topframe,omitempty"`

	// Attribute:
	//   expdir
	// Type:
	//   integer array
	// Description:
	//   Directions in which the banner may expand. Refer to List: Expandable
	//   Directions in AdCOM 1.0.
	// Note:
	//   OpenRTB <=2.5 defined only directions 1..5.
	ExpDir []adcom1.ExpandableDirection `json:"expdir,omitempty"`

	// Attribute:
	//   api
	// Type:
	//   integer array
	// Description:
	//   List of supported API frameworks for this impression. Refer to List: API
	//   Frameworks in AdCOM 1.0. If an API is not explicitly listed, it is assumed
	//   not to be supported.
	// Note:
	//   OpenRTB <=2.5 defined only frameworks 1..6.
	API []adcom1.APIFramework `json:"api,omitempty"`

	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   Unique identifier for this banner object. Recommended when
	//   Banner objects are used with a Video object (Section 3.2.7) to
	//   represent an array of companion ads. Values usually start at 1
	//   and increase with each object; should be unique within an
	//   impression.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   vcm
	// Type:
	//   integer
	// Description:
	//   Relevant only for Banner objects used with a Video object
	//   (Section 3.2.7) in an array of companion ads. Indicates the
	//   companion banner rendering mode relative to the associated
	//   video, where 0 = concurrent, 1 = end-card.
	Vcm *int8 `json:"vcm,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}

// 3.2.10 Object: Format
//
// This object represents an allowed size (i.e., height and width combination) or Flex Ad parameters for a banner impression.
// These are typically used in an array where multiple sizes are permitted.
// It is recommended that either the w/h pair or the wratio/hratio/wmin set (i.e., for Flex Ads) be specified.
type Format struct {

	// Attribute:
	//   w
	// Type:
	//   integer
	// Description:
	//   Width in device independent pixels (DIPS).
	W int64 `json:"w,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer
	// Description:
	//   Height in device independent pixels (DIPS).
	H int64 `json:"h,omitempty"`

	// Attribute:
	//   wratio
	// Type:
	//   integer
	// Description:
	//   Relative width when expressing size as a ratio
	WRatio int64 `json:"wratio,omitempty"`

	// Attribute:
	//   hratio
	// Type:
	//   Integer
	// Description:
	//   Relative height when expressing size as a ratio.
	HRatio int64 `json:"hratio,omitempty"`

	// Attribute:
	//   wmin
	// Type:
	//   integer
	// Description:
	//   The minimum width in device independent pixels (DIPS) at
	//   which the ad will be displayed the size is expressed as a ratio.
	WMin int64 `json:"wmin,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}

// BannerAdType
//
// Types of ads that can be accepted by the exchange unless restricted by publisher site settings.
type BannerAdType int8

const (
	BannerAdTypeXHTMLTextAd   BannerAdType = 1 // XHTML Text Ad (usually mobile)
	BannerAdTypeXHTMLBannerAd BannerAdType = 2 // XHTML Banner Ad. (usually mobile)
	BannerAdTypeJavaScriptAd  BannerAdType = 3 // JavaScript Ad; must be valid XHTML (i.e., Script Tags Included)
	BannerAdTypeIframe        BannerAdType = 4 // iframe
)

// Object: DurFloors
//
// This object allows sellers to specify price floors for video and audio creatives, whose price varies based on time.
// For example: 1-15 seconds at a floor of $5; 16-30 seconds at a floor of $10, > 31 seconds at a floor of $20.
// There are no explicit constraints on the defined ranges, nor guarantees that they don't overlap.
// In cases where multiple ranges may apply, it is up to the buyer and seller to coordinate on which floor is applicable.
type DurFloors struct {

	// Attribute:
	//   mindur
	// Type:
	//   integer
	// Description:
	//   An integer indicating the low end of a duration range. If this
	//   value is missing, the low end is unbounded. Either mindur or maxdur
	//   is required, but not both.
	MinDur int64 `json:"mindur,omitempty"`

	// Attribute:
	//   maxdur
	// Type:
	//   integer
	// Description:
	//   An integer indicating the high end of a duration range. If this
	//   value is missing, the high end is unbounded. Either mindur or maxdur
	//   is required, but not both.
	MaxDur int64 `json:"maxdur,omitempty"`

	// Attribute:
	//   bidfloor
	// Type:
	//   float; default 0
	// Description:
	//   Minimum bid for a given impression opportunity, if bidding with a
	//   creative in this duration range, expressed in CPM. For any creatives
	//   whose durations are outside of the defined min/max, the bidfloor at
	//   the Imp level will serve as the default floor.
	BidFloor float64 `json:"bidfloor,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Placeholder for vendor specific extensions to this object.
	Ext json.RawMessage `json:"ext,omitempty"`
}
