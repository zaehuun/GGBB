package GGBB

import (
	"fmt"
	"regexp"
)

type Game interface {
}

type gameNumbers struct {
	nums string
}

func CreateGame(nums string) (Game, error) {
	//입력 값이 3~4 자리인지
	var length int = len(nums)
	if length < 3 || length > 4 {
		return nil, fmt.Errorf("invalued nums: %s", nums)
	}

	//입력 값이 숫자로만 이루어졌는지
	matched, _ := regexp.MatchString("^[0-9]*$", nums)
	if !matched {
		return nil, fmt.Errorf("invalued nums: %s", nums)
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