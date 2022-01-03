package GGBB

import "testing"

type Game interface{

}

type gameNumbers struct {
	nums string
}

func CreateGame(nums string) Game {
	return &gameNumbers{nums: nums}
}

func TestCreateGame(t *testing.T){
	var game Game = CreateGame("012")
	if game == nil{
		t.Fatalf("no game created")
	}
}
