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

func UpdateQuality(items []*Item) {
	for _, item := range items {
		switch item.name {
		case "Aged Brie":
			item.sellIn--
			if item.sellIn < 0 {
				item.quality = validQuality(item.quality + 2)
			} else {
				item.quality = validQuality(item.quality + 1)
			}
		case "Backstage passes to a TAFKAL80ETC concert":
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
		case "Sulfuras, Hand of Ragnaros":
			// Do nothing
		default:
			item.sellIn--
			if item.sellIn < 0 {
				item.quality = validQuality(item.quality - 2)
			} else {
				item.quality = validQuality(item.quality - 1)
			}
		}
	}
}
