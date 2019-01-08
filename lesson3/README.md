# lesson3

## 最終成果物

映画 API

```
$ curl -isS http://localhost:1323/films?pretty
HTTP/1.1 200 OK
...

[
  {
    "filmId": 1,
    "title": "ACADEMY DINOSAUR",
    "description": "A Epic Drama of a Feminist And a Mad ...",
    "releaseYear": 2006,
    "languageId": 1
  },
  ...
]
```

- 複数取得：GET /films
- 単一取得：GET /films/:id
- 新規作成：POST /films
- 削除：DELETE /films/:id
- 更新：PUT /films/:id

## 実行方法

DB 起動

```
$ docker-compose up --build
```

API 起動

```
$ go run main.go
```

または以下のコマンドだと go ファイル変更時に自動再起動する。

```
$ realize start
```
