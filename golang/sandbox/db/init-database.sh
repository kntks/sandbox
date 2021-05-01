#!/usr/bin/env bash

#run the setup script to create the DB and the schema in the DB
mysql -udocker -pdocker test_database < "/docker-entrypoint-initdb.d/001-create-tables.sql"