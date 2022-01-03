package GGBB

import (
	"fmt"
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
	return &gameNumbers{nums: nums}, nil
}

func TestInvalidNums(t *testing.T){
	_, err := CreateGame("01")
	if err == nil {
		t.Fatalf("error must be returned: invalie nums: %s", "01")
	}

	_, err2 := CreateGame("01456")
	if err2 == nil {
		t.Fatalf("error must be returned: invalie nums: %s", "01456")
	}
}

func TestCreateGame(t *testing.T){
	game, err := CreateGame("123")
	if err != nil || game == nil{
		t.Fatalf("game be returned")
	}
}


