package main

type dsu struct {
	parent []int
	rank   []int
}

func newDSU(n int) *dsu {
	obj := &dsu{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		obj.parent[i] = i
		obj.rank[i] = 0
	}
	return obj
}

func (d *dsu) FindParent(x int) int {
	if d.parent[x] == x {
		return d.parent[x]
	}
	d.parent[x] = d.FindParent(d.parent[x])
	return d.parent[x]
}

func (d *dsu) Union(x, y int) {
	parx := d.FindParent(x)
	pary := d.FindParent(y)
	if parx == pary {
		return
	}
	switch {
	case d.rank[parx] < d.rank[pary]:
		d.parent[parx] = pary
	case d.rank[parx] > d.rank[pary]:
		d.parent[pary] = parx
	default: // When both ranks are equal
		d.parent[pary] = parx
		d.rank[parx]++
	}
}
