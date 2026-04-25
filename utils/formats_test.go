package formats

import (
	"bytes"
	"testing"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

// MockRowProvider satisfies the RowProvider interface
type MockRowProvider struct {
	schema bigquery.Schema
	rows   [][]bigquery.Value
	index  int
}

func (m *MockRowProvider) Next(dst interface{}) error {
	if m.index >= len(m.rows) {
		return iterator.Done
	}
	d := dst.(*[]bigquery.Value)
	*d = m.rows[m.index]
	m.index++
	return nil
}

func (m *MockRowProvider) Schema() bigquery.Schema {
	return m.schema
}

func testRowProviderContract(t *testing.T, p RowProvider, expectedSchema bigquery.Schema, expectedRows [][]bigquery.Value) {
	t.Helper()

	// Check Schema
	if len(p.Schema()) != len(expectedSchema) {
		t.Errorf("expected schema length %d, got %d", len(expectedSchema), len(p.Schema()))
	}

	// Check Rows
	for i, expectedRow := range expectedRows {
		var row []bigquery.Value
		err := p.Next(&row)
		if err != nil {
			t.Fatalf("row %d: unexpected error: %v", i, err)
		}
		if len(row) != len(expectedRow) {
			t.Errorf("row %d: expected length %d, got %d", i, len(expectedRow), len(row))
		}
		for j, val := range row {
			if val != expectedRow[j] {
				t.Errorf("row %d, col %d: expected %v, got %v", i, j, expectedRow[j], val)
			}
		}
	}

	// Check Done
	var lastRow []bigquery.Value
	err := p.Next(&lastRow)
	if err != iterator.Done {
		t.Errorf("expected iterator.Done, got %v", err)
	}
}

func TestMockRowProvider(t *testing.T) {
	schema := bigquery.Schema{{Name: "Col1"}, {Name: "Col2"}}
	rows := [][]bigquery.Value{
		{"val1", 10},
		{"val2", 20},
	}

	mock := &MockRowProvider{
		schema: schema,
		rows:   rows,
	}

	testRowProviderContract(t, mock, schema, rows)
}

func TestCalculateRowWidths(t *testing.T) {
	tests := []struct {
		name           string
		schema         bigquery.Schema
		rows           [][]bigquery.Value
		expectedWidths []int
	}{
		{
			name: "mixed widths",
			schema: bigquery.Schema{
				{Name: "Name"},
				{Name: "Age"},
			},
			rows: [][]bigquery.Value{
				{"Alice", 30},
				{"Bob", 250},
			},
			expectedWidths: []int{5, 3}, // "Alice" (5), "Age" (3) vs "250" (3)
		},
		{
			name: "column names longer than data",
			schema: bigquery.Schema{
				{Name: "Department"},
				{Name: "ID"},
			},
			rows: [][]bigquery.Value{
				{"HR", 1},
				{"IT", 2},
			},
			expectedWidths: []int{10, 2}, // "Department" (10), "ID" (2)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockRowProvider{
				schema: tt.schema,
				rows:   tt.rows,
			}

			// This should fail to compile as CalculateRowWidths still expects *bigquery.RowIterator
			widths, err := CalculateRowWidths(mock)
			if err != nil {
				t.Fatalf("CalculateRowWidths failed: %v", err)
			}

			if len(widths) != len(tt.expectedWidths) {
				t.Fatalf("expected %d widths, got %d", len(tt.expectedWidths), len(widths))
			}

			for i, w := range widths {
				if w != tt.expectedWidths[i] {
					t.Errorf("index %d: expected width %d, got %d", i, tt.expectedWidths[i], w)
				}
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"a greater than b", 10, 5, 10},
		{"b greater than a", 5, 10, 10},
		{"a equals b", 7, 7, 7},
		{"negative numbers", -1, -5, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := max(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("max(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestPrintFormatedRow(t *testing.T) {
	tests := []struct {
		name     string
		row      []bigquery.Value
		widths   []int
		expected string
	}{
		{
			name:     "simple row",
			row:      []bigquery.Value{"Alice", 30},
			widths:   []int{10, 5},
			expected: "Alice      | 30   \n",
		},
		{
			name:     "single column",
			row:      []bigquery.Value{"OnlyOne"},
			widths:   []int{10},
			expected: "OnlyOne   \n",
		},
		{
			name:     "empty values",
			row:      []bigquery.Value{"", 0},
			widths:   []int{5, 3},
			expected: "      | 0  \n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			PrintFormatedRow(&buf, tt.row, tt.widths)
			result := buf.String()
			if result != tt.expected {
				t.Errorf("PrintFormatedRow() = %q; expected %q", result, tt.expected)
			}
		})
	}
}
