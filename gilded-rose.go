package main

const (
	// Minimum quality value
	minQuality = 0
	// Maximum quality value
	maxQuality = 50
	// Legendary quality
	lengendaryQuality = 80
)

type Item struct {
	name            string
	sellIn, quality int
}

// Simple function that return a valid quality
func validQuality(quality int) int {
	if quality > maxQuality {
		return maxQuality
	} else if quality < minQuality {
		return minQuality
	} else {
		return quality
	}
}

// augmentedItem helps to extend the default item
//  to have quality update implementation
//  in each struct.
//  Should make it easier to maintain.
type augmentedItem interface {
	updateQuality()
}

// name of available augmented types
const (
	agedBrieType        = "Aged Brie"
	backstagePassesType = "Backstage Passes"
	sulfurasType        = "Sulfuras"
	conjuredType        = "Conjured"
)

// Map of item -> augmentedType
// I hope this will live in some kind of database :)
var itemMapping map[string]string = map[string]string{
	"Aged Brie":                                 agedBrieType,
	"Backstage passes to a TAFKAL80ETC concert": backstagePassesType,
	"Sulfuras, Hand of Ragnaros":                sulfurasType,
	"Conjured":                                  conjuredType,
}

// Now the items structs declaration
type agedBrie struct {
	*Item
}

func (item *agedBrie) updateQuality() {
	item.sellIn--
	if item.sellIn < 0 {
		item.quality = validQuality(item.quality + 2)
	} else {
		item.quality = validQuality(item.quality + 1)
	}
}

type backstagePasses struct {
	*Item
}

func (item *backstagePasses) updateQuality() {
	item.sellIn--
	if item.sellIn < 0 {
		item.quality = validQuality(minQuality)
	} else if item.sellIn < 5 {
		item.quality = validQuality(item.quality + 3)
	} else if item.sellIn < 10 {
		item.quality = validQuality(item.quality + 2)
	} else {
		item.quality = validQuality(item.quality + 1)
	}
}

type sulfuras struct {
	*Item
}

func (item *sulfuras) updateQuality() {
	// Do nothing
}

type conjured struct {
	*Item
}

func (item *conjured) updateQuality() {
	item.sellIn--
	if item.sellIn < 0 {
		item.quality = validQuality(item.quality - 4)
	} else {
		item.quality = validQuality(item.quality - 2)
	}
}

type defaultItem struct {
	*Item
}

func (item *defaultItem) updateQuality() {
	item.sellIn--
	if item.sellIn < 0 {
		item.quality = validQuality(item.quality - 2)
	} else {
		item.quality = validQuality(item.quality - 1)
	}
}

// return augmented Item given an item
func augmentItem(item *Item) augmentedItem {
	t, ok := itemMapping[item.name]
	if ok {
		switch t {
		case agedBrieType:
			return &agedBrie{item}
		case backstagePassesType:
			return &backstagePasses{item}
		case sulfurasType:
			return &sulfuras{item}
		case conjuredType:
			return &conjured{item}
		}
	}
	return &defaultItem{item}
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		augmentItem(item).updateQuality()
	}
}
