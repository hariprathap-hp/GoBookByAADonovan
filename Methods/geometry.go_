package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
	fmt.Println(p.X, p.Y)
}

func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			fmt.Println(p[i-1], p[i])
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}

func main() {
	p := Point{2, 5}
	q := Point{5, 7}

	mypath := Path{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	fmt.Println(p.Distance(q))
	fmt.Println(Distance(p, q))
	fmt.Println(mypath.Distance())
	(&p).ScaleBy(5.0)
	fmt.Println(p.X, p.Y)

}
