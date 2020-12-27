package pleasanter

type View struct {
	NearCompletionTime bool         `json:"NearCompletionTime,omitempty"`
	ColumnFilterHash   ColumnFilter `json:"ColumnFilterHash,omitempty"`
	ColumnSorterHash   ColumnSorter `json:"ColumnSorterHash,omitempty"`
}

type ColumnFilter map[string]string

type ColumnSorter map[string]string
