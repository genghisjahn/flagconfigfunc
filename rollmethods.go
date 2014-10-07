package main

import (
	"math/rand"
	"sort"
	"time"
)

type fnDieRollMethod func() (int, error)

var (
	RollMethods = map[string]fnDieRollMethod{
		"r3d6":  R3D6,
		"r1d20": R1D20,
		"r4d6":  R4D6,
	}
)

func R3D6() (int, error) {
	rand.Seed(time.Now().UnixNano())
	result := 0
	for i := 0; i < 3; i++ {
		result += rand.Intn(6) + 1
	}
	return result, nil
}

func R1D20() (int, error) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(20) + 1, nil
}

func R4D6() (int, error) {
	rand.Seed(time.Now().UnixNano())
	result := 0
	rolls := make([]int, 4, 4)
	for k, _ := range rolls {
		rolls[k] = rand.Intn(6) + 1
	}

	sort.Sort(intArray(rolls))
	for _, v := range rolls[1:4] {
		result += v
	}
	return result, nil
}

type intArray []int

func (s intArray) Len() int           { return len(s) }
func (s intArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s intArray) Less(i, j int) bool { return s[i] < s[j] }
