#!/usr/bin/env bash
sqlDir="`dirname $0`/sql"
echo "Sql dir: $sqlDir"
echo "Should init sql scripts: `ls $sqlDir`"
docker run --restart=always --name mysql -v `pwd`/$sqlDir:/docker-entrypoint-initdb.d -v `pwd`/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=pjoc -e MYSQL_USER=pjoc -e MYSQL_PASSWORD=pjoc_pay -e MYSQL_DATABASE=pay_gateway  -p 3306:3306 -d mysql:latest --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
#sql="`cat ddl.sql`"
#docker exec  mysql $sql
