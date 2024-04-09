package adcom1

// PodDedupe
//
// Types of pod deduplication settings.
type PodDedupe int8

const (
	PodDedupeADomain      PodDedupe = 1
	PodDedupeIABCategory  PodDedupe = 2
	PodDedupeCreativeID   PodDedupe = 3
	PodDedupeMediafileURL PodDedupe = 4
)
