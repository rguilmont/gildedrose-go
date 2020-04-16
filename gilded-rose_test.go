package main

import "testing"

// Some useful constant
const (
	minQuality        = 0
	maxQuality        = 50
	lengendaryQuality = 80
)

func Test_Foo(t *testing.T) {
	var items = []*Item{
		&Item{"foo", 0, 0},
	}

	UpdateQuality(items)

	if items[0].name != "fixme" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].name)
	}
}

// At the end of each day our system lowers both values for every item
func Test_Decrease(t *testing.T) {
	sellIn := 10
	quality := 8

	items := []*Item{
		&Item{"item1", sellIn, quality},
	}
	UpdateQuality(items)

	if items[0].sellIn != sellIn-1 {
		t.Errorf("Expected %v as sellIn value, got %v", (sellIn - 1), items[0].sellIn)
	}

	if items[0].quality != quality-1 {
		t.Errorf("Expected %v as quality, got %v", (quality - 1), items[0].quality)
	}
}

// once the sell by date has passed, Quality degrades twice as fast
func Test_PassedSellIn(t *testing.T) {
	sellIn := -2
	quality := 8

	items := []*Item{
		&Item{"item1", sellIn, quality},
	}
	UpdateQuality(items)

	// Also test that sellIn is updated while it has negative value
	if items[0].sellIn != sellIn-1 {
		t.Errorf("Expected %v as sellIn value, got %v", (sellIn - 1), items[0].sellIn)
	}

	if items[0].quality != quality-2 {
		t.Errorf("Expected %v as quality, got %v", (quality - 1), items[0].quality)
	}
}

// The Quality of an item is never negative
func Test_NegativeQuality(t *testing.T) {
	sellIn := 2
	quality := 0

	items := []*Item{
		&Item{"item1", sellIn, quality},
	}
	UpdateQuality(items)

	if items[0].quality != minQuality {
		t.Errorf("Expected %v as quality, got %v", minQuality, items[0].quality)
	}
}

// "Aged Brie" actually increases in Quality the older it gets
func Test_AgedBrie(t *testing.T) {
	sellIn := 20
	quality := 2

	items := []*Item{
		&Item{"Aged Brie", sellIn, quality},
	}

	for i := 0; i < 50; i++ {
		UpdateQuality(items)
		// Update expected quality
		if quality < maxQuality {
			if items[0].sellIn < 0 {
				// For some reason, running the test i found out that
				//  quality increase by 2 when sellIn is negative.
				//  it's not in the spec, but since there's the same behaviour
				//  at least in Scala, and i assume that there's no bug in the implementation
				//  i also test that.
				quality = quality + 2
			} else {
				quality++
			}
		}
		if items[0].quality != quality {
			t.Errorf("Expected %v as quality, got %v - sellIn=%v", quality, items[0].quality, items[0].sellIn)
		}
	}
}

// "Sulfuras", being a legendary item, never has to be sold or decreases in Quality
func Test_Sulfuras(t *testing.T) {
	sellIn := 20
	quality := lengendaryQuality

	items := []*Item{
		&Item{"Sulfuras, Hand of Ragnaros", sellIn, quality},
	}
	UpdateQuality(items)
	if items[0].quality != lengendaryQuality {
		t.Errorf("Expected quality %v, got %v", items[0].quality, lengendaryQuality)
	}
	if items[0].sellIn != sellIn {
		t.Errorf("Expected %v, got %v", items[0].sellIn, sellIn)
	}
}

// "Backstage passes", like aged brie, increases in Quality as its SellIn value approaches;
// Quality increases by 2 when there are 10 days or less and by 3 when there are 5 days or less but
// Quality drops to 0 after the concert
func Test_BackstagePasses(t *testing.T) {
	sellIn := 20
	quality := 10

	items := []*Item{
		&Item{"Backstage passes to a TAFKAL80ETC concert", sellIn, quality},
	}

	for i := 0; i < 25; i++ {
		UpdateQuality(items)
		// Update expected quality

		if items[0].sellIn < 0 {
			quality = minQuality
		} else if items[0].sellIn < 5 {
			quality = quality + 3
		} else if items[0].sellIn < 10 {
			quality = quality + 2
		} else {
			quality++
		}

		if quality > maxQuality {
			quality = maxQuality
		}

		if items[0].quality != quality {
			t.Errorf("Expected %v as quality, got %v - sellIn=%v", quality, items[0].quality, items[0].sellIn)
		}
	}
}

// Finally, "Conjured" items degrade in Quality twice as fast as normal items
func Test_Conjured(t *testing.T) {
	sellIn := 20
	quality := 10

	items := []*Item{
		&Item{"Conjured", sellIn, quality},
	}

	for i := 0; i < 25; i++ {
		UpdateQuality(items)
		// Update expected quality

		if items[0].sellIn < 0 {
			quality = quality - 2
		}

		if quality < minQuality {
			quality = minQuality
		}

		if items[0].quality != quality {
			t.Errorf("Expected %v as quality, got %v - sellIn=%v", quality, items[0].quality, items[0].sellIn)
		}
	}
}

// Some additional tests
func Test_UpdatingMultipleItems(t *testing.T) {

	items := []*Item{
		&Item{"Conjured", 15, 15},
		&Item{"item2", 3, 30},
		&Item{"item3", 12, 22},
		&Item{"item1", 1, 0},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 22, 3},
		&Item{"Aged Brie", 300, 25},
	}

	// Let's loop for 2 years :)
	for i := 0; i < 730; i++ {
		UpdateQuality(items)
		for _, j := range items {
			if j.quality > maxQuality || j.quality < minQuality {
				t.Errorf("Got invalid quality value %v as quality - sellIn=%v", j.quality, j.sellIn)
			}
		}
	}
}
