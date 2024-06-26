package math

import (
	"github.com/miguelbernadi/maroto/v2/pkg/core/entity"
	"github.com/miguelbernadi/maroto/v2/pkg/props"
)

type math struct{}

// New create a Math.
func New() *math {
	return &math{}
}

// GetInnerCenterCell define a inner cell formatted inside outer cell centered.
func (s *math) GetInnerCenterCell(inner *entity.Dimensions, outer *entity.Dimensions, percent float64) *entity.Cell {
	percent /= 100.0

	innerProportion := inner.Height / inner.Width
	outerProportion := outer.Height / outer.Width

	innerCell := &entity.Cell{}
	if innerProportion > outerProportion {
		newInnerWidth := outer.Height / innerProportion * percent
		newInnerHeight := newInnerWidth * innerProportion

		widthCorrection := s.GetCenterCorrection(outer.Width, newInnerWidth)
		heightCorrection := s.GetCenterCorrection(outer.Height, newInnerHeight)

		innerCell.X = widthCorrection
		innerCell.Y = heightCorrection
		innerCell.Width = newInnerWidth
		innerCell.Height = newInnerHeight
	} else {
		newInnerWidth := outer.Width * percent
		newInnerHeight := newInnerWidth * innerProportion

		widthCorrection := s.GetCenterCorrection(outer.Width, newInnerWidth)
		heightCorrection := s.GetCenterCorrection(outer.Height, newInnerHeight)

		innerCell.X = widthCorrection
		innerCell.Y = heightCorrection
		innerCell.Width = newInnerWidth
		innerCell.Height = newInnerHeight
	}

	return innerCell
}

// GetInnerNonCenterCell define a inner cell formatted inside outer cell non centered.
func (s *math) GetInnerNonCenterCell(inner *entity.Dimensions, outer *entity.Dimensions, prop *props.Rect) *entity.Cell {
	percent := prop.Percent / 100.0

	innerProportion := inner.Height / inner.Width
	outerProportion := outer.Height / outer.Width

	innerCell := &entity.Cell{}
	if innerProportion > outerProportion {
		newInnerWidth := outer.Height / innerProportion * percent
		newInnerHeight := newInnerWidth * innerProportion

		innerCell.X = prop.Left
		innerCell.Y = prop.Top
		innerCell.Width = newInnerWidth
		innerCell.Height = newInnerHeight
	} else {
		newInnerWidth := outer.Width * percent
		newInnerHeight := newInnerWidth * innerProportion

		innerCell.X = prop.Left
		innerCell.Y = prop.Top
		innerCell.Width = newInnerWidth
		innerCell.Height = newInnerHeight
	}

	return innerCell
}

// GetCenterCorrection return the correction of space in X or Y to
// centralize a line in relation with another line.
func (s *math) GetCenterCorrection(outerSize, innerSize float64) float64 {
	return (outerSize - innerSize) / 2.0
}
