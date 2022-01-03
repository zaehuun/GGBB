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
	// return &gameNumbers{nums: nums}, nil
	return nil, fmt.Errorf("invalued nums: %s", nums)
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


