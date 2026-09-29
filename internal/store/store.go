// Package store はタスクファイルの保存先と読み書きを扱う。
package store

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/tomotargz/memotodo/internal/task"
)

// DataDir は環境変数とホームディレクトリから保存先ディレクトリを決める。
// ホームディレクトリは環境変数で決まらないときだけ取得する。
func DataDir(getenv func(string) string, home func() (string, error)) (string, error) {
	if dir := getenv("MEMOTODO_DIR"); dir != "" {
		return dir, nil
	}
	if xdg := getenv("XDG_DATA_HOME"); xdg != "" {
		return filepath.Join(xdg, "memotodo"), nil
	}
	h, err := home()
	if err != nil {
		return "", err
	}
	return filepath.Join(h, ".local", "share", "memotodo"), nil
}

// Store は保存先ディレクトリ配下のタスクファイルを扱う。
type Store struct {
	Dir string
}

// Add は新しい ID でタスクを作成して保存する。
func (s Store) Add(title string, created time.Time) (task.Task, error) {
	dir := filepath.Join(s.Dir, "tasks")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return task.Task{}, err
	}
	id, err := nextID(dir)
	if err != nil {
		return task.Task{}, err
	}
	t := task.Task{ID: id, Title: title, Status: task.StatusTodo, Created: created}
	path := filepath.Join(dir, fmt.Sprintf("%04d.md", id))
	if err := os.WriteFile(path, task.Marshal(t), 0o644); err != nil {
		return task.Task{}, err
	}
	return t, nil
}

// nextID はファイル名から読み取った ID の最大値 + 1 を返す。
func nextID(dir string) (int, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, err
	}
	maxID := 0
	for _, e := range entries {
		name, ok := strings.CutSuffix(e.Name(), ".md")
		if !ok {
			continue
		}
		id, err := strconv.Atoi(name)
		if err != nil || id <= 0 {
			continue
		}
		maxID = max(maxID, id)
	}
	return maxID + 1, nil
}
