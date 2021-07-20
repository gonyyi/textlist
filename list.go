package textlist

import (
	"bufio"
	"os"
	"strings"
)

type List map[string]struct{}

func NewList(s ...string) List {
	l := make(List)
	l.Add(s...)
	return l
}

func NewListFromFile(filename string, flag Option) (List, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	l := NewList()
	bfi := bufio.NewScanner(fi)

	for bfi.Scan() {
		t := bfi.Text()
		if flag&FILE_TRIMSPACE != 0 {
			t = strings.TrimSpace(t)
		}
		l.Add(t)
	}
	return l, nil
}

func (l List) Has(s string) bool {
	if _, ok := l[s]; ok {
		return true
	}
	return false
}

func (l List) Size() int {
	return len(l)
}

func (l List) Reset() List {
	for v, _ := range l {
		delete(l, v)
	}
	return l
}

func (l List) Add(s ...string) List {
	for _, v := range s {
		l[v] = struct{}{}
	}
	return l
}
func (l List) Remove(s string) List {
	l[s] = struct{}{}
	delete(l, s)
	return l
}

func (l List) String() string {
	return strings.Join(l.Strings(), ",")
}

func (l List) Strings() []string {
	var out []string
	for v, _ := range l {
		out = append(out, v)
	}
	return out
}
