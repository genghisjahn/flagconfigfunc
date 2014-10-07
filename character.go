package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

type Character struct {
	Class        string
	Strength     int
	Intelligence int
	Wisdom       int
	Dexterity    int
	Constitution int
	Charisma     int
}

func NewCharacter(cfg Config) Character {
	result := Character{}
	result.Class = cfg.Class.Name

	dieRoller := RollMethods[cfg.Class.RollMethod]

	scores := make([]int, 6, 6)
	for k, _ := range scores {
		scores[k], _ = dieRoller()
	}

	sortedScores := scores
	sort.Sort(intArray(sortedScores))
	maxScore := sortedScores[len(sortedScores)-1]

	remainingScores := sortedScores[:5]

	for i := range remainingScores {
		j := rand.Intn(i + 1)
		remainingScores[i], remainingScores[j] = remainingScores[j], remainingScores[i]
	}

	switch cfg.Class.MaxAbility {
	case "strength":
		result.Strength = maxScore
		result.Intelligence = remainingScores[0]
		result.Wisdom = remainingScores[1]
	case "intelligence":
		result.Strength = remainingScores[0]
		result.Intelligence = maxScore
		result.Wisdom = remainingScores[1]
	case "wisdom":
		result.Strength = remainingScores[0]
		result.Intelligence = remainingScores[1]
		result.Wisdom = maxScore
	}
	result.Dexterity = remainingScores[2]
	result.Constitution = remainingScores[3]
	result.Charisma = remainingScores[4]

	return result
}

func (c Character) String() string {
	charTemplate := "Class: %v\nStrength: %v\nIntelligence: %v\nWisdom: %v\nDexterity: %v\nConstitution: %v\nCharisma: %v\n"
	return fmt.Sprintf(charTemplate, strings.Title(c.Class), c.Strength, c.Intelligence, c.Wisdom, c.Dexterity, c.Constitution, c.Charisma)
}

func (c Character) Error() string {
	return "This character has errors, or flaws if you will."
}
