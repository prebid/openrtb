package openrtb2

import (
	"github.com/prebid/openrtb/v20/common"
)

// Object: DurFloors
//
// This object allows sellers to specify price floors for video and audio creatives, whose price varies based on time.
// For example: 1-15 seconds at a floor of $5; 16-30 seconds at a floor of $10, > 31 seconds at a floor of $20.
// There are no explicit constraints on the defined ranges, nor guarantees that they don't overlap.
// In cases where multiple ranges may apply, it is up to the buyer and seller to coordinate on which floor is applicable.
type DurFloors = common.DurFloors
