package GGBB

import (
	"fmt"
	"testing"
)

func TestGuessInvalidNums(t *testing.T) {
	game, _ := CreateGame("123")
	_, err := game.guess("12")
	if err == nil {
		fmt.Errorf("error must be returned: %s", "12")
	}
}