package math_test

import (
	"fmt"
	"testing"

	"github.com/miguelbernadi/maroto/v2/pkg/core/entity"

	"github.com/miguelbernadi/maroto/v2/pkg/props"

	"github.com/miguelbernadi/maroto/v2/internal/math"
	"github.com/stretchr/testify/assert"
)

func TestNewMath(t *testing.T) {
	sut := math.New()

	assert.NotNil(t, sut)
	assert.Equal(t, "*math.math", fmt.Sprintf("%T", sut))
}

func TestMath_GetCenterCorrection(t *testing.T) {
	t.Run("should get center correction correctly", func(t *testing.T) {
		// Arrange
		sut := math.New()
		outerSize := 100.0
		innerSize := 50.0

		// Act
		correction := sut.GetCenterCorrection(outerSize, innerSize)

		// Assert
		assert.Equal(t, 25.0, correction)
	})
}

func TestMath_GetInnerCenterCell(t *testing.T) {
	t.Run("there is not side-effect", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 100.0
		inner := &entity.Dimensions{Width: 100, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		_ = sut.GetInnerCenterCell(inner, outer, percent)

		// Assert
		assert.Equal(t, 100.0, inner.Width)
		assert.Equal(t, 100.0, inner.Height)
		assert.Equal(t, 100.0, outer.Width)
		assert.Equal(t, 100.0, outer.Height)
	})
	t.Run("inner same size, inner same proportion, inner 100%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 100.0
		inner := &entity.Dimensions{Width: 100, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer, percent)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})
	t.Run("inner same size, inner same proportion, inner 75%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 75.0
		inner := &entity.Dimensions{Width: 100, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer, percent)

		// Assert
		assert.Equal(t, 12.5, cell.X)
		assert.Equal(t, 12.5, cell.Y)
		assert.Equal(t, 75.0, cell.Width)
		assert.Equal(t, 75.0, cell.Height)
	})
	t.Run("inner smaller, inner same proportion, inner 100%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 100.0
		inner := &entity.Dimensions{Width: 80, Height: 80}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer, percent)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})
	t.Run("inner smaller, inner same proportion, inner 75%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 75.0
		inner := &entity.Dimensions{Width: 80, Height: 80}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer, percent)

		// Assert
		assert.Equal(t, 12.5, cell.X)
		assert.Equal(t, 12.5, cell.Y)
		assert.Equal(t, 75.0, cell.Width)
		assert.Equal(t, 75.0, cell.Height)
	})
	t.Run("inner greater, inner same proportion, inner 100%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 100.0
		inner := &entity.Dimensions{Width: 120, Height: 120}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer, percent)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})
	t.Run("inner greater, inner same proportion, inner 75%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 75.0
		inner := &entity.Dimensions{Width: 120, Height: 120}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer, percent)

		// Assert
		assert.Equal(t, 12.5, cell.X)
		assert.Equal(t, 12.5, cell.Y)
		assert.Equal(t, 75.0, cell.Width)
		assert.Equal(t, 75.0, cell.Height)
	})
	t.Run("inner smaller, inner width proportion greater, 100%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 100.0
		outer := &entity.Dimensions{Width: 100, Height: 100}
		inner := &entity.Dimensions{Width: 100, Height: 80}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer, percent)

		// Assert
		assert.InDelta(t, 0.0, cell.X, 0.1)
		assert.InDelta(t, 10.0, cell.Y, 0.1)
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 80.0, cell.Height)
	})
	t.Run("inner smaller, inner width proportion greater, 75%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 75.0
		outer := &entity.Dimensions{Width: 100, Height: 100}
		inner := &entity.Dimensions{Width: 100, Height: 80}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer, percent)

		// Assert
		assert.Equal(t, 12.5, cell.X)
		assert.Equal(t, 20.0, cell.Y)
		assert.Equal(t, 75.0, cell.Width)
		assert.Equal(t, 60.0, cell.Height)
	})
	t.Run("inner smaller, inner height proportion greater, 100%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 100.0
		inner := &entity.Dimensions{Width: 80, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer, percent)

		// Assert
		assert.Equal(t, 10.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 80.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})
	t.Run("inner smaller, inner height proportion greater, 75%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 75.0
		inner := &entity.Dimensions{Width: 80, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer, percent)

		// Assert
		assert.Equal(t, 20.0, cell.X)
		assert.Equal(t, 12.5, cell.Y)
		assert.Equal(t, 60.0, cell.Width)
		assert.Equal(t, 75.0, cell.Height)
	})
	t.Run("inner greater, inner height proportion greater, 100%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 100.0
		inner := &entity.Dimensions{Width: 100, Height: 125}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer, percent)

		// Assert
		assert.Equal(t, 10.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 80.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})
	t.Run("inner greater, inner height proportion greater, 75%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 75.0
		inner := &entity.Dimensions{Width: 100, Height: 125}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer, percent)

		// Assert
		assert.Equal(t, 20.0, cell.X)
		assert.Equal(t, 12.5, cell.Y)
		assert.Equal(t, 60.0, cell.Width)
		assert.Equal(t, 75.0, cell.Height)
	})
	t.Run("inner greater, inner width proportion greater, 100%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 100.0
		inner := &entity.Dimensions{Width: 125, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer, percent)

		// Assert
		assert.InDelta(t, 0.0, cell.X, 0.1)
		assert.InDelta(t, 10.0, cell.Y, 0.1)
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 80.0, cell.Height)
	})
	t.Run("inner greater, inner width proportion greater, 75%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 75.0
		inner := &entity.Dimensions{Width: 125, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer, percent)

		// Assert
		assert.Equal(t, 12.5, cell.X)
		assert.Equal(t, 20.0, cell.Y)
		assert.Equal(t, 75.0, cell.Width)
		assert.Equal(t, 60.0, cell.Height)
	})
}

