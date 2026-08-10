package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/richtext"
)

// BlockTable represents a rich message table block builder.
type BlockTable struct {
	block *gotgbot.InputRichBlockTable
	rows  g.Slice[g.Slice[*BlockTableCell]]
}

// Table creates a new BlockTable builder.
func Table() *BlockTable {
	return &BlockTable{
		block: new(gotgbot.InputRichBlockTable),
		rows:  g.NewSlice[g.Slice[*BlockTableCell]](),
	}
}

// Row appends a row of cells to the table.
func (bt *BlockTable) Row(cells ...*BlockTableCell) *BlockTable {
	row := g.Slice[*BlockTableCell](cells).
		Iter().
		Exclude(func(cell *BlockTableCell) bool { return cell == nil }).
		Collect().
		Slice()

	bt.rows.Push(row)

	return bt
}

// Bordered draws borders around the table cells.
func (bt *BlockTable) Bordered() *BlockTable {
	bt.block.IsBordered = true
	return bt
}

// Striped highlights alternating table rows.
func (bt *BlockTable) Striped() *BlockTable {
	bt.block.IsStriped = true
	return bt
}

// Caption sets the table caption.
func (bt *BlockTable) Caption(caption *richtext.Text) *BlockTable {
	bt.block.Caption = caption.Std()
	return bt
}

// Build creates the gotgbot.InputRichBlock.
func (bt *BlockTable) Build() gotgbot.InputRichBlock {
	block := *bt.block

	if !bt.rows.IsEmpty() {
		block.Cells = g.TransformSlice(bt.rows, func(row g.Slice[*BlockTableCell]) []gotgbot.RichBlockTableCell {
			return g.TransformSlice(row, (*BlockTableCell).Std)
		})
	}

	return block
}
