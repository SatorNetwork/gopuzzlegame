package main

import (
	"github.com/SatorNetwork/gopuzzlegame/util"
	"math"
)

type Puzzle struct {
	tiles []*Tile

}

func (p *Puzzle) GetDimension() int {
	return int(math.Sqrt(float64(len(p.tiles))))
}

func (p *Puzzle) GetWhitespaceTile() *Tile {
	for _, tile := range p.tiles {
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

	for _, tile := range p.tiles {
		if tile.CurrentPosition.X == whitespaceTile.CurrentPosition.X + relativeOffset.Dx &&
			tile.CurrentPosition.Y == whitespaceTile.CurrentPosition.Y + relativeOffset.Dy {
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
	for _, tile := range p.tiles {
		if tile != whitespaceTile && tile.CurrentPosition == tile.CorrectPosition {
			numberOfCorrectTiles++
		}
	}

	return numberOfCorrectTiles
}

func (p *Puzzle) IsComplete() bool {
	return (len(p.tiles) - 1) - p.GetNumberOfCorrectTiles() == 0
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
	for a := 0; a < len(p.tiles); a++ {
		tileA := p.tiles[a]
		if tileA.IsWhitespace { continue }
		for b := a + 1; b < len(p.tiles); b++ {
			tileB := p.tiles[b]
			if p.isInversion(tileA, tileB) {
				count++
			}
		}
	}
	return count
}

func (p *Puzzle) IsSolvable() bool {
	size := p.GetDimension()
	height := len(p.tiles)/size
	if size * height != len(p.tiles) {
		panic("tiles must be equal to size * height")
	}

	inversions := p.CountInversions()
	if size % 2 == 1 {
		return inversions % 2 == 0
	}

	whitespace := p.GetWhitespaceTile()
	whitespaceRow := whitespace.CurrentPosition.Y

	if (height - whitespaceRow + 1) % 2 == 1 {
		return inversions % 2 == 0
	} else {
		return inversions % 2 == 1
	}
}

func (p *Puzzle) MoveTiles(tile *Tile, tilesToSwap []*Tile) *Puzzle {
	whitespaceTile := p.GetWhitespaceTile()
	if whitespaceTile == nil {
		return nil
	}

	dx := whitespaceTile.CurrentPosition.X - tile.CurrentPosition.X
	dy := whitespaceTile.CurrentPosition.Y - tile.CurrentPosition.Y

	if math.Abs(float64(dx)) + math.Abs(float64(dy)) > 1 {
		shiftPointX := tile.CurrentPosition.X + util.GetSign(dx)
		shiftPointY := tile.CurrentPosition.Y + util.GetSign(dx)
		var tileToSwapWith *Tile
		for _, t := range p.tiles {
			if t.CurrentPosition.X == shiftPointX && t.CurrentPosition.Y == shiftPointY {
				tileToSwapWith = t
			}
		}
		tilesToSwap = append(tilesToSwap, tileToSwapWith)
		return p.MoveTiles(tileToSwapWith, tilesToSwap)
	} else {
		tilesToSwap = append(tilesToSwap, tile)
		return p.SwapTiles(tilesToSwap)
	}
}

func (p *Puzzle) SwapTiles(tilesToSwap []*Tile) *Puzzle {
	Reverse(tilesToSwap)
	for _, tileToSwap := range tilesToSwap {
		tileIndex := IndexOfTileInTiles(p.tiles, tileToSwap)
		tile := p.tiles[tileIndex]
		whitespaceTile := p.GetWhitespaceTile()
		if whitespaceTile == nil {
			return nil
		}
		whitespaceTileIndex := IndexOfTileInTiles(p.tiles, whitespaceTile)

		p.tiles[tileIndex] = tile.CopyWith(whitespaceTile.CurrentPosition)
		p.tiles[whitespaceTileIndex] = whitespaceTile.CopyWith(tile.CurrentPosition)
	}

	return &Puzzle{tiles: p.tiles}
}

