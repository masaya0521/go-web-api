# 初回だけ

docker volume create pq_todo

# 起動（コンテナイメージ作成処理で init フォルダ以下が/docker-entrypoint-initdb.d/にコピーされ、create と copy が実行される）

docker-compose up --build

# postgres の中に入る

docker exec -it postgre-todo bash
psql -U postgres

# 停止

#docker-compose down
