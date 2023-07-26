package openrtb2

import (
	"encoding/json"

	"github.com/prebid/openrtb/v20/adcom1"
)

// 3.2.1 Object: BidRequest
//
// The top-level bid request object contains a globally unique bid request or auction ID.
// This id attribute is required as is at least one impression object (Section 3.2.4).
// Other attributes in this top-level object establish rules and restrictions that apply to all impressions being offered.
//
// There are also several subordinate objects that provide detailed data to potential buyers.
// Among these are the Site and App objects, which describe the type of published media in which the impression(s) appear.
// These objects are highly recommended, but only one applies to a given bid request depending on whether the media is browser-based web content or a non-browser application, respectively.
type BidRequest struct {

	// Attribute:
	//   id
	// Type:
	//   string; required
	// Description:
	//   Unique ID of the bid request, provided by the exchange.
	ID string `json:"id"`

	// Attribute:
	//   imp
	// Type:
	//   object array; required
	// Description:
	//   Array of Imp objects (Section 3.2.4) representing the
	//   impressions offered. At least 1 Imp object is required.
	Imp []Imp `json:"imp"`

	// Attribute:
	//   site
	// Type:
	//   object; recommended
	// Description:
	//    Details via a Site object (Section 3.2.13) about the publisher’s
	//    website. Only applicable and recommended for websites.
	Site *Site `json:"site,omitempty"`

	// Attribute:
	//   app
	// Type:
	//   object; recommended
	// Description:
	//    Details via an App object (Section 3.2.14) about the publisher’s
	//    app (i.e., non-browser applications). Only applicable and
	//    recommended for apps.
	App *App `json:"app,omitempty"`

	// Attribute:
	//   dooh
	// Type:
	//   object
	// Description:
	//   This object should be included if the ad supported content is a
	//   Digital Out-Of-Home screen. A bid request with a DOOH object must
	//   not contain a site or app object.
	DOOH *DOOH `json:"dooh,omitempty"`

	// Attribute:
	//   device
	// Type:
	//   object; recommended
	// Description:
	//   Details via a Device object (Section 3.2.18) about the user’s
	//   device to which the impression will be delivered.
	Device *Device `json:"device,omitempty"`

	// Attribute:
	//   user
	// Type:
	//   object; recommended
	// Description:
	//    Details via a User object (Section 3.2.20) about the human
	//    user of the device; the advertising audience.
	User *User `json:"user,omitempty"`

	// Attribute:
	//   test
	// Type:
	//   integer; default 0
	// Description:
	//    Indicator of test mode in which auctions are not billable,
	//    where 0 = live mode, 1 = test mode.
	Test int8 `json:"test,omitempty"`

	// Attribute:
	//   at
	// Type:
	//   integer; default 2
	// Description:
	//    Auction type, where 1 = First Price, 2 = Second Price Plus.
	//    Exchange-specific auction types can be defined using values
	//    500 and greater.
	AT int64 `json:"at,omitempty"`

	// Attribute:
	//   tmax
	// Type:
	//   integer
	// Description:
	//    Maximum time in milliseconds the exchange allows for bids to
	//    be received including Internet latency to avoid timeout. This
	//    value supersedes any a priori guidance from the exchange.
	TMax int64 `json:"tmax,omitempty"`

	// Attribute:
	//   wseat
	// Type:
	//   string array
	// Description:
	//   Allowed list of buyer seats (e.g., advertisers, agencies)
	//   allowed to bid on this impression. IDs of seats and knowledge
	//   of the buyer’s customers to which they refer must be
	//   coordinated between bidders and the exchange a priori. At
	//   most, only one of wseat and bseat should be used in the
	//   same request. Omission of both implies no seat restrictions.
	WSeat []string `json:"wseat,omitempty"`

	// Attribute:
	//   bseat
	// Type:
	//   string array
	// Description:
	//   Block list of buyer seats (e.g., advertisers, agencies) restricted
	//   from bidding on this impression. IDs of seats and knowledge
	//   of the buyer’s customers to which they refer must be
	//   coordinated between bidders and the exchange a priori. At
	//   most, only one of wseat and bseat should be used in the
	//   same request. Omission of both implies no seat restrictions.
	BSeat []string `json:"bseat,omitempty"`

	// Attribute:
	//   allimps
	// Type:
	//   integer; default 0
	// Description:
	//   Flag to indicate if Exchange can verify that the impressions
	//   offered represent all of the impressions available in context
	//   (e.g., all on the web page, all video spots such as pre/mid/post
	//   roll) to support road-blocking. 0 = no or unknown, 1 = yes, the
	//   impressions offered represent all that are available.
	AllImps int8 `json:"allimps,omitempty"`

	// Attribute:
	//   cur
	// Type:
	//   string array
	// Description:
	//   Array of allowed currencies for bids on this bid request using
	//   ISO-4217 alpha codes. Recommended only if the exchange
	//   accepts multiple currencies.
	Cur []string `json:"cur,omitempty"`

	// Attribute:
	//   wlang
	// Type:
	//   string array
	// Description:
	//   Allowed list of languages for creatives using ISO-639-1-alpha-2.
	//   Omission implies no specific restrictions, but buyers would be
	//   advised to consider language attribute in the Device and/or
	//   Content objects if available. Only one of wlang or wlangb
	//   should be present.
	WLang []string `json:"wlang,omitempty"`

	// Attribute:
	//   wlangb
	// Type:
	//   string array
	// Description:
	//   Allowed list of languages for creatives using IETF BCP 47I.
	//   Omission implies no specific restrictions, but buyers would be
	//   advised to consider language attribute in the Device and/or
	//   Content objects if available. Only one of wlang or wlangb
	//   should be present.
	WLangB []string `json:"wlangb,omitempty"`

	// Attribute:
	//   bcat
	// Type:
	//   string array
	// Description:
	//   Blocked advertiser categories using the specified
	//   category taxonomy.
	//   The taxonomy to be used is defined by the cattax field. If no
	//   cattax field is supplied IAB Content Category Taxonomy 1.0 is
	//   assumed.
	BCat []string `json:"bcat,omitempty"`

	// Attribute:
	//   cattax
	// Type:
	//   integer; default 1
	// Description:
	//   The taxonomy in use for bcat. Refer to the AdCOM
	//   1.0 list List: Category Taxonomies for values
	CatTax adcom1.CategoryTaxonomy `json:"cattax,omitempty"`

	// Attribute:
	//   badv
	// Type:
	//   string array
	// Description:
	//   Block list of advertisers by their domains (e.g., “ford.com”).
	BAdv []string `json:"badv,omitempty"`

	// Attribute:
	//   bapp
	// Type:
	//   string array
	// Description:
	//   Block list of applications by their app store IDs. See OTT/CTV
	//   Store Assigned App Identification Guidelines for more details
	//   about expected strings for CTV app stores. For mobile apps in
	//   Google Play Store, these should be bundle or package names
	//   (e.g. com.foo.mygame). For apps in Apple App Store, these
	//   should be a numeric ID.
	BApp []string `json:"bapp,omitempty"`

	// Attribute:
	//   source
	// Type:
	//   object
	// Description:
	//   A Sorce object (Section 3.2.2) that provides data about the
	//   inventory source and which entity makes the final decision.
	Source *Source `json:"source,omitempty"`

	// Attribute:
	//   regs
	// Type:
	//   object
	// Description:
	//   A Regs object (Section 3.2.3) that specifies any industry, legal,
	//   or governmental regulations in force for this request.
	Regs *Regs `json:"regs,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
