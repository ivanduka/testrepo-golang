package main

// Square represents the rectangular shape with both sides being equal
type Square struct {
	width, height int
}

// Area returns an area of the square
func (s *Square) Area() int {
	return w.width * w.height
}
