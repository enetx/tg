package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/richtext"
)

// BlockTableCell represents a rich message table cell builder.
type BlockTableCell struct {
	cell *gotgbot.RichBlockTableCell
}

// Cell creates a new BlockTableCell builder with the given text.
// A cell without text is invisible.
func Cell(text *richtext.Text) *BlockTableCell {
	return &BlockTableCell{
		cell: &gotgbot.RichBlockTableCell{
			Text:   text.Std(),
			Align:  AlignLeft.String(),
			Valign: ValignTop.String(),
		},
	}
}

// Text sets the cell text.
func (btc *BlockTableCell) Text(text *richtext.Text) *BlockTableCell {
	btc.cell.Text = text.Std()
	return btc
}

// Header marks the cell as a header cell.
func (btc *BlockTableCell) Header() *BlockTableCell {
	btc.cell.IsHeader = true
	return btc
}

// Colspan sets the number of columns the cell spans.
func (btc *BlockTableCell) Colspan(colspan int64) *BlockTableCell {
	btc.cell.Colspan = colspan
	return btc
}

// Rowspan sets the number of rows the cell spans.
func (btc *BlockTableCell) Rowspan(rowspan int64) *BlockTableCell {
	btc.cell.Rowspan = rowspan
	return btc
}

// Align sets the horizontal alignment of the cell content.
func (btc *BlockTableCell) Align(align CellAlign) *BlockTableCell {
	btc.cell.Align = align.String()
	return btc
}

// Valign sets the vertical alignment of the cell content.
func (btc *BlockTableCell) Valign(valign CellValign) *BlockTableCell {
	btc.cell.Valign = valign.String()
	return btc
}

// Std returns the underlying gotgbot.RichBlockTableCell.
func (btc *BlockTableCell) Std() gotgbot.RichBlockTableCell {
	return *btc.cell
}
