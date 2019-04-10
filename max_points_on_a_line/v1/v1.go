package v1

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

type direction struct {
	x float64
	y float64
}

func (d *direction) equals(nd *direction) bool {
	return (d.x == nd.x && d.y == nd.y) ||
		(d.x != 0 && d.y != 0 && nd.x/d.x == nd.y/d.y)
}

func (d *direction) string() string {
	return fmt.Sprintf("(%v,%v)", d.x, d.y)
}

type line struct {
	base  Point
	count int
	pm    map[string]Point
	dir   *direction
}

func (l *line) add(p Point, pc map[string]int) bool {
	sub := pString(p)
	if _, ok := l.pm[sub]; ok {
		return false
	}
	l.pm[sub] = p
	l.count += pc[pString(p)]
	return true
}

func (l *line) inline(p Point, pc map[string]int) (bool, int) {
	nd := calDir(l.base, p)
	if l.dir.equals(nd) ||
		(l.base.X == p.X && l.dir.x == 0) ||
		(l.base.Y == p.Y && l.dir.y == 0) {
		return l.add(p, pc), l.count
	}
	return false, 0
}

func newLine(s, t Point, d *direction, pc map[string]int) *line {
	l := &line{
		base: s,
		dir:  d,
		pm:   map[string]Point{},
	}
	l.add(s, pc)
	l.add(t, pc)
	return l
}

func calDir(s, t Point) *direction {
	return &direction{
		x: float64(t.X - s.X),
		y: float64(t.Y - s.Y),
	}
}

func pString(p Point) string {
	return fmt.Sprintf("%v,%v", p.X, p.Y)
}

func maxPoints(points []Point) int {
	max := 0
	unps := []Point{}
	pc := map[string]int{}
	for _, p := range points {
		s := pString(p)
		if _, ok := pc[s]; ok {
			pc[s] += 1
		} else {
			pc[s] = 1
			unps = append(unps, p)
		}
	}
	if len(unps) > 1 {
		lines := []*line{}
		for i := 0; i < len(unps); i += 1 {
			s := unps[i]
			for j := i + 1; j < len(unps); j += 1 {
				t := unps[j]
				inline := false
				for _, l := range lines {
					if v, count := l.inline(t, pc); v {
						inline = true
						if count > max {
							max = count
						}
						break
					}
				}
				if !inline {
					nl := newLine(s, t, calDir(s, t), pc)
					lines = append(lines, nl)
					if nl.count > max {
						max = nl.count
					}
				}
			}
		}
	} else {
		max = len(points)
	}
	return max
}
