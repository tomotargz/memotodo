# CLAUDE.md

memotodo: タスク管理とメモを統合した CLI アプリ（Go 製、モジュール `github.com/tomotargz/memotodo`）。開発者が日常的に使うツールとして開発する。

## 開発プロセス

詳細は [docs/development-process.md](docs/development-process.md) を参照し、必ず従うこと。要点:

- TDD 中心。1 機能 = 1 Issue = 1 ブランチ = 1 PR。
- 実装前に `docs/specs/` に仕様メモを書き、開発者の承認を得る。
- Red / Green / Refactor の **各ステップ完了時に作業を止めて承認を得る**。承認後にそのステップをコミットする。
- コミットは Conventional Commits（type は英語、本文は日本語）。
- PR 作成後は `/code-review` でセルフレビューする。**マージは開発者が行い、Claude はマージしない。**
- ドキュメント・Issue・PR・コミット本文はすべて日本語で書く。

## プロセス改善

- **機能開発の PR とプロセス改善の PR は必ず分ける。** 機能ブランチにプロセス関連ファイルの変更を含めない。
- PR マージ後、`docs/retrospectives/notes.md`（`.gitignore` 対象、コミットしない）にプロセスについての気づきを 1〜3 行追記する（案を出して開発者と確定）。
- `notes.md` の記録が 5 件たまったら、KPT 形式の振り返りを提案する。
- 振り返りごとに `process/retro-<連番>` ブランチで、振り返りファイルとプロセス文書の更新を含むプロセス改善 PR を作る。プロセスはこの PR 経由でのみ変更する。

## コマンド

- テスト: `go test ./...`
- カバレッジ: `go test -cover ./...`
- フォーマット確認: `gofmt -l .`
- 静的解析: `go vet ./...`
