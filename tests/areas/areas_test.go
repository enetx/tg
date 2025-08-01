package areas_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	. "github.com/enetx/tg/areas"
)

func TestNew(t *testing.T) {
	areas := New()

	if areas == nil {
		t.Error("Expected New() to create an Areas instance")
	}

	if areas.Count() != 0 {
		t.Errorf("Expected empty areas, got count: %d", areas.Count())
	}

	if len(areas.Std()) != 0 {
		t.Error("Expected empty StoryArea slice")
	}
}

func TestAreas_Position_Location(t *testing.T) {
	areas := New()
	x := 25.5
	y := 75.0

	// Test creating a location area with position
	result := areas.Position(x, y).Location()

	if result != areas {
		t.Error("Expected Location() to return parent Areas")
	}

	if areas.Count() != 1 {
		t.Error("Expected one area to be added")
	}

	storyAreas := areas.Std()
	if len(storyAreas) != 1 {
		t.Error("Expected one story area")
		return
	}

	area := storyAreas[0]
	if area.Position.XPercentage != x {
		t.Errorf("Expected X percentage %f, got %f", x, area.Position.XPercentage)
	}

	if area.Position.YPercentage != y {
		t.Errorf("Expected Y percentage %f, got %f", y, area.Position.YPercentage)
	}

	// Verify it's a location area
	if area.Type == nil {
		t.Error("Expected area type to be set")
	}

	if _, ok := area.Type.(*gotgbot.StoryAreaTypeLocation); !ok {
		t.Error("Expected area type to be StoryAreaTypeLocation")
	}
}

func TestAreaBuilder_Rotate(t *testing.T) {
	areas := New()
	angle := 45.0

	// Test rotation by creating an area and checking the result
	areas.Position(10, 20).Rotate(angle).Location()

	storyAreas := areas.Std()
	if len(storyAreas) != 1 {
		t.Error("Expected one story area")
		return
	}

	area := storyAreas[0]
	if area.Position.RotationAngle != angle {
		t.Errorf("Expected rotation angle %f, got %f", angle, area.Position.RotationAngle)
	}
}

func TestAreaBuilder_Rounded(t *testing.T) {
	areas := New()
	radius := 25.0

	// Test corner radius by creating an area and checking the result
	areas.Position(10, 20).Rounded(radius).Location()

	storyAreas := areas.Std()
	if len(storyAreas) != 1 {
		t.Error("Expected one story area")
		return
	}

	area := storyAreas[0]
	if area.Position.CornerRadiusPercentage != radius {
		t.Errorf("Expected corner radius %f, got %f", radius, area.Position.CornerRadiusPercentage)
	}
}

func TestAreaBuilder_Size(t *testing.T) {
	areas := New()
	width := 50.0
	height := 30.0

	// Test size by creating an area and checking the result
	areas.Position(10, 20).Size(width, height).Location()

	storyAreas := areas.Std()
	if len(storyAreas) != 1 {
		t.Error("Expected one story area")
		return
	}

	area := storyAreas[0]
	if area.Position.WidthPercentage != width {
		t.Errorf("Expected width percentage %f, got %f", width, area.Position.WidthPercentage)
	}

	if area.Position.HeightPercentage != height {
		t.Errorf("Expected height percentage %f, got %f", height, area.Position.HeightPercentage)
	}
}

func TestAreaBuilder_Reaction(t *testing.T) {
	areas := New()
	emoji := g.String("😀")

	result := areas.Position(10, 20).Reaction(emoji)

	if result != areas {
		t.Error("Expected Reaction() to return parent Areas")
	}

	if areas.Count() != 1 {
		t.Error("Expected one area to be added")
	}

	storyAreas := areas.Std()
	area := storyAreas[0]

	// Verify it's a reaction area
	reactionType, ok := area.Type.(*gotgbot.StoryAreaTypeSuggestedReaction)
	if !ok {
		t.Error("Expected area type to be StoryAreaTypeSuggestedReaction")
		return
	}

	emojiReaction, ok := reactionType.ReactionType.(*gotgbot.ReactionTypeEmoji)
	if !ok {
		t.Error("Expected reaction type to be ReactionTypeEmoji")
		return
	}

	if emojiReaction.Emoji != emoji.Std() {
		t.Errorf("Expected emoji %s, got %s", emoji.Std(), emojiReaction.Emoji)
	}
}

func TestAreaBuilder_Link(t *testing.T) {
	areas := New()
	url := g.String("https://example.com")

	result := areas.Position(10, 20).Link(url)

	if result != areas {
		t.Error("Expected Link() to return parent Areas")
	}

	if areas.Count() != 1 {
		t.Error("Expected one area to be added")
	}

	storyAreas := areas.Std()
	area := storyAreas[0]

	// Verify it's a link area
	linkType, ok := area.Type.(*gotgbot.StoryAreaTypeLink)
	if !ok {
		t.Error("Expected area type to be StoryAreaTypeLink")
		return
	}

	if linkType.Url != url.Std() {
		t.Errorf("Expected URL %s, got %s", url.Std(), linkType.Url)
	}
}

func TestAreaBuilder_Weather(t *testing.T) {
	areas := New()

	result := areas.Position(10, 20).Weather()

	if result != areas {
		t.Error("Expected Weather() to return parent Areas")
	}

	if areas.Count() != 1 {
		t.Error("Expected one area to be added")
	}

	storyAreas := areas.Std()
	area := storyAreas[0]

	// Verify it's a weather area
	if _, ok := area.Type.(*gotgbot.StoryAreaTypeWeather); !ok {
		t.Error("Expected area type to be StoryAreaTypeWeather")
	}
}

