#!/bin/bash
set -eE -o functrace

if [[ -z "${DATABASE_USER}" ]] || [[ -z "${DATABASE_PASSWORD}" ]]; then
  echo "Some variables are not set, please use '$ source .env'"
  exit 0
fi

failure() {
  local lineno=$1
  local msg=$2
  echo "Failed at line $lineno: $msg"
}
trap 'failure ${LINENO} "$BASH_COMMAND"' ERR

mysql -u root -e "CREATE USER IF NOT EXISTS '${DATABASE_USER}'@'localhost' IDENTIFIED BY '${DATABASE_PASSWORD}';"
mysql -u root -e "GRANT ALL PRIVILEGES ON *.* TO '${DATABASE_USER}'@'localhost';"
mysql -u root -e "CREATE DATABASE petfeederdb";