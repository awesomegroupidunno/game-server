package collision

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

	points := [4]point{}

	points[0] = point{X: cen_x - width/2,
		Y: cen_y - height/2}

	points[1] = point{X: cen_x + width/2,
		Y: cen_y - height/2}

	points[2] = point{X: cen_x - width/2,
		Y: cen_y + height/2}

	points[3] = point{X: cen_x + width/2,
		Y: cen_y + height/2}
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

			normal := point{X: p2.Y - p1.Y, Y: p1.X - p2.X}

			var minA, maxA float64
			var minAnill, maxAnill bool
			minAnill = true
			maxAnill = true
			for _, p := range a.Points() {
				projected := normal.X*p.X + normal.Y*p.Y
				if minAnill == true || projected < minA {
					minA = projected
					minAnill = false
				}
				if maxAnill == true || projected > maxA {
					maxA = projected
					maxAnill = false
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

	return true
}

func Collides(b, t Box2d) bool {
	return isPolygonIntersect(boxToPoly(b), boxToPoly(t))
}
