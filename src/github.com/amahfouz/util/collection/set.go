package collection

type StringSet struct {
    items map[string] bool
}

func (set *StringSet) Add(s string) {
    set.items[s] = true
}

func (set *StringSet) Contains(s string) bool {
    return set.items[s]
}

func (set *StringSet) NumElements() int {
    return len(set.items)
}

func (set StringSet) Union(other StringSet) StringSet {
	union := NewStringSet()
	
	for k,_ := range set.items {
	    union.Add(k)
	}
	for k,_ := range other.items {
	    union.Add(k)
	}
	return union
}

func (set StringSet) Intersect(other StringSet) StringSet {
    intersection := NewStringSet()
    
    var smaller, bigger StringSet
    if set.NumElements() < other.NumElements() {
        smaller = set
        bigger = other
    } else {
        smaller = other
        bigger = set
    }
    
    for k, _ := range smaller.items {
        if bigger.Contains(k) {
			intersection.Add(k)            
        }
    }
    
    return intersection
}

func NewStringSet() StringSet {
    return StringSet{make(map[string]bool)}
}

