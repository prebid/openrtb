package adcom1

// DOOHVenueTaxonomy describes the locations and contexts in which Out-Of-Home media may be experienced. Taxonomies entries
// are expected to refer to a specific version, unless a given taxonomy has explicit semantics for forward compatibility and
// handling updates.
type DOOHVenueTaxonomy int

// Digital out-of-home venue taxonomies.
const (
	VenueTaxonomyAdCom      DOOHVenueTaxonomy = 0 // AdCom DOOH Venue Types (deprecated)
	VenueTaxonomyOpenOOH10  DOOHVenueTaxonomy = 1 // OpenOOH Venue Taxonomy 1.0 https://github.com/openooh/venue-taxonomy/blob/main/specification-1.0.md
	VenueTaxonomyDPAA       DOOHVenueTaxonomy = 2 // DPAA Device Venue Types https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/master/AdCOM%20v1.0%20FINAL.md#list--dooh-venue-types-
	VenueTaxonomyDMI11      DOOHVenueTaxonomy = 3 // DMI Categorization of Venues 1.1 https://www.dmi-org.com/download/DMI_Standards_for_DOOH_Venues.pdf
	VenueTaxonomyOMAJan2022 DOOHVenueTaxonomy = 4 // OMA taxonomy Jan 2022 https://www.oma.org.au/industry-standards
	VenueTaxonomyOpenOOH11  DOOHVenueTaxonomy = 5 // OpenOOH Venue Taxonomy 1.1 https://github.com/openooh/venue-taxonomy/blob/main/specification-1.1.md
)

// Ptr returns pointer to own value.
func (t DOOHVenueTaxonomy) Ptr() *DOOHVenueTaxonomy {
	return &t
}

// Val safely dereferences pointer, returning default value (ConnectionUnknown) for nil.
func (t *DOOHVenueTaxonomy) Val() DOOHVenueTaxonomy {
	if t == nil {
		return VenueTaxonomyOpenOOH10
	}
	return *t
}
