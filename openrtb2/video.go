package openrtb2

import (
	"encoding/json"

	"github.com/prebid/openrtb/v20/adcom1"
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
	//   poddedupe
	// Type:
	//   enum array; provisional
	// Description:
	//   Indicates pod deduplication settings that will be applied to bid
	//   responses. Refer to List: Pod Deduplication in AdCOM 1.0.
	PodDedupe []adcom1.PodDedupe `json:"poddedupe,omitempty"`

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
