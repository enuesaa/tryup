# loadii
A CLI tool to watch file changes and execute a task.

## Usage
```console
$ go run . --help
A CLI tool to watch file changes and execute a task

USAGE:
  loadii [global options] command [command options] [arguments...]

COMMANDS:
   exec   exec commands
   run    run app. Currently, this command supports golang app.
   serve  serve instant web server

GLOBAL OPTIONS:
  --watch value, -w value  watch dir (default: ".")
  --help, -h               show help
  --version, -v            print the version
```

## Development Plan
### Planning Usage
```bash
loadii serve # this serve static content
loadii run main.go # this run `go run main.go` and also, do hot reload
loadii exec pnpm build
loadii serve | loadii exec pnpm build
loadii run main.go | loadii -w admin exec pnpm build
```

## モチベーション
- 開発の際にウェブサーバをぱっと立ち上げたい時があるが nginx や apache のセットアップは面倒
- SSG方式のサイトでコンテンツを更新したときに即座にビルド処理を走らせたい時がある (hot reload)
- 上記2つは本来無関係だが使用するタイミングが近く、一緒のコマンドになっていると便利そう
