package shapes

import "math"

type point struct {
        x float64
        y float64
}

type circle struct {
        center point
        radius float64
}

func (c circle) Square() float64 {
        return math.Pi * c.radius * c.radius
}

type rectangle struct {
        p1 point
        p2 point
}

func (r rectangle) Square() float64 {
    return math.Abs(r.p1.x - r.p2.x) * math.Abs(r.p1.y - r.p2.y)
}
