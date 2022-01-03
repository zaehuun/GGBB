package GGBB

import (
	"fmt"
	"regexp"
	"testing"
)

type Game interface{

}

type gameNumbers struct {
	nums string
}

func CreateGame(nums string) (Game, error) {
	var length  int = len(nums)
	if length < 3 || length > 4 {
		return nil, fmt.Errorf("invalued nums: %s", nums)
	}

	matched, _ := regexp.MatchString("^[0-9]*$", nums)
	if !matched{
		return nil, fmt.Errorf("invalued nums: %s", nums)
	}
	return &gameNumbers{nums: nums}, nil
}

func TestInvalidNums(t *testing.T){

	assertError(t, "01")
	assertError(t, "01456")
	assertError(t, "abc")
}

func TestCreateGame(t *testing.T){
	
	game, err := CreateGame("123")
	if err != nil || game == nil{
		t.Fatalf("game be returned")
	}
}

func assertError(t *testing.T, nums string) {
	_, err := CreateGame(nums)
	if err == nil {
		t.Fatalf("error must be returned: invalie nums: %s", nums)
	}
}


