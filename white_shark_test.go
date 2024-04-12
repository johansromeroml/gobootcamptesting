package hunt_test

import (
	hunt "testdoubles"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		// arrange
		ws := hunt.NewWhiteShark(true, false, 200)
		tu := hunt.NewTuna("A", 100)
		// act

		err := ws.Hunt(tu)

		// assert
		assert.NoError(t, err)
		assert.True(t, ws.Tired)
		assert.False(t, ws.Hungry)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		// arrange
		ws := hunt.NewWhiteShark(false, false, 200)
		tu := hunt.NewTuna("A", 100)
		// act

		err := ws.Hunt(tu)

		// assert
		assert.Error(t, err)
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		// arrange
		ws := hunt.NewWhiteShark(true, true, 200)
		tu := hunt.NewTuna("A", 100)
		// act

		err := ws.Hunt(tu)

		// assert
		assert.Error(t, err)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		// arrange
		ws := hunt.NewWhiteShark(true, false, 80)
		tu := hunt.NewTuna("A", 100)
		// act

		err := ws.Hunt(tu)

		// assert
		assert.Error(t, err)
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		// arrange
		ws := hunt.NewWhiteShark(false, false, 200)
		// act

		err := ws.Hunt(nil)

		// assert
		//assert.Panics(t, func() { ws.Hunt(nil) })
		assert.Error(t, err)
	})
}
