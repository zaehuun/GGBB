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

func TestGuess(t *testing.T) {
	g, _ := CreateGame("123")
	assertScore(t, g, "567", 0, 0)
	assertScore(t, g, "156", 1, 0)
	assertScore(t, g, "025", 1, 0)
	assertScore(t, g, "453", 1, 0)
	assertScore(t, g, "129", 2, 0)
	assertScore(t, g, "103", 2, 0)
	assertScore(t, g, "023", 2, 0)
	assertScore(t, g, "123", 3, 0)

	assertScore(t, g, "910", 0, 1)
	assertScore(t, g, "901", 0, 1)
	assertScore(t, g, "294", 0, 1)
	assertScore(t, g, "892", 0, 1)
}

func assertScore(t *testing.T, g Game, nums string, expectedStrikes int, expectedBalls int){
	score, _ := g.guess(nums)
	if score.strikes != expectedStrikes {
		t.Fatalf("guess number %s: strikes exprected %d, but %d", nums, expectedStrikes, score.strikes)
	}

	if score.balls != expectedBalls {
		t.Fatalf("guess number %s: balls exprected %d, but %d", nums, expectedBalls, score.balls)
	}
}