package GGBB

import (
	"fmt"
	"regexp"
)

type Game interface {
	guess(s string) (Score, error)
}

type Score struct {
	
}

type gameNumbers struct {
	nums string
}

func (g gameNumbers) guess(nums string) (Score, error) {
	//정답과 추측 값의 길이가 같은지
	if len(g.nums) != len(nums){
		return Score{}, fmt.Errorf("length mismatch")
	}

	// 추측 값이 숫자로만 이루어졌는지
	matched, _ := regexp.MatchString("^[0-9]*$", nums)
	if !matched {
		return Score{}, fmt.Errorf("invalid nums: %s", nums)
	}

	//중복된 숫자가 있는지
	used := make(map[uint8]bool)
	for i := 0; i < len(nums); i++ {
		_, found := used[nums[i]]
		fmt.Print(nums[i])
		if found {
			return Score{}, fmt.Errorf("invalued nums: %s", nums)
		}
		used[nums[i]] = true
	}
	return Score{}, nil
}

func CreateGame(nums string) (Game, error) {
	//입력 값이 3~4 자리인지
	var length int = len(nums)
	if length < 3 || length > 4 {
		return nil, fmt.Errorf("invalid nums: %s", nums)
	}

	//입력 값이 숫자로만 이루어졌는지
	matched, _ := regexp.MatchString("^[0-9]*$", nums)
	if !matched {
		return nil, fmt.Errorf("invalid nums: %s", nums)
	}

	//중복된 숫자가 있는지
	used := make(map[uint8]bool)
	for i := 0; i < length; i++ {
		_, found := used[nums[i]]
		fmt.Print(nums[i])
		if found {
			return nil, fmt.Errorf("invalued nums: %s", nums)
		}
		used[nums[i]] = true
	}
	return &gameNumbers{nums: nums}, nil
}