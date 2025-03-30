package todo

import (
	"encoding/json"
	"os"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	Position int
	Done     bool
}

func (item *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		item.Priority = 1
	case 3:
		item.Priority = 3
	default:
		item.Priority = 2
	}
}

func (item *Item) GetPriority() string {
	if item.Priority == 1 {
		return "(H)"
	}
	if item.Priority == 2 {
		return "(M)"
	}
	if item.Priority == 3 {
		return "(L)"
	}

	return ""
}

func (item *Item) GetPosition() string {
	return strconv.Itoa(item.Position) + "."
}

func (item *Item) DisplayDone() string {
	if item.Done {
		return "✓"
	}

	return ""
}

// ByPriority Implements a sort.Sort() interface  for array of Item base of priority and position
type ByPriority []Item

func (a ByPriority) Len() int {
	return len(a)
}

func (a ByPriority) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByPriority) Less(i, j int) bool {

	if a[i].Priority == a[j].Priority {
		return a[i].Position < a[j].Position
	}

	if a[i].Done != a[j].Done {
		return a[i].Done
	}

	return a[i].Position < a[j].Position
}

func SaveItems(filename string, items []Item) error {

	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {

	b, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}

	var items []Item
	err = json.Unmarshal(b, &items)
	if err != nil {
		return []Item{}, err
	}

	for i := range items {
		items[i].Position = i + 1
	}

	return items, nil
}
