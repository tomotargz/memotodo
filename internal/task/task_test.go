package task

import (
	"errors"
	"testing"
	"time"
)

func TestNormalizeTitle(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr error
	}{
		{name: "1 つの引数", args: []string{"APIを調査する"}, want: "APIを調査する"},
		{name: "複数の引数はスペースで連結", args: []string{"設計", "レビュー", "を依頼する"}, want: "設計 レビュー を依頼する"},
		{name: "前後の空白を取り除く", args: []string{"  APIを調査する \t"}, want: "APIを調査する"},
		{name: "引数なし", args: nil, wantErr: ErrEmptyTitle},
		{name: "空白のみ", args: []string{"  ", "\t"}, wantErr: ErrEmptyTitle},
		{name: "改行を含む", args: []string{"API\n調査"}, wantErr: ErrTitleNewline},
		{name: "復帰を含む", args: []string{"API\r調査"}, wantErr: ErrTitleNewline},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeTitle(tt.args)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("エラー = %v, want %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("タイトル = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestMarshal(t *testing.T) {
	jst := time.FixedZone("JST", 9*60*60)
	tk := Task{
		ID:      1,
		Title:   "APIを調査する",
		Status:  StatusTodo,
		Created: time.Date(2026, 9, 28, 10, 0, 0, 123, jst),
	}
	want := `---
id: 1
title: APIを調査する
status: todo
created: 2026-09-28T10:00:00+09:00
---

## メモ
`
	if got := string(Marshal(tk)); got != want {
		t.Errorf("Marshal() =\n%s\nwant\n%s", got, want)
	}
}
