# express　+ prisma

## 初期構築 メモ
これは空のディレクトリから構築するときに入力したコマンド
```
$ npm init -y
$ npm i -D typescript @types/node ts-node
$ npm i express
$ npm i -D @types/express
$ npm i -D ts-node-dev 
$ npx prisma init
```

## 環境構築

.envファイルをテンプレートからコピーして編集する
```
$ cp env-template .env
```

以下のコマンドを順番に入力する
```
$ docker compose run --rm app npm ci
$ docker compose up 
$ docker compose run --rm app npx prisma migrate dev --name init
```

参考:
- [expressの開発にTypeScriptを利用する](https://qiita.com/zaburo/items/69726cc42ef774990279)
- [Connect your database](https://www.prisma.io/docs/getting-started/setup-prisma/start-from-scratch/relational-databases/connect-your-database-typescript-mysql)

# 既存のスキーマをprismaに反映させる
```
$ docker compose exec app npx prisma db pull
$ docker compose exec app npx prisma generate
```
https://www.prisma.io/docs/concepts/components/introspection

# prisma error handling
- [Handling exceptions and errors](https://www.prisma.io/docs/concepts/components/prisma-client/handling-exceptions-and-errors)

- [Errors reference](https://www.prisma.io/docs/reference/api-reference/error-reference)


# trouble shooting

```
$ npx prisma migrate dev --name init

Error: Migration engine error:
Error querying the database: Unknown authentication plugin `sha256_password'.

$ docker compose exec db /bin/bash  
bash-4.4# mysql -uroot -p

mysql> select user, host, plugin from mysql.user;
+------------------+-----------+-----------------------+
| user             | host      | plugin                |
+------------------+-----------+-----------------------+
| app              | %         | caching_sha2_password |
| healthchecker    | localhost | caching_sha2_password |
| mysql.infoschema | localhost | caching_sha2_password |
| mysql.session    | localhost | caching_sha2_password |
| mysql.sys        | localhost | caching_sha2_password |
| root             | localhost | caching_sha2_password |
+------------------+-----------+-----------------------+
6 rows in set (0.01 sec)
```

my.cnfの設定を以下のようにする
```
[mysqld]
character-set-server=utf8
default_authentication_plugin=mysql_native_password
```

再度docker composeをするとpluginがmysql_native_passwordになった  
※ 変更されない場合は`mysql/data`を消す
```
mysql> select user,host,plugin from mysql.user;
+------------------+-----------+-----------------------+
| user             | host      | plugin                |
+------------------+-----------+-----------------------+
| app              | %         | mysql_native_password |
| root             | %         | mysql_native_password |
| mysql.infoschema | localhost | caching_sha2_password |
| mysql.session    | localhost | caching_sha2_password |
| mysql.sys        | localhost | caching_sha2_password |
| root             | localhost | mysql_native_password |
+------------------+-----------+-----------------------+
6 rows in set (0.01 sec)
```

# メモ
必要そう
https://zenn.dev/nori_k/articles/45399999ff39f2#prisma-client%E3%82%92%E5%B0%8E%E5%85%A5%E3%81%99%E3%82%8B


