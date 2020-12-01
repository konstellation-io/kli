package text

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestText(t *testing.T) {
	t.Run("Sanitize", func(t *testing.T) {
		str := " some     long     string    "
		expected := "some long string"
		require.Equal(t, Sanitize(str), expected)
	})

	t.Run("Normalilze", func(t *testing.T) {
		str := " SoME     lONg     STring    "
		expected := "some long string"
		require.Equal(t, Normalize(str), expected)
	})

	t.Run("LinesTrim", func(t *testing.T) {
		str := "String \t More   \n    Test   \n   New"
		expected := "String More\nTest\nNew"
		require.Equal(t, LinesTrim(str), expected)
	})
}
