#!/bin/bash
export PGPASSWORD=chess
psql -h localhost --port=9090 --username=chess -a -d chess -f create_tables.sql
