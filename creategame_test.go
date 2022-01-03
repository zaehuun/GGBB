package GGBB

import "testing"



func TestCreateGame(t *testing.T){
	var game Game = CreateGame("012")
	if game == nil{
		t.Fatalf("no game created")
	}
}
