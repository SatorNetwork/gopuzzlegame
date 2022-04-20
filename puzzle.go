package gopuzzlegame

import (
	"math"
	"sort"

	"github.com/SatorNetwork/gopuzzlegame/util"
)

type Puzzle struct {
	Tiles []*Tile
}

func (p *Puzzle) GetDimension() int {
	return int(math.Sqrt(float64(len(p.Tiles))))
}

func (p *Puzzle) GetWhitespaceTile() *Tile {
	for _, tile := range p.Tiles {
		if tile.IsWhitespace {
			return tile
		}
	}
	return nil
}

func (p *Puzzle) GetTileRelativeToWhitespaceTile(relativeOffset Offset) *Tile {
	whitespaceTile := p.GetWhitespaceTile()
	if whitespaceTile == nil {
		return nil
	}

	for _, tile := range p.Tiles {
		if tile.CurrentPosition.X == whitespaceTile.CurrentPosition.X+relativeOffset.Dx &&
			tile.CurrentPosition.Y == whitespaceTile.CurrentPosition.Y+relativeOffset.Dy {
			return tile
		}
	}

	return nil
}

func (p *Puzzle) GetNumberOfCorrectTiles() int {
	whitespaceTile := p.GetWhitespaceTile()
	if whitespaceTile == nil {
		return 0
	}

	numberOfCorrectTiles := 0
	for _, tile := range p.Tiles {
		if tile != whitespaceTile && tile.CurrentPosition == tile.CorrectPosition {
			numberOfCorrectTiles++
		}
	}

	return numberOfCorrectTiles
}

func (p *Puzzle) IsComplete() bool {
	return (len(p.Tiles)-1)-p.GetNumberOfCorrectTiles() == 0
}

func (p *Puzzle) IsTileMovable(tile *Tile) bool {
	whitespaceTile := p.GetWhitespaceTile()
	if whitespaceTile == nil {
		return false
	}

	if tile == whitespaceTile {
		return false
	}

	if whitespaceTile.CurrentPosition.X != tile.CurrentPosition.X &&
		whitespaceTile.CurrentPosition.Y != tile.CorrectPosition.Y {
		return false
	}

	return true
}

func (p *Puzzle) isInversion(a, b *Tile) bool {
	if !b.IsWhitespace && a.Value != b.Value {
		if b.Value < a.Value {
			return b.CurrentPosition.CompareTo(a.CurrentPosition) > 0
		} else {
			return a.CurrentPosition.CompareTo(b.CurrentPosition) > 0
		}
	}
	return false
}

func (p *Puzzle) CountInversions() int {
	count := 0
	for a := 0; a < len(p.Tiles); a++ {
		tileA := p.Tiles[a]
		if tileA.IsWhitespace {
			continue
		}
		for b := a + 1; b < len(p.Tiles); b++ {
			tileB := p.Tiles[b]
			if p.isInversion(tileA, tileB) {
				count++
			}
		}
	}
	return count
}

func (p *Puzzle) IsSolvable() bool {
	size := p.GetDimension()
	height := len(p.Tiles) / size
	if size*height != len(p.Tiles) {
		panic("tiles must be equal to size * height")
	}

	inversions := p.CountInversions()
	if size%2 == 1 {
		return inversions%2 == 0
	}

	whitespace := p.GetWhitespaceTile()
	whitespaceRow := whitespace.CurrentPosition.Y

	if (height-whitespaceRow+1)%2 == 1 {
		return inversions%2 == 0
	} else {
		return inversions%2 == 1
	}
}

func (p *Puzzle) MoveTiles(tile *Tile, tilesToSwap []*Tile) *Puzzle {
	whitespaceTile := p.GetWhitespaceTile()
	if whitespaceTile == nil {
		return nil
	}

	dx := whitespaceTile.CurrentPosition.X - tile.CurrentPosition.X
	dy := whitespaceTile.CurrentPosition.Y - tile.CurrentPosition.Y

	if math.Abs(float64(dx))+math.Abs(float64(dy)) > 1 {
		shiftPointX := tile.CurrentPosition.X + util.GetSign(dx)
		shiftPointY := tile.CurrentPosition.Y + util.GetSign(dy)
		var tileToSwapWith *Tile
		for _, t := range p.Tiles {
			if t.CurrentPosition.X == shiftPointX && t.CurrentPosition.Y == shiftPointY {
				tileToSwapWith = t
				break
			}
		}
		tilesToSwap = append(tilesToSwap, tile)
		return p.MoveTiles(tileToSwapWith, tilesToSwap)
	} else {
		tilesToSwap = append(tilesToSwap, tile)
		return p.SwapTiles(tilesToSwap)
	}
}

func (p *Puzzle) SwapTiles(tilesToSwap []*Tile) *Puzzle {
	Reverse(tilesToSwap)
	for _, tileToSwap := range tilesToSwap {
		tileIndex := IndexOfTileInTiles(p.Tiles, tileToSwap)
		tile := p.Tiles[tileIndex]
		whitespaceTile := p.GetWhitespaceTile()
		if whitespaceTile == nil {
			return nil
		}
		whitespaceTileIndex := IndexOfTileInTiles(p.Tiles, whitespaceTile)

		p.Tiles[tileIndex] = tile.CopyWith(whitespaceTile.CurrentPosition)
		p.Tiles[whitespaceTileIndex] = whitespaceTile.CopyWith(tile.CurrentPosition)
	}

	return &Puzzle{Tiles: p.Tiles}
}

func (p *Puzzle) Sort() {
	sort.Slice(p.Tiles, func(i, j int) bool {
		return p.Tiles[i].CurrentPosition.CompareToBool(p.Tiles[j].CurrentPosition)
	})
}