func TestMath_GetInnerNonCenterCell(t *testing.T) {
	t.Run("there is not side-effect", func(t *testing.T) {
		// Arrange
		sut := math.New()

		rect := &props.Rect{}
		rect.MakeValid()
		inner := &entity.Dimensions{Width: 100, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		_ = sut.GetInnerNonCenterCell(inner, outer, rect)

		// Assert
		assert.Equal(t, 100.0, inner.Width)
		assert.Equal(t, 100.0, inner.Height)
		assert.Equal(t, 100.0, outer.Width)
		assert.Equal(t, 100.0, outer.Height)
	})
	t.Run("top and left works", func(t *testing.T) {
		// Arrange
		sut := math.New()

		rect := &props.Rect{
			Percent: 50,
			Top:     15,
			Left:    10,
		}
		inner := &entity.Dimensions{Width: 100, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerNonCenterCell(inner, outer, rect)

		// Assert
		assert.Equal(t, 10.0, cell.X)
		assert.Equal(t, 15.0, cell.Y)
		assert.Equal(t, 50.0, cell.Width)
		assert.Equal(t, 50.0, cell.Height)
	})
	t.Run("inner same size, inner same proportion, inner 100%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		rect := &props.Rect{
			Percent: 100,
		}
		rect.MakeValid()
		inner := &entity.Dimensions{Width: 100, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerNonCenterCell(inner, outer, rect)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})
	t.Run("inner same size, inner same proportion, inner 75%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		rect := &props.Rect{
			Percent: 75,
		}
		rect.MakeValid()
		inner := &entity.Dimensions{Width: 100, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerNonCenterCell(inner, outer, rect)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 75.0, cell.Width)
		assert.Equal(t, 75.0, cell.Height)
	})
	t.Run("inner smaller, inner same proportion, inner 100%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		rect := &props.Rect{
			Percent: 100,
		}
		rect.MakeValid()
		inner := &entity.Dimensions{Width: 80, Height: 80}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerNonCenterCell(inner, outer, rect)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})
	t.Run("inner smaller, inner same proportion, inner 75%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		rect := &props.Rect{
			Percent: 75,
		}
		rect.MakeValid()
		inner := &entity.Dimensions{Width: 80, Height: 80}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerNonCenterCell(inner, outer, rect)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 75.0, cell.Width)
		assert.Equal(t, 75.0, cell.Height)
	})
	t.Run("inner greater, inner same proportion, inner 100%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		rect := &props.Rect{
			Percent: 100,
		}
		rect.MakeValid()
		inner := &entity.Dimensions{Width: 120, Height: 120}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerNonCenterCell(inner, outer, rect)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})
	t.Run("inner greater, inner same proportion, inner 75%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		rect := &props.Rect{
			Percent: 75,
		}
		rect.MakeValid()
		inner := &entity.Dimensions{Width: 120, Height: 120}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerNonCenterCell(inner, outer, rect)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 75.0, cell.Width)
		assert.Equal(t, 75.0, cell.Height)
	})
	t.Run("inner smaller, inner width proportion greater, 100%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		rect := &props.Rect{
			Percent: 100,
		}
		rect.MakeValid()
		outer := &entity.Dimensions{Width: 100, Height: 100}
		inner := &entity.Dimensions{Width: 100, Height: 80}

		// Act
		cell := sut.GetInnerNonCenterCell(inner, outer, rect)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 80.0, cell.Height)
	})
	t.Run("inner smaller, inner width proportion greater, 75%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		rect := &props.Rect{
			Percent: 75,
		}
		rect.MakeValid()
		outer := &entity.Dimensions{Width: 100, Height: 100}
		inner := &entity.Dimensions{Width: 100, Height: 80}

		// Act
		cell := sut.GetInnerNonCenterCell(inner, outer, rect)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 75.0, cell.Width)
		assert.Equal(t, 60.0, cell.Height)
	})
	t.Run("inner smaller, inner height proportion greater, 100%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		rect := &props.Rect{
			Percent: 100,
		}
		rect.MakeValid()
		inner := &entity.Dimensions{Width: 80, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerNonCenterCell(inner, outer, rect)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 80.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})
	t.Run("inner smaller, inner height proportion greater, 75%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		rect := &props.Rect{
			Percent: 75,
		}
		rect.MakeValid()
		inner := &entity.Dimensions{Width: 80, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerNonCenterCell(inner, outer, rect)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 60.0, cell.Width)
		assert.Equal(t, 75.0, cell.Height)
	})
	t.Run("inner greater, inner height proportion greater, 100%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		rect := &props.Rect{
			Percent: 100,
		}
		rect.MakeValid()
		inner := &entity.Dimensions{Width: 100, Height: 125}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerNonCenterCell(inner, outer, rect)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 80.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})
	t.Run("inner greater, inner height proportion greater, 75%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		rect := &props.Rect{
			Percent: 75,
		}
		rect.MakeValid()
		inner := &entity.Dimensions{Width: 100, Height: 125}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerNonCenterCell(inner, outer, rect)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 60.0, cell.Width)
		assert.Equal(t, 75.0, cell.Height)
	})
	t.Run("inner greater, inner width proportion greater, 100%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		rect := &props.Rect{
			Percent: 100,
		}
		rect.MakeValid()
		inner := &entity.Dimensions{Width: 125, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerNonCenterCell(inner, outer, rect)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 80.0, cell.Height)
	})
	t.Run("inner greater, inner width proportion greater, 75%", func(t *testing.T) {
		// Arrange
		sut := math.New()

		rect := &props.Rect{
			Percent: 75,
		}
		rect.MakeValid()
		inner := &entity.Dimensions{Width: 125, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerNonCenterCell(inner, outer, rect)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
		assert.Equal(t, 75.0, cell.Width)
		assert.Equal(t, 60.0, cell.Height)
	})
}
