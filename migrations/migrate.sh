#!/bin/bash
set -eE -o functrace

if [[ -z "${DATABASE_USER}" ]] || [[ -z "${DATABASE_PASSWORD}" ]] || [[ -z "${DATABASE_NAME}" ]]; then
  echo "Some variables are not set, please use '$ source .env'"
  exit 0
fi

failure() {
  local lineno=$1
  local msg=$2
  echo "Failed at line $lineno: $msg"
}
trap 'failure ${LINENO} "$BASH_COMMAND"' ERR

MIGRATIONS="./*-up.sql"
for f in ${MIGRATIONS}
do
  if test -f "${f}"; then
    echo "Running migration $f"
    mysql -h localhost -u "${DATABASE_USER}" --password="${DATABASE_PASSWORD}" "${DATABASE_NAME}" < "${f}"
  fi
done
