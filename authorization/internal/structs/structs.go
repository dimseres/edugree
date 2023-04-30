package structs

type WithPagination struct {
	Total   int64
	MaxPage int64
	Data    interface{}
}
