## Go-PostgreSQL

RestAPI のデモ

### 1. テーブルの作成

`./internal/utility/sql/create_tables.sql` にあるファイルを DB に実行

```
# DB接続情報
POSTGRES_USER: postgres
POSTGRES_PASSWORD: postgres
POSTGRES_DB: postgres
POSTGRES_HOSTNAME: localhost
```

### 2. サーバー起動

```shell
go run cmd/main.go
```

### 3. HTTP でアクセス

詳細は `./internal/utility/openapi/users.yml` を参照

| 処理内容                            | HTTP メソッド | URL                       |
| ----------------------------------- | ------------- | ------------------------- |
| ユーザーデータを全て取得する        | GET           | localhost:8080/users      |
| ユーザーを一件登録する              | POST          | localhost:8080/users      |
| id と一致するユーザーを一件取得する | GET           | localhost:8080/users/{id} |
| id と一致するユーザーを一件更新する | PUT           | localhost:8080/users/{id} |
| id と一致するユーザーを一件削除する | DELETE        | localhost:8080/users/{id} |

### [レスポンスメモ](https://blog.pepese.com/design-rest-api/)

#### GET

| HTTP Status Code | 意味                                                     | ボディ         | その他 |
| ---------------- | -------------------------------------------------------- | -------------- | ------ |
| 200 OK           | 正常に処理が完了                                         | 取得対象データ |        |
| 204 No Content   | コレクションは存在するが、対象 ID のリソースが存在しない | 空             |        |

#### POST

| HTTP Status Code | 意味                                                     | ボディ                   | その他                                |
| ---------------- | -------------------------------------------------------- | ------------------------ | ------------------------------------- |
| 200 OK           | 新しいリソース作成以外の処理が正常に完了                 | 処理結果データ           |                                       |
| 201 Created      | 正常に処理が完了                                         | 新規作成リソースデータ   | Location ヘッダに作成したリソース URI |
| 204 No Content   | コレクションは存在するが、対象 ID のリソースが存在しない | 空                       |                                       |
| 400 Bad Request  | クライアントから無効なリソース登録要求                   | エラー内容そのものや URI |                                       |

#### PUT

| HTTP Status Code         | 意味                                   | ボディ                   | その他                                |
| ------------------------ | -------------------------------------- | ------------------------ | ------------------------------------- |
| 200 OK or 204 No Content | リソースの更新処理が正常に完了         | 処理結果データ           |                                       |
| 201 Created              | 正常に処理が完了                       | 新規作成リソースデータ   | Location ヘッダに作成したリソース URI |
| 400 Bad Request          | クライアントから無効なリソース登録要求 | エラー内容そのものや URI |                                       |
| 409 Conflict             | リソースの現状の状態の矛盾している     |                          | 更新順序の前後等により発生            |

#### DELETE

| HTTP Status Code | 意味             | ボディ | その他 |
| ---------------- | ---------------- | ------ | ------ |
| 204 No Content   | 正常に処理が完了 | 空     |        |

#### 共通

| HTTP Status Code          | 意味                                                               | 補足                                                       |
| ------------------------- | ------------------------------------------------------------------ | ---------------------------------------------------------- |
| 301 MovedPermanently      | アクセスしたリソースが別の URI に恒久的に移動した                  | Location ヘッダに移動先の URI を付与する                   |
| 303 See Other             | 別の URI にアクセスしてほしい                                      | Location ヘッダに移動先の URI を付与する                   |
| 304 Not Modified          | リソースの変更をしなかった                                         | POST、PUT、DELETE の結果                                   |
| 307 Temporary Redirect    | 一時的なリダイレクト                                               | 閉塞時等、Location ヘッダにリダイレクト先の URI を付与する |
| 400 Bad Request           | クライアントからの HTTP リクエストに誤りがありサーバで処理できない |                                                            |
| 401 Unauthorized          | クライアントが認証されていない                                     |                                                            |
| 403 Forbidden             | クライアントの権限不足                                             |                                                            |
| 404 Not Found             | アクセスした URI・コレクションが存在しない                         |                                                            |
| 406 Not Acceptable        | リクエストの Accept ヘッダがサーバで受け入れられない               | コンテンツネゴシエーション失敗                             |
| 415 Unsupported MediaTYpe | リクエストの Content-Type ヘッダがサーバで受け入れられない         | サーバ側で Body を解釈できない                             |
| 409 Conflict              | リソースの現状の状態の矛盾している                                 | 更新順序の前後等                                           |
| 500 Internal Server Error | サーバ内でエラーが発生した                                         | 上記以外のサーバエラー                                     |
