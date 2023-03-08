package adcom1

import (
	"github.com/pinpr/backend-shared-kit/vld"
	"golang.org/x/exp/slices"
)

// CategoryTaxonomy identifies the taxonomy in effect when content categories are listed.
type CategoryTaxonomy int
type CategoryTaxonomyEnum string

// CategoryTaxonomy options.
//
// Values of 500+ hold vendor-specific codes.
const (
	CatTaxIABContent10  CategoryTaxonomyEnum = "content1.0"  // 1	IAB Content Category Taxonomy 1.0.
	CatTaxIABContent20  CategoryTaxonomyEnum = "content2.0"  // 2	IAB Content Category Taxonomy 2.0: www.iab.com/guidelines/taxonomy
	CatTaxIABProduct10  CategoryTaxonomyEnum = "product1.0"  // 3	IAB Ad Product Taxonomy 1.0.
	CatTaxIABAudience11 CategoryTaxonomyEnum = "audience1.1" // 4	IAB Audience Taxonomy 1.1.
	CatTaxIABContent21  CategoryTaxonomyEnum = "content2.1"  // 5	IAB Content Category Taxonomy 2.1.
	CatTaxIABContent22  CategoryTaxonomyEnum = "content2.2"  // 6	IAB Content Category Taxonomy 2.2
)

var allCatTaxes = []CategoryTaxonomyEnum{
	CatTaxIABContent10,
	CatTaxIABContent20,
	CatTaxIABProduct10,
	CatTaxIABAudience11,
	CatTaxIABContent21,
	CatTaxIABContent22,
}

func (c CategoryTaxonomyEnum) IsValid() error {
	return vld.EnumIsValid(&c, allCatTaxes)
}

func (c CategoryTaxonomyEnum) ToValue() CategoryTaxonomy {
	return CategoryTaxonomy(slices.Index(allCatTaxes, c) + 1)
}

func (e CategoryTaxonomy) ToValue() CategoryTaxonomyEnum {
	return allCatTaxes[e-1]
}
