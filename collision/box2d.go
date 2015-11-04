package collision

import (
	"math"
)

// A simple interface for rotatable 2d rectangles
type Box2d interface {
	Position() (float64, float64)
	Size() (float64, float64)
	AngleDegrees() float64
}

type point struct {
	X float64
	Y float64
}

type boxPoly struct {
	points []point
}

type polygon interface {
	Points() []point
}

// converts a Box2d to a polygon
func boxToPoly(b Box2d) polygon {
	cen_x, cen_y := b.Position()
	width, height := b.Size()
	angle := b.AngleDegrees()

	points := [4]point{}

	points[0] = point{X: getX(-width/2, -height/2, angle) + cen_x,
		Y: getY(-width/2, -height/2, angle) + cen_y}

	points[1] = point{X: getX(width/2, -height/2, angle) + cen_x,
		Y: getY(width/2, -height/2, angle) + cen_y}

	points[2] = point{X: getX(-width/2, height/2, angle) + cen_x,
		Y: getY(-width/2, height/2, angle) + cen_y}

	points[3] = point{X: getX(width/2, height/2, angle) + cen_x,
		Y: getY(width/2, height/2, angle) + cen_y}
	z := boxPoly{points: points[:4]}

	return z
}

//
func (b boxPoly) Points() []point {
	return b.points
}

// Determines if 2 polygons intersect
// original algorithm can be found here: http://stackoverflow.com/questions/10962379/how-to-check-intersection-between-2-rotated-rectangles
func isPolygonIntersect(a, b polygon) bool {
	shapes := [2]polygon{a, b}
	for _, polygon := range shapes {
		for i1 := 0; i1 < len(polygon.Points()); i1++ {

			i2 := (i1 + 1) % len(polygon.Points())
			p1 := polygon.Points()[i1]
			p2 := polygon.Points()[i2]

			normal := point{Y: p2.Y - p1.Y, X: p1.X - p2.X}

			minA := math.MaxFloat64
			maxA := (math.MaxFloat64 * -1)
			for _, p := range a.Points() {
				projected := normal.X*p.X + normal.Y*p.Y
				if projected < minA {
					minA = projected
				}
				if projected > maxA {
					maxA = projected
				}
			}

			minB := math.MaxFloat64
			maxB := (math.MaxFloat64 * -1)
			for _, p := range b.Points() {
				projected := normal.X*p.X + normal.Y*p.Y
				if projected < minB {
					minB = projected
				}
				if projected > maxB {
					maxB = projected
				}
			}

			if maxA < minB || maxB < minA {
				return false
			}
		}
	}
	return true
}

// Checks 2 Box2d Objects for collisions
// returns true if they collide
func Collides(b, t Box2d) bool {
	return isPolygonIntersect(boxToPoly(b), boxToPoly(t))
}

// Returns the Y value of the endpoint
func getY(x, y, theta float64) float64 {
	theta = theta * math.Pi / 180
	return x*math.Sin(theta) + y*math.Cos(theta)
}

// Returns the Y value of the endpoint
func getX(x, y, theta float64) float64 {
	theta = theta * math.Pi / 180
	return x*math.Cos(theta) - y*math.Sin(theta)
}
