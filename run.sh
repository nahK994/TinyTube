#!/bin/bash

echo "--- AUTH ---"
echo "1) Start APP"
echo "2) Kill APP"
echo "3) Start DB"
echo "4) Kill DB"
echo "5) Run PLAYGROUND"
echo "------------------"
echo "--- USER MANAGEMENT ---"
echo "6) Start APP"
echo "7) Kill APP"
echo "8) Start DB"
echo "9) Kill DB"
echo "10) Run PLAYGROUND"
echo "------------------"
echo "--- RABBITMQ ---"
echo "31) Start"
echo "32) Kill"
echo "------------------"
read -p "Type: " cmd

if [[ $cmd == 31 ]]; then
    pushd backend/ || exit
    docker compose -f rabbitmq.yml up -d && echo "RabbitMQ started"
    popd || exit
    exit 0
elif [[ $cmd == 32 ]]; then
    pushd backend/ || exit
    docker compose -f rabbitmq.yml down && echo "RabbitMQ stopped"
    popd || exit
    exit 0
fi

app_dir=''
db_dir=''

if [[ $cmd -ge 1 && $cmd -le 5 ]]; then
    app_dir=backend/services/auth/
    db_dir=backend/db/auth/
elif [[ $cmd -ge 6 && $cmd -le 10 ]]; then
    app_dir=backend/services/user-management/
    db_dir=backend/db/user-management/
else
    echo "Invalid option"
    exit 1
fi

if [[ $((cmd % 5)) == 1 ]]; then
    pushd "$app_dir/cmd/run/" || exit
    go run main.go && echo "Auth/User-Management APP started"
    popd || exit
elif [[ $((cmd % 5)) == 2 ]]; then
    pushd "$app_dir/cmd/kill/" || exit
    go run main.go && echo "Auth/User-Management APP stopped"
    popd || exit
elif [[ $((cmd % 5)) == 3 ]]; then
    pushd "$db_dir" || exit
    docker compose -f db.yml up -d && echo "Database started"
    popd || exit
    pushd "$db_dir/.." || exit
    docker compose -f pgadmin.yml up -d && echo "pgAdmin started"
    popd || exit
elif [[ $((cmd % 5)) == 4 ]]; then
    pushd "$db_dir" || exit
    docker compose -f db.yml down && echo "Database stopped"
    popd || exit
    pushd "$db_dir/.." || exit
    docker compose -f pgadmin.yml down && echo "pgAdmin stopped"
    popd || exit
elif [[ $((cmd % 5)) == 0 ]]; then
    pushd "$app_dir/cmd/playground/" || exit
    go run main.go && echo "Playground started"
    popd || exit
fi
