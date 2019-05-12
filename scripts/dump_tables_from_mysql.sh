#!/usr/bin/env bash

sqlDir="`dirname $0`/sql"

mysqldump -h127.0.0.1 -uroot -ppjoc -P3306 -d pay_gateway --skip-opt --add-drop-table=false > `pwd`/$sqlDir/tables.sql
