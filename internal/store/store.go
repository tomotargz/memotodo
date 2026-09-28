// Package store はタスクファイルの保存先と読み書きを扱う。
package store

import (
	"time"

	"github.com/tomotargz/memotodo/internal/task"
)

// DataDir は環境変数とホームディレクトリから保存先ディレクトリを決める。
func DataDir(getenv func(string) string, home string) string {
	return ""
}

// Store は保存先ディレクトリ配下のタスクファイルを扱う。
type Store struct {
	Dir string
}

// Add は新しい ID でタスクを作成して保存する。
func (s Store) Add(title string, created time.Time) (task.Task, error) {
	return task.Task{}, nil
}
