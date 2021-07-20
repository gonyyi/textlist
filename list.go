package textlist

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

type List map[string]struct{}

func NewList(s ...string) List {
	l := make(List)
	l.Add(s...)
	return l
}

func NewListFromFile(filename, delim string, flag Option) (List, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	l := NewList()
	bfi := bufio.NewScanner(fi)

	bfi.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		// bytes.IndexAny will take string, but if any of the string matches, it will trigger.
		// therefore, can't use it.
		if i := bytes.Index(data, []byte(delim)); i >= 0 {
			return i + len(delim), data[0:i], nil
		}
		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return len(data), data, nil
		}
		// Request more data.
		return 0, nil, nil
	})

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

func (l List) Remove(s ...string) List {
	for _, v := range s {
		delete(l, v)
	}
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
