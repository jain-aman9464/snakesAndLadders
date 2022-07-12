package snakes_ladders

type Jumper struct {
	startPoint, endPoint int
}

func NewJumper(start, end int) Jumper {
	return Jumper{
		startPoint: start,
		endPoint:   end,
	}
}
