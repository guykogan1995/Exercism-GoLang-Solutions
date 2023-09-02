package kindergarten

import (
	"errors"
	"sort"
)

type Garden struct {
	diagram  string
	children []string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	var g Garden
	g.diagram = diagram
	g.children = children
	if diagram == "" || len(children) == 0 || diagram[0] != '\n' || len(diagram)%2 != 0 {
		return &g, errors.New("diagram cannot be empty and children cannot be none")
	}
	return &g, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	newSortedChildren := g.children
	sort.Strings(newSortedChildren)
	isChild := false
	count := 0
	for _, s := range newSortedChildren {
		count += 2
		if s == child {
			isChild = true
			break
		}
	}
	if count == 0 || isChild == false {
		return []string{}, false
	}
	plants := changeToPlant(g.diagram[count-1 : count+1])
	plants = append(plants, changeToPlant(g.diagram[len(g.diagram)/2+count-1:len(g.diagram)/2+count+1])...)
	return plants, true
}

func changeToPlant(convert string) []string {
	var plants []string
	for _, l := range convert {
		switch l {
		case 'R':
			plants = append(plants, "radishes")
		case 'V':
			plants = append(plants, "violets")
		case 'C':
			plants = append(plants, "clover")
		case 'G':
			plants = append(plants, "grass")
		}
	}
	return plants
}
