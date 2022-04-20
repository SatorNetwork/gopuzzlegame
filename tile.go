package gopuzzlegame

type Tile struct {
	Image           []byte
	Value           int
	CorrectPosition Position
	CurrentPosition Position
	IsWhitespace    bool
}

func Reverse(tiles []*Tile) {
	for i, j := 0, len(tiles)-1; i < j; i, j = i+1, j-1 {
		tiles[i], tiles[j] = tiles[j], tiles[i]
	}
}

func IndexOfTileInTiles(tiles []*Tile, tile *Tile) int {
	for i := range tiles {
		if tiles[i] == tile {
			return i
		}
	}
	return -1
}

func (t *Tile) CopyWith(currentPosition Position) *Tile {
	return &Tile{
		Image:           t.Image,
		Value:           t.Value,
		CorrectPosition: t.CorrectPosition,
		CurrentPosition: currentPosition,
		IsWhitespace:    t.IsWhitespace,
	}
}
