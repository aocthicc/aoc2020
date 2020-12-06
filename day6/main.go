package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Group struct {
	People []Person
}

func (g *Group) AddPerson(person Person) {
	g.People = append(g.People, person)
}

func (g *Group) GetTotalUniqueYesAnswers() int {
	answers := make(map[string]bool)
	for _, person := range g.People {
		for _, answer := range person.GetAnswers() {
			if _, ok := answers[answer]; !ok {
				answers[answer] = true
			}
		}
	}

	return len(answers)
}

type Person struct {
	Answers string
}

func (p *Person) GetAnswers() []string {
	return strings.Split(p.Answers, "")
}

func main() {
	data, err := ioutil.ReadFile("input")
	if nil != err {
		panic(err)
	}

	groups := mapIntoGroups(string(data))
	totalYesAnswers := 0
	for _, group := range groups {
		totalYesAnswers += group.GetTotalUniqueYesAnswers()
	}
	fmt.Println(totalYesAnswers)
}

func mapIntoGroups(data string) []Group {
	var groups []Group
	var current Group
	parsed := strings.Split(data, "\r")
	for i, line := range parsed {
		if i == len(parsed) - 1 {
			current.AddPerson(Person{Answers: strings.TrimSpace(line)})
		}
		if line == "\n" || i == len(parsed) -1 {
			groups = append(groups, current)
			current = Group{}
		} else {
			current.AddPerson(Person{Answers: strings.TrimSpace(line)})
		}
	}
	return groups
}

