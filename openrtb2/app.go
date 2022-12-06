package openrtb2

import (
	"encoding/json"

	"github.com/prebid/openrtb/v17/adcom1"
)

// Object: App
//
// This object should be included if the ad supported content is a non-browser application (typically in mobile) as opposed to a website.
// A bid request with an app object must not contain a site or DOOH object.
// At a minimum, it is useful to provide an App ID or bundle, but this is not strictly required.
type App struct {

	// Attribute:
	//   id
	// Type:
	//   string; recommended
	// Description:
	//   Exchange-specific app ID.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   App name (may be aliased at the publisher’s request).
	Name string `json:"name,omitempty"`

	// Attribute:
	//   bundle
	// Type:
	//   string
	// Description:
	//   The store ID of the app in an app store. See OTT/CTV Store
	//   Assigned App Identification Guidelines for more details about
	//   expected strings for CTV app stores. For mobile apps in
	//   Google Play Store, these should be bundle or package names
	//   (e.g. com.foo.mygame). For apps in Apple App Store, these
	//   should be a numeric ID.
	Bundle string `json:"bundle,omitempty"`

	// Attribute:
	//   domain
	// Type:
	//   string
	// Description:
	//   Domain of the app (e.g., “mygame.foo.com”).
	Domain string `json:"domain,omitempty"`

	// Attribute:
	//   storeurl
	// Type:
	//   string
	// Description:
	//   App store URL for an installed app; for IQG 2.1 compliance.
	StoreURL string `json:"storeurl,omitempty"`

	// Attribute:
	//   cattax
	// Type:
	//   integer; default 1
	// Description:
	//   The taxonomy in use. Refer to the AdCOM list List: Category
	//   Taxonomies for values.
	CatTax adcom1.CategoryTaxonomy `json:"cattax,omitempty"`

	// Attribute:
	//   cat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories of the app. The taxonomy to be
	//   used is defined by the cattax field. If no cattax field is supplied
	//   IAB Content Category Taxonomy 1.0 is assumed.
	Cat []string `json:"cat,omitempty"`

	// Attribute:
	//   sectioncat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories that describe the current
	//   section of the app.
	//   The taxonomy to be used is defined by the cattax field.
	SectionCat []string `json:"sectioncat,omitempty"`

	// Attribute:
	//   pagecat
	// Type:
	//   string array
	// Description:
	//   Array of IAB content categories that describe the current page
	//   or view of the app.
	//   The taxonomy to be used is defined by the cattax field.
	PageCat []string `json:"pagecat,omitempty"`

	// Attribute:
	//   ver
	// Type:
	//   string
	// Description:
	//   Application version.
	Ver string `json:"ver,omitempty"`

	// Attribute:
	//   privacypolicy
	// Type:
	//   integer
	// Description:
	//   Indicates if the app has a privacy policy, where 0 = no, 1 = yes.
	PrivacyPolicy int8 `json:"privacypolicy,omitempty"`

	// Attribute:
	//   paid
	// Type:
	//   integer
	// Description:
	//   0 = app is free, 1 = the app is a paid version.
	Paid int8 `json:"paid,omitempty"`

	// Attribute:
	//   publisher
	// Type:
	//   object
	// Description:
	//   Details about the Publisher (Section 3.2.15) of the app.
	Publisher *Publisher `json:"publisher,omitempty"`

	// Attribute:
	//   content
	// Type:
	//   object
	// Description:
	//   Details about the Content (Section 3.2.16) within the app
	Content *Content `json:"content,omitempty"`

	// Attribute:
	//   keywords
	// Type:
	//   string
	// Description:
	//   Comma separated list of keywords about the app. Only one of
	//   ‘keywords’ or ‘kwarray’ may be present.
	Keywords string `json:"keywords,omitempty"`

	// Attribute:
	//   kwarray
	// Type:
	//   string
	// Description:
	//   Array of keywords about the site. Only one of ‘keywords’ or
	//   ‘kwarray’ may be present.
	KwArray []string `json:"kwarray,omitempty"`

	// Attribute:
	//   inventorypartnerdomain
	// Type:
	//   string
	// Description:
	//   A domain to be used for inventory authorization in the case of inventory
	//   sharing arrangements between an app owner and content owner. This field
	//   is typically used by authorization crawlers to establish the domain of the
	//   content owner, who has the right to monetize some portion of ad inventory
	//   within the app. The content owner's domain should be listed in the app
	//   owner's app-ads.txt file as an inventorypartnerdomain. Authorization for
	//   supply from the inventorypartnerdomain will be published in the ads.txt
	//   file on the root of that domain. Refer to the ads.txt 1.1 spec for more details.
	InventoryPartnerDomain string `json:"inventorypartnerdomain,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
