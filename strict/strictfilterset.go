package strict

import "fmt"

type FilterSet struct {
	filters map[string]struct{}
}

func SayHello() {
	fmt.Println("Hello Go!")
}

// NewFilterSet constructs a FilterSet of exact string matches.
func NewFilterSet(filters []string) *FilterSet {
	fs := &FilterSet{
		filters: make(map[string]struct{}, len(filters)),
	}

	for _, f := range filters {
		fs.filters[f] = struct{}{}
	}

	return fs
}

// Matches returns true if the given string matches any of the FilterSet's filters.
func (sfs *FilterSet) Matches(toMatch string) bool {
	_, ok := sfs.filters[toMatch]
	return ok
}
