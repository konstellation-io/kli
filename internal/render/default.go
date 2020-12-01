package render

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

// Renderer interface that can be used to render in different formats.
type Renderer interface {
	SetHeader([]string)
	Append([]string)
	Render()
}

// DefaultRenderer creates a table renderer used as default for all commands.
func DefaultRenderer(w io.Writer) Renderer {
	table := tablewriter.NewWriter(w)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator(" ")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding(" ") // pad with spaces
	table.SetNoWhiteSpace(true)
	table.SetBorder(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	return table
}
