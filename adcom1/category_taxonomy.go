package adcom1

// CategoryTaxonomy identifies the taxonomy in effect when content categories are listed.
type CategoryTaxonomy int64

// CategoryTaxonomy options.
//
// Values of 500+ hold vendor-specific codes.
const (
	CatTaxIABContent10  CategoryTaxonomy = 1 // IAB Content Category Taxonomy 1.0.
	CatTaxIABContent20  CategoryTaxonomy = 2 // IAB Content Category Taxonomy 2.0: www.iab.com/guidelines/taxonomy
	CatTaxIABProduct10  CategoryTaxonomy = 3 // IAB Ad Product Taxonomy 1.0.
	CatTaxIABAudience11 CategoryTaxonomy = 4 // IAB Audience Taxonomy 1.1.
	CatTaxIABContent21  CategoryTaxonomy = 5 // IAB Content Category Taxonomy 2.1.
	CatTaxIABContent22  CategoryTaxonomy = 6 // IAB Content Category Taxonomy 2.2.
	CatTaxIABContent30  CategoryTaxonomy = 7 // IAB Content Category Taxonomy 3.0.
)
