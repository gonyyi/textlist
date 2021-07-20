package textlist

func Compare(from List, to List) (added []string, removed []string) {
	for v, _ := range from {
		if !to.Has(v) {
			removed = append(removed, v)
		}
	}

	for v, _ := range to {
		if !from.Has(v) {
			added = append(added, v)
		}
	}
	return
}
