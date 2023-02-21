package models

type Operation struct {
	Income      bool
	Expense     bool
	Category    Category
	Description string
	Sum         float64
}
