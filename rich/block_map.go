package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
)

// BlockMap represents a rich message map block builder.
type BlockMap struct {
	block *gotgbot.InputRichBlockMap
}

// Map creates a new BlockMap builder with the required location.
func Map(latitude, longitude float64) *BlockMap {
	return &BlockMap{
		block: &gotgbot.InputRichBlockMap{
			Location: gotgbot.Location{Latitude: latitude, Longitude: longitude},
		},
	}
}

// Accuracy sets the radius of uncertainty for the location, in meters.
func (bm *BlockMap) Accuracy(accuracy float64) *BlockMap {
	bm.block.Location.HorizontalAccuracy = accuracy
	return bm
}

// Zoom sets the map zoom level.
func (bm *BlockMap) Zoom(zoom int64) *BlockMap {
	bm.block.Zoom = zoom
	return bm
}

// Size sets the map width and height.
func (bm *BlockMap) Size(width, height int64) *BlockMap {
	bm.block.Width = width
	bm.block.Height = height

	return bm
}

// Caption sets the block caption.
func (bm *BlockMap) Caption(caption *Caption) *BlockMap {
	bm.block.Caption = caption.Std()
	return bm
}

// Build creates the gotgbot.InputRichBlock.
func (bm *BlockMap) Build() gotgbot.InputRichBlock {
	return *bm.block
}
