package sets

type Element interface {
	Key() string
	Equals(interface{}) bool
}

type Set []Element

func Diff(a, b Set) (Set, Set) {
	ma := make(map[string]Element)
	for _, e := range a {
		n := string(e.Key())
		if _, ok := ma[n]; !ok {
			ma[n] = e
		}
	}

	mb := make(map[string]Element)
	for _, e := range b {
		n := string(e.Key())
		if _, ok := mb[n]; !ok {
			mb[n] = e
		}
	}

	var del Set
	for k, v := range ma {
		if _, ok := mb[k]; !ok {
			del = append(del, v)
		}
	}

	var add Set
	for k, v := range mb {
		if _, ok := ma[k]; !ok {
			add = append(add, v)
		}
	}

	return add, del
}

type stringElement string

func (e stringElement) Key() string { return string(e) }

func (e stringElement) Equals(v interface{}) bool {
	if c, ok := v.(stringElement); ok && c == e {
		return true
	} else {
		return false
	}
}

func DiffStrings(a, b []string) ([]string, []string) {
	la, lb := len(a), len(b)

	ea := make(Set, la)
	for i := 0; i < la; i++ {
		ea[i] = stringElement(a[i])
	}

	eb := make(Set, lb)
	for i := 0; i < lb; i++ {
		eb[i] = stringElement(b[i])
	}

	add, del := Diff(ea, eb)

	adds := make([]string, len(add))
	for i, e := range add {
		adds[i] = string(e.(stringElement))
	}

	dels := make([]string, len(del))
	for i, e := range del {
		dels[i] = string(e.(stringElement))
	}

	return adds, dels
}
