# MyBatis Migration
MyBatis MigrationはJava系のマイグレーションツール

## フォルダ構成
```
.
├─drivers            ・・・JDBCドライバを格納場所
└─sample_db          ・・・データベース単位の設定ファイルとDDLを格納場所
    ├─environments   ・・・設定ファイルの格納場所
    └─scripts        ・・・DDLの格納場所
```
## セットアップ
### データベース単位の設定ファイル作成
データベース単位のフォルダを作成して、設定ファイルを作成する。  
```sh
$ docker ps
CONTAINER ID   IMAGE                  COMMAND                  CREATED             STATUS             PORTS                                                  NAMES
c5def9f27be6   mybatis                "/bin/bash"              About an hour ago   Up About an hour                                                          mybatis
$ docker exec -it ${CONTAINER ID} sh # mybatisコンテナに入る
$ mkdir ${データベース名}
$ migrate init # 設定ファイルや変更履歴管理テーブルのcreate文を自動で作成
```

### JDBCドライバのダウンロード
データベースに接続するためのJDBCドライバをダウンロードして、driversに格納する
TODO：データベース単位の設定ファイル作成と合わせてツール化
```sh
$ docker ps
CONTAINER ID   IMAGE                  COMMAND                  CREATED             STATUS             PORTS                                                  NAMES
c5def9f27be6   mybatis                "/bin/bash"              About an hour ago   Up About an hour                                                          mybatis
$ docker exec -it ${CONTAINER ID} sh # mybatisコンテナに入る
$ wget wget http://dev.mysql.com/get/Downloads/Connector-J/mysql-connector-java-5.1.30.tar.gz -O mysql-connector.tar.gz  # 例
$ tar -xvzf mysql-connector.tar.gz
$ cp mysql-connector-java-5.1.30/mysql-connector-java-5.1.30-bin.jar /mybatis/drivers
```

### development.properties修正
* time_zoneを`Asia/Tokyo`に変更
  ```bash
  time_zone=Asia/Tokyo
  ```
* `script_char_set=UTF-8`をアンコメント
* `## JDBC connection properties.`下の設定追記
  ```bash
  driver=com.mysql.cj.jdbc.Driver # driverのクラスを追記
  url=jdbc:mysql://mysql:3306/sample_db # jdbc:${データベースの種類}/${サーバー名}:${ポート番号}/${データベース名}
  username=sample_root
  password=password
  ```

### 変更履歴管理テーブルの作成
コンテナ内で実行する。
TODO ここもツール化したい
```sh
$ cd ${データベース名}
$ migrate status      # 状態を確認
ID             Applied At          Description
==================================================================
20141211083812    ...pending...    create changelog
20141211083813    ...pending...    first migration
$ migrate up          # pendingのものが実行される
$ migrate ver ${ID}   # ID指定で実行する
```
## マイグレーション
### DDLの作成
コンテナ内で実行する。
TODO ここもツール化したい
```sh
$ cd ${データベース名}
$ migrate new ${テーブル名} # 年月日時分秒_${テーブル名}.sqlのファイルが作成される。
```
作成したファイルにDDLを追記
```sql
-- // First migration.
-- 以下に適用したいSQLを記載する
CREATE TABLE user (
    id int not null,
    first_name VARCHAR(20),
    last_name VARCHAR(20)
);

-- //@UNDO
-- 以下に上記で適用したものを元に戻すSQLを記載する（記載しなくてもいい）
DROP TABLE user;
```
### テーブルの作成
作成したDDLを実行する。
```sh
$ cd ${データベース名}
$ migrate up
```

## 参考コマンド
|コマンド|用途|
|----|----|
|migrate up|最新の状態まで一気に進める|
|migrate down|一世代前のDB変更に戻す|
|migrate ver [ID]|[ID]適用後の状態にする|
|migrate status|現状の状態を表示する|