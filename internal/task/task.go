// Package task はタスクのデータモデルとファイル形式を扱う。
package task

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Status はタスクのステータス。
type Status string

const (
	StatusTodo  Status = "todo"
	StatusDoing Status = "doing"
	StatusDone  Status = "done"
)

var (
	ErrEmptyTitle   = errors.New("タイトルを指定してください")
	ErrTitleNewline = errors.New("タイトルに改行は使えません")
)

// Task は 1 つのタスク。
type Task struct {
	ID      int
	Title   string
	Status  Status
	Created time.Time
}

// NormalizeTitle はコマンド引数からタイトルを組み立てて検証する。
func NormalizeTitle(args []string) (string, error) {
	title := strings.TrimSpace(strings.Join(args, " "))
	if title == "" {
		return "", ErrEmptyTitle
	}
	if strings.ContainsAny(title, "\r\n") {
		return "", ErrTitleNewline
	}
	return title, nil
}

// Marshal はタスクをマークダウンファイルの内容に変換する。
func Marshal(t Task) []byte {
	return fmt.Appendf(nil, `---
id: %d
title: %s
status: %s
created: %s
---

## メモ
`, t.ID, t.Title, t.Status, t.Created.Format(time.RFC3339))
}
