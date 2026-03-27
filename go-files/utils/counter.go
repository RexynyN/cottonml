package utils

import "sort"

type Counter struct {
	values map[any]int
}

type CounterItem struct {
	Item  any
	Count int
}

func NewCounter() *Counter {
	return &Counter{}
}

func CounterFromSlice(items []any) *Counter {
	c := &Counter{}
	c.AddSlice(items)

	return c
}

func (c *Counter) Add(item any) {
	c.values[item]++
}

func (c *Counter) DeleteItem(item any) {
	delete(c.values, item)
}

func (c *Counter) Subtract(item any) {
	c.values[item]--
}

// Add all items from a slice of items
func (c *Counter) AddSlice(items []any) {
	for _, item := range items {
		c.Add(item)
	}
}

// Returns the most common Item from the Counter as a CounterItem
func (c *Counter) MostCommonOne() CounterItem {
	var maxItem any
	maxCount := 0
	for item, count := range c.values {
		if count >= maxCount {
			maxItem = item
			maxCount = count
		}
	}

	return CounterItem{Item: maxItem, Count: maxCount}
}

// Returns all the items that have the max value inside the Counter
func (c *Counter) MostCommonItems() []CounterItem {
	maxCount := 0
	for _, count := range c.values {
		if count >= maxCount {
			maxCount = count
		}
	}

	winners := []CounterItem{}
	for item, count := range c.values {
		if maxCount == count {
			winners = append(winners, CounterItem{Item: item, Count: count})
		}
	}

	return winners
}

// Returns the N most common items from Counter as a slice of CounterItem
func (c *Counter) TopNCommon(n uint32) []CounterItem {
	sortedItems := c.GetCounterItems()

	return sortedItems[:n]
}

// Return the raw Counter items as map of counts per items
func (c *Counter) GetValues() map[any]int {
	return c.values
}

// Return the Counter as a sorted slice of CounterItem
func (c *Counter) GetCounterItems() (items []CounterItem) {
	for item, count := range c.values {
		items = append(items, CounterItem{Item: item, Count: count})
	}

	sort.Slice(items, func(x, y int) bool {
		return items[x].Count > items[y].Count
	})

	return
}
