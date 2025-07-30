package areas

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// Areas provides a builder for creating StoryArea slices with fluent API.
type Areas struct {
	areas g.Slice[gotgbot.StoryArea]
}

// AreaBuilder represents a single area being built with position/rotation that can be reused.
type AreaBuilder struct {
	areas    *Areas
	position gotgbot.StoryAreaPosition
}

// New creates a new Areas builder.
func New() *Areas {
	return &Areas{
		areas: g.NewSlice[gotgbot.StoryArea](),
	}
}

// Position creates an AreaBuilder with the specified position (x, y).
func (a *Areas) Position(x, y float64) *AreaBuilder {
	return &AreaBuilder{
		areas: a,
		position: gotgbot.StoryAreaPosition{
			XPercentage: x,
			YPercentage: y,
		},
	}
}

// Rotate sets the rotation angle for this area (in degrees).
func (ab *AreaBuilder) Rotate(angle float64) *AreaBuilder {
	ab.position.RotationAngle = angle
	return ab
}

// Rounded sets the corner radius percentage for this area (0-50%).
func (ab *AreaBuilder) Rounded(radiusPercent float64) *AreaBuilder {
	ab.position.CornerRadiusPercentage = radiusPercent
	return ab
}

// Size sets the size (width, height) for this area.
func (ab *AreaBuilder) Size(width, height float64) *AreaBuilder {
	ab.position.WidthPercentage = width
	ab.position.HeightPercentage = height
	return ab
}

// Location adds a location area type to this position.
func (ab *AreaBuilder) Location() *Areas {
	area := gotgbot.StoryArea{
		Position: ab.position,
		Type:     &gotgbot.StoryAreaTypeLocation{},
	}

	ab.areas.areas.Push(area)

	return ab.areas
}

// Reaction adds a suggested reaction area type to this position.
func (ab *AreaBuilder) Reaction(emoji g.String) *Areas {
	area := gotgbot.StoryArea{
		Position: ab.position,
		Type: &gotgbot.StoryAreaTypeSuggestedReaction{
			ReactionType: &gotgbot.ReactionTypeEmoji{
				Emoji: emoji.Std(),
			},
		},
	}

	ab.areas.areas.Push(area)

	return ab.areas
}

// Link adds a clickable link area type to this position.
func (ab *AreaBuilder) Link(url g.String) *Areas {
	area := gotgbot.StoryArea{
		Position: ab.position,
		Type: &gotgbot.StoryAreaTypeLink{
			Url: url.Std(),
		},
	}

	ab.areas.areas.Push(area)

	return ab.areas
}

// Weather adds a weather information area type to this position.
func (ab *AreaBuilder) Weather() *Areas {
	area := gotgbot.StoryArea{
		Position: ab.position,
		Type:     &gotgbot.StoryAreaTypeWeather{},
	}

	ab.areas.areas.Push(area)

	return ab.areas
}

// UniqueGift adds a unique gift area type to this position (max 1 per story).
func (ab *AreaBuilder) UniqueGift() *Areas {
	area := gotgbot.StoryArea{
		Position: ab.position,
		Type:     &gotgbot.StoryAreaTypeUniqueGift{},
	}

	ab.areas.areas.Push(area)

	return ab.areas
}

// Add inserts a raw StoryArea.
func (a *Areas) Add(area gotgbot.StoryArea) *Areas {
	a.areas.Push(area)
	return a
}

// Import appends existing StoryArea slice.
func (a *Areas) Import(src []gotgbot.StoryArea) *Areas {
	if src != nil {
		a.areas.Push(src...)
	}

	return a
}

// Clear removes all areas.
func (a *Areas) Clear() *Areas {
	a.areas = a.areas[:0]
	return a
}

// Count returns the number of areas.
func (a *Areas) Count() g.Int {
	return a.areas.Len()
}

// Std returns the underlying StoryArea slice.
func (a *Areas) Std() []gotgbot.StoryArea {
	return a.areas
}
