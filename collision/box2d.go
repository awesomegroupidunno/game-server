package collision

type Box2d interface {
	Position() (float64, float64)
	Size() (float64, float64)
	AngleDegrees() float64
}

func naiveCollision(b, t Box2d) bool {
	box1x, box1y := b.Position()
	box1w, box1h := b.Size()
	//box1a := b.AngleDegrees()

	box2x, box2y := t.Position()
	box2w, box2h := t.Size()
	//box2a := t.AngleDegrees()

	return box1x < box2x+box2w &&
		box1x+box1w > box2x &&
		box1y < box2y+box2h &&
		box1h+box1y > box2y
}

func Collides(b, t Box2d) bool {
	return naiveCollision(b, t)
}