func TestAreaBuilder_UniqueGift(t *testing.T) {
	areas := New()

	result := areas.Position(10, 20).UniqueGift()

	if result != areas {
		t.Error("Expected UniqueGift() to return parent Areas")
	}

	if areas.Count() != 1 {
		t.Error("Expected one area to be added")
	}

	storyAreas := areas.Std()
	area := storyAreas[0]

	// Verify it's a unique gift area
	if _, ok := area.Type.(*gotgbot.StoryAreaTypeUniqueGift); !ok {
		t.Error("Expected area type to be StoryAreaTypeUniqueGift")
	}
}

func TestAreaBuilder_ChainedOperations(t *testing.T) {
	areas := New()

	// Test chaining multiple operations
	result := areas.Position(10, 20).
		Rotate(45).
		Rounded(25).
		Size(50, 30).
		Location()

	if result != areas {
		t.Error("Expected chained operations to return parent Areas")
	}

	storyAreas := areas.Std()
	if len(storyAreas) != 1 {
		t.Error("Expected one story area")
		return
	}

	area := storyAreas[0]
	if area.Position.XPercentage != 10 {
		t.Error("Expected X percentage to be preserved")
	}
	if area.Position.YPercentage != 20 {
		t.Error("Expected Y percentage to be preserved")
	}
	if area.Position.RotationAngle != 45 {
		t.Error("Expected rotation angle to be preserved")
	}
	if area.Position.CornerRadiusPercentage != 25 {
		t.Error("Expected corner radius to be preserved")
	}
	if area.Position.WidthPercentage != 50 {
		t.Error("Expected width percentage to be preserved")
	}
	if area.Position.HeightPercentage != 30 {
		t.Error("Expected height percentage to be preserved")
	}
}

func TestAreas_Add(t *testing.T) {
	areas := New()

	customArea := gotgbot.StoryArea{
		Position: gotgbot.StoryAreaPosition{
			XPercentage: 50,
			YPercentage: 60,
		},
		Type: &gotgbot.StoryAreaTypeLocation{},
	}

	result := areas.Add(customArea)

	if result != areas {
		t.Error("Expected Add() to return Areas")
	}

	if areas.Count() != 1 {
		t.Error("Expected one area to be added")
	}

	storyAreas := areas.Std()
	if len(storyAreas) != 1 {
		t.Error("Expected one story area in slice")
	}
}

func TestAreas_Import(t *testing.T) {
	areas := New()

	importAreas := []gotgbot.StoryArea{
		{
			Position: gotgbot.StoryAreaPosition{XPercentage: 10, YPercentage: 20},
			Type:     &gotgbot.StoryAreaTypeLocation{},
		},
		{
			Position: gotgbot.StoryAreaPosition{XPercentage: 30, YPercentage: 40},
			Type:     &gotgbot.StoryAreaTypeWeather{},
		},
	}

	result := areas.Import(importAreas)

	if result != areas {
		t.Error("Expected Import() to return Areas")
	}

	if areas.Count() != 2 {
		t.Errorf("Expected 2 areas, got %d", areas.Count())
	}

	// Test importing nil slice
	result = areas.Import(nil)
	if result != areas {
		t.Error("Expected Import(nil) to return Areas")
	}

	if areas.Count() != 2 {
		t.Error("Expected count to remain unchanged after importing nil")
	}
}

func TestAreas_Clear(t *testing.T) {
	areas := New()

	// Add some areas first
	areas.Position(10, 20).Location()
	areas.Position(30, 40).Weather()

	if areas.Count() != 2 {
		t.Error("Expected 2 areas before clear")
	}

	result := areas.Clear()

	if result != areas {
		t.Error("Expected Clear() to return Areas")
	}

	if areas.Count() != 0 {
		t.Error("Expected 0 areas after clear")
	}

	if len(areas.Std()) != 0 {
		t.Error("Expected empty slice after clear")
	}
}

func TestAreas_MultipleAreas(t *testing.T) {
	areas := New()

	// Add multiple different area types
	areas.Position(10, 20).Location()
	areas.Position(30, 40).Reaction(g.String("👍"))
	areas.Position(50, 60).Link(g.String("https://example.com"))
	areas.Position(70, 80).Weather()

	if areas.Count() != 4 {
		t.Errorf("Expected 4 areas, got %d", areas.Count())
	}

	storyAreas := areas.Std()
	if len(storyAreas) != 4 {
		t.Error("Expected 4 story areas in slice")
	}

	// Verify each area type
	for i, area := range storyAreas {
		switch i {
		case 0:
			if _, ok := area.Type.(*gotgbot.StoryAreaTypeLocation); !ok {
				t.Errorf("Expected area %d to be Location type", i)
			}
		case 1:
			if _, ok := area.Type.(*gotgbot.StoryAreaTypeSuggestedReaction); !ok {
				t.Errorf("Expected area %d to be SuggestedReaction type", i)
			}
		case 2:
			if _, ok := area.Type.(*gotgbot.StoryAreaTypeLink); !ok {
				t.Errorf("Expected area %d to be Link type", i)
			}
		case 3:
			if _, ok := area.Type.(*gotgbot.StoryAreaTypeWeather); !ok {
				t.Errorf("Expected area %d to be Weather type", i)
			}
		}
	}
}
