package todo

import (
	"encoding/json"
	"io/ioutil"
	"os"
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
	_, err := os.Stat(s.FilePath)
	if os.IsNotExist(err) {
		return &List{}, nil // ファイルが存在しない場合は空のTodoリストを返す。
	}
	b, err := ioutil.ReadFile(s.FilePath)
	if err != nil {
		return nil, err
	}

	var list List
	err = json.Unmarshal(b, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
