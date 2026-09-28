package store

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/tomotargz/memotodo/internal/task"
)

func TestDataDir(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
		want string
	}{
		{
			name: "MEMOTODO_DIR を最優先",
			env:  map[string]string{"MEMOTODO_DIR": "/data/memotodo", "XDG_DATA_HOME": "/xdg"},
			want: "/data/memotodo",
		},
		{
			name: "XDG_DATA_HOME 配下",
			env:  map[string]string{"XDG_DATA_HOME": "/xdg"},
			want: "/xdg/memotodo",
		},
		{
			name: "どちらもなければホーム配下",
			env:  map[string]string{},
			want: "/home/user/.local/share/memotodo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getenv := func(key string) string { return tt.env[key] }
			if got := DataDir(getenv, "/home/user"); got != tt.want {
				t.Errorf("DataDir() = %q, want %q", got, tt.want)
			}
		})
	}
}

var created = time.Date(2026, 9, 28, 10, 0, 0, 0, time.FixedZone("JST", 9*60*60))

func TestAdd_最初のタスク(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "notexist")
	s := Store{Dir: dir}

	got, err := s.Add("APIを調査する", created)
	if err != nil {
		t.Fatalf("Add() エラー: %v", err)
	}
	want := task.Task{ID: 1, Title: "APIを調査する", Status: task.StatusTodo, Created: created}
	if got != want {
		t.Errorf("Add() = %+v, want %+v", got, want)
	}
	assertFile(t, filepath.Join(dir, "tasks", "0001.md"), string(task.Marshal(want)))
}

func TestAdd_既存の最大IDの次(t *testing.T) {
	dir := t.TempDir()
	writeFiles(t, filepath.Join(dir, "tasks"), "0001.md", "0003.md", "memo.txt", "abc.md")
	s := Store{Dir: dir}

	got, err := s.Add("次のタスク", created)
	if err != nil {
		t.Fatalf("Add() エラー: %v", err)
	}
	if got.ID != 4 {
		t.Errorf("ID = %d, want 4", got.ID)
	}
	assertFile(t, filepath.Join(dir, "tasks", "0004.md"), string(task.Marshal(got)))
}

func TestAdd_5桁のID(t *testing.T) {
	dir := t.TempDir()
	writeFiles(t, filepath.Join(dir, "tasks"), "9999.md")
	s := Store{Dir: dir}

	got, err := s.Add("5桁", created)
	if err != nil {
		t.Fatalf("Add() エラー: %v", err)
	}
	if got.ID != 10000 {
		t.Errorf("ID = %d, want 10000", got.ID)
	}
	assertFile(t, filepath.Join(dir, "tasks", "10000.md"), string(task.Marshal(got)))
}

func TestAdd_保存先に書き込めない(t *testing.T) {
	blocker := filepath.Join(t.TempDir(), "blocker")
	if err := os.WriteFile(blocker, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	s := Store{Dir: blocker}

	if _, err := s.Add("書けない", created); err == nil {
		t.Error("Add() がエラーを返さなかった")
	}
}

func writeFiles(t *testing.T, dir string, names ...string) {
	t.Helper()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	for _, name := range names {
		if err := os.WriteFile(filepath.Join(dir, name), nil, 0o644); err != nil {
			t.Fatal(err)
		}
	}
}

func assertFile(t *testing.T, path, want string) {
	t.Helper()
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ファイルを読めない: %v", err)
	}
	if string(got) != want {
		t.Errorf("%s の内容 =\n%s\nwant\n%s", path, got, want)
	}
}
