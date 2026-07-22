package adcom1

import "encoding/json"

// Data and segment objects together allow additional data about the related object (e.g., user, content) to be specified.
// This data may be from multiple sources whether from the exchange itself or third parties as specified by the id attribute.
// When in use, vendor-specific IDs should be communicated a priori among the parties.
type Data struct {
	// Attribute:
	//   id
	// Type:
	//   string
	// Definition:
	//   Vendor-specific ID for the data provider.
	ID string `json:"id,omitempty"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Definition:
	//   Vendor-specific displayable name for the data provider.
	Name string `json:"name,omitempty"`

	// Attribute:
	//   cids
	// Type:
	//   string array
	// Definition:
	//   Array of Extended Content IDs, representing one or more identifiers for the video or audio content from the ID source specified in the 'name' field of the 'data' object.
	CIDs []string `json:"cids,omitempty"`

	// Attribute:
	//   segment
	// Type:
	//   object array
	// Definition:
	//   Array of Segment objects that contain the actual data values.
	//   Refer to Object: Segment.
	Segment []Segment `json:"segment,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Definition:
	//   Optional vendor-specific extensions.
	Ext json.RawMessage `json:"ext,omitempty"`
}
