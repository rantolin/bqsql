/*
Copyright © 2024 Roberto Antolin <rantolin DOT geo AT gmail DOT com>
*/
package formats

import (
	"fmt"
	"io"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func CalculateRowWidths(it *bigquery.RowIterator, schema bigquery.Schema) ([]int, error) {
	widths := make([]int, len(it.Schema))
	for {
		var row []bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		for i, val := range row {
			width := max(len(fmt.Sprint(val)), len(schema[i].Name))
			if width > widths[i] {
				widths[i] = width
			}
		}
	}

	return widths, nil
}

func PrintFormatedRow(w io.Writer, row []bigquery.Value, widths []int) {
	for i, val := range row {
		if i != 0 {
			fmt.Fprint(w, " | ")
		}
		fmt.Fprintf(w, "%-*v", widths[i], val)
	}
	fmt.Fprintln(w)
}
