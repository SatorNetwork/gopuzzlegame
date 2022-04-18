package main

func main() {
	var images [][]byte
	for i := 0; i < 16; i++ {
		images = append(images, []byte{})
	}
	puzzle := GeneratePuzzle(images, 4, true)
	puzzle.IsSolvable()
}