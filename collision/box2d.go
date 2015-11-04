package collision

import (
	"log"
	"math"
)

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

func boxToPoly(b Box2d) boxPoly {
	cen_x, cen_y := b.Position()
	width, height := b.Size()
	angle := b.AngleDegrees()

	points := [4]point{}

	points[0] = point{X: getX(-width/2, -height/2, angle) + cen_x,
		Y: getY(-width/2, -height/2, angle) + cen_y}

	points[1] = point{X: getX(width/2, -height/2, angle) + cen_x,
		Y: getY(width/2, -height/2, angle) + cen_y}

	points[2] = point{X: getX(-width/2, +height/2, angle) + cen_x,
		Y: getY(-width/2, height/2, angle) + cen_y}

	points[3] = point{X: getX(width/2, height/2, angle) + cen_x,
		Y: getY(width/2, height/2, angle) + cen_y}
	z := boxPoly{points: points[:4]}

	return z
}

func (b boxPoly) Points() []point {
	return b.points
}

func isPolygonIntersect(a, b polygon) bool {
	shapes := [2]polygon{a, b}
	for _, polygon := range shapes {
		for i1 := 0; i1 < len(polygon.Points()); i1++ {

			i2 := (i1 + 1) % len(polygon.Points())
			p1 := polygon.Points()[i1]
			p2 := polygon.Points()[i2]

			normal := point{Y: p2.Y - p1.Y, X: p1.X - p2.X}

			var minA, maxA float64
			minAnil := true
			maxAnil := true
			for _, p := range a.Points() {
				projected := normal.X*p.X + normal.Y*p.Y
				if minAnil == true || projected < minA {
					minA = projected
					minAnil = false
				}
				if maxAnil == true || projected > maxA {
					maxA = projected
					maxAnil = false
				}
			}

			var minB, maxB float64
			minBnil := true
			maxBnil := true
			for _, p := range b.Points() {
				projected := normal.X*p.X + normal.Y*p.Y
				if minBnil == true || projected < minB {
					minB = projected
					minBnil = false
				}
				if maxBnil == true || projected > maxB {
					maxB = projected
					maxBnil = false
				}
			}

			if maxA < minB || maxB < minA {
				return false
			}
		}
	}
	log.Println(a)
	log.Println(b)
	return true
}

func Collides(b, t Box2d) bool {
	return isPolygonIntersect(boxToPoly(b), boxToPoly(t))
}

func getY(x, y, theta float64) float64 {
	theta = theta * math.Pi / 180
	return x*math.Sin(theta) + y*math.Cos(theta)
}
func getX(x, y, theta float64) float64 {
	theta = theta * math.Pi / 180
	return x*math.Cos(theta) - y*math.Sin(theta)
}
