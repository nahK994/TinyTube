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
echo "---- PGADMIN ----"
echo "33) Start"
echo "34) Kill"
echo "------------------"
read -p "Type: " cmd

# Function to handle Docker Compose operations
run_docker_compose() {
    local config_file=$1
    local action=$2
    pushd backend/config/ || exit
    docker compose -f "$config_file" "$action" -d && echo "$config_file ${action}ed"
    popd || exit
}

# RabbitMQ commands
if [[ $cmd == 31 ]]; then
    run_docker_compose "rabbitmq.yml" "up"
elif [[ $cmd == 32 ]]; then
    run_docker_compose "rabbitmq.yml" "down"
fi

# PgAdmin commands
if [[ $cmd == 33 ]]; then
    run_docker_compose "pgadmin.yml" "up"
elif [[ $cmd == 34 ]]; then
    run_docker_compose "pgadmin.yml" "down"
fi

# Set variables for application directories and database configuration files
app_dir=''
db_file=''

if [[ $cmd -ge 1 && $cmd -le 5 ]]; then
    app_dir="backend/services/auth"
    db_file="auth.yml"
elif [[ $cmd -ge 6 && $cmd -le 10 ]]; then
    app_dir="backend/services/user-management"
    db_file="user-management.yml"
else
    echo "Invalid option"
    exit 1
fi

# Function to start or stop the application
run_go_app() {
    local action_dir=$1
    local message=$2
    pushd "$app_dir/cmd/$action_dir/" || exit
    go run main.go && echo "$message"
    popd || exit
}

# Handle commands for starting/stopping app, DB, and playground
case $((cmd % 5)) in
    1)  run_go_app "run" "App started" ;;
    2)  run_go_app "kill" "App stopped" ;;
    3)  
        pushd backend/config/db/ || exit
        docker compose -f "$db_file" up -d && echo "Database started"
        popd || exit
        ;;
    4)  
        pushd backend/config/db/ || exit
        docker compose -f "$db_file" down && echo "Database stopped"
        popd || exit
        ;;
    0)  run_go_app "playground" "Playground started" ;;
    *)  echo "Invalid command" ;;
esac
