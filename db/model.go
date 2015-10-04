package db

type Paging struct {
	Offset int
	Limit  int
	Count  int
	Total  int
}
