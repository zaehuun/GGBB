package GGBB

import (
	"testing"
)





func TestInvalidNums(t *testing.T){

	assertError(t, "01")
	assertError(t, "01456")
	assertError(t, "abc")
	assertError(t, "14z")
	assertError(t, "as23")
	assertError(t, "1122")
	assertError(t, "2142")
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


