package todo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Store はTodoリストの保管場所を表すインターフェイス。
type Store interface {
	Save(*List) error
	Load() (*List, error)
}

// JSONFileStore はTodoリストの保管場所としてJSONファイルを使用する実装。
type JSONFileStore struct {
	FilePath string
}

// NewJSONFileStore はJSONファイルのパスを指定してJSONFileStoreのポインターを得る。
func NewJSONFileStore(filePath string) *JSONFileStore {
	return &JSONFileStore{filePath}
}

// Save はTodoリストをJSONファイルに保存する。
func (s *JSONFileStore) Save(todoList *List) error {
	b, err := json.MarshalIndent(todoList, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(s.FilePath, b, 0644)
}

// Load はTodoリストをJSONファイルから読み込む。
func (s *JSONFileStore) Load() (*List, error) {
	return nil, errors.New("Not implemented")
}
