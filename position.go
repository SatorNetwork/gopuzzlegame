package main

type Position struct {
	X int
	Y int
}

func (p *Position) CompareTo(other Position) int {
	if p.Y < other.Y {
		return -1
	} else if p.Y > other.Y {
		return 1
	} else {
		if p.X < other.X {
			return -1
		} else if p.X > other.X {
			return 1
		} else {
			return 0
		}
	}
}

func (p *Position) CompareToBool(other Position) bool {
	if p.Y < other.Y {
		return false
	} else if p.Y > other.Y {
		return true
	} else {
		if p.X < other.X {
			return false
		} else if p.X > other.X {
			return true
		} else {
			return false
		}
	}
}

func ShufflePositions(positions []Position) {

}