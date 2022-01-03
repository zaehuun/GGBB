package GGBB

import (
	"testing"
)

func TestGuessInvalidNums(t *testing.T) {
	game, _ := CreateGame("123")
	assertGuessReturnError(t, game, "12")
	assertGuessReturnError(t, game, "1234")
	assertGuessReturnError(t, game, "12a")
	assertGuessReturnError(t, game, "Z34")
	assertGuessReturnError(t, game, "112")
	
	_, err := game.guess("876")
	if err != nil{
		t.Fatalf("oops")
	}
}

func assertGuessReturnError(t *testing.T, game Game, nums string){
	_, err := game.guess(nums)
	if err == nil {
		t.Fatalf("error must be returned: %s", nums)
	}
}