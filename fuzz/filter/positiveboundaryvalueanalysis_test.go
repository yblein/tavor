package filter

import (
	"testing"

	. "github.com/stretchr/testify/assert"

	"github.com/zimmski/tavor/token"
	"github.com/zimmski/tavor/token/primitives"
)

func TestNewPositiveBoundaryValueAnalysisFilterToBeFilter(t *testing.T) {
	var filt *Filter

	Implements(t, filt, &PositiveBoundaryValueAnalysisFilter{})
}

func TestNewPositiveBoundaryValueAnalysisFilter(t *testing.T) {
	f := NewPositiveBoundaryValueAnalysisFilter()

	// single value range
	{
		root := primitives.NewRangeInt(10, 10)
		replacements, err := f.Apply(root)
		Nil(t, err)
		Equal(t, replacements, []token.Token{
			primitives.NewConstantInt(10),
		})
	}
	// two value range
	{
		root := primitives.NewRangeInt(10, 11)
		replacements, err := f.Apply(root)
		Nil(t, err)
		Equal(t, replacements, []token.Token{
			primitives.NewConstantInt(10),
			primitives.NewConstantInt(11),
		})
	}
	// three value range
	{
		root := primitives.NewRangeInt(10, 12)
		replacements, err := f.Apply(root)
		Nil(t, err)
		Equal(t, replacements, []token.Token{
			primitives.NewConstantInt(10),
			primitives.NewConstantInt(11),
			primitives.NewConstantInt(12),
		})
	}
	// four value range
	{
		root := primitives.NewRangeInt(10, 13)
		replacements, err := f.Apply(root)
		Nil(t, err)
		Equal(t, replacements, []token.Token{
			primitives.NewConstantInt(10),
			primitives.NewConstantInt(11),
			primitives.NewConstantInt(13),
		})
	}
	// five value range
	{
		root := primitives.NewRangeInt(10, 14)
		replacements, err := f.Apply(root)
		Nil(t, err)
		Equal(t, replacements, []token.Token{
			primitives.NewConstantInt(10),
			primitives.NewConstantInt(12),
			primitives.NewConstantInt(14),
		})
	}
}
