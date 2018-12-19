package todo

// List はTodoリストを表す。
type List struct {
	Items []*Item
}

// Item はTodoリストの項目を表す。
type Item struct {
	Name string
	Done bool
}

// NewItem は指定した名前のItemを作成してポインターを取得する。
func NewItem(name string) *Item {
	return &Item{Name: name}
}
