package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/tomotargz/memotodo/internal/store"
	"github.com/tomotargz/memotodo/internal/task"
)

const usage = `使い方: memotodo <コマンド> [引数]

コマンド:
  add <タイトル>  タスクを追加する
`

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(stderr, usage)
		return 1
	}
	switch args[0] {
	case "add":
		return add(args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "エラー: 不明なコマンドです: %s\n", args[0])
		fmt.Fprint(stderr, usage)
		return 1
	}
}

func add(args []string, stdout, stderr io.Writer) int {
	title, err := task.NormalizeTitle(args)
	if err != nil {
		return fail(stderr, err)
	}
	dir, err := store.DataDir(os.Getenv, os.UserHomeDir)
	if err != nil {
		return fail(stderr, err)
	}
	s := store.Store{Dir: dir}
	t, err := s.Add(title, time.Now())
	if err != nil {
		return fail(stderr, err)
	}
	fmt.Fprintf(stdout, "追加しました: %d %s\n", t.ID, t.Title)
	return 0
}

func fail(stderr io.Writer, err error) int {
	fmt.Fprintf(stderr, "エラー: %v\n", err)
	return 1
}
