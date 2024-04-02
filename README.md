# go-sample-module-a

GoモジュールサンプルA

## 作成手順

```
go mod init github.com/yagrush/go-sample-module-a
```

## テスト

```
make test
```

## memo

### タグだけをpushする
```
git push origin main --tags
```

### キャッシュを無視してgo install
```
GOPROXY=direct go install <package>@latest
```
