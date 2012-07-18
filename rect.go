package tulib

type Rect struct {
	X int
	Y int
	Width int
	Height int
}

func (a Rect) Intersection(b Rect) Rect {
	// no intersection cases
	if a.X >= b.Width || a.Y >= b.Height {
		return Rect{0, 0, 0, 0}
	}

	if a.X + a.Width <= 0 || a.Y + a.Height <= 0 {
		return Rect{0, 0, 0, 0}
	}

	// adjust X or Width
	if a.X + a.Width >= b.Width {
		a.Width = b.Width - a.X
	}

	if a.X < 0 {
		a.Width += a.X
		a.X = 0
	}

	// adjust Y or Height
	if a.Y + a.Height >= b.Height {
		a.Height = b.Height - a.Y
	}

	if a.Y < 0 {
		a.Height += a.Y
		a.Y = 0
	}

	return a
}



