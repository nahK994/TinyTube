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
    cd backend/
    docker compose -f rabbitmq.yml up -d 
    cd ..
    exit 0
elif [[ $cmd == 32 ]]; then
    cd backend/
    docker compose -f rabbitmq.yml down
    cd ..
    exit 0
fi


if [[ $cmd -ge 1 && $cmd -le 5 ]]; then
    cd backend/services/auth/ || exit
elif [[ $cmd -ge 6 && $cmd -le 10 ]]; then
    cd backend/services/user-management/ || exit
else
    echo "Invalid option"
    exit 1
fi


if [[ $((cmd % 5)) == 1 ]]; then
    cd cmd/run/ || exit
    go run main.go || exit
    cd ../..
elif [[ $((cmd % 5)) == 2 ]]; then
    cd cmd/kill/ || exit
    go run main.go || exit
    cd ../..
elif [[ $((cmd % 5)) == 3 ]]; then
    docker compose -f db.yml up -d || exit
elif [[ $((cmd % 5)) == 4 ]]; then
    docker compose -f db.yml down || exit
elif [[ $((cmd % 5)) == 0 ]]; then
    cd cmd/playground/ || exit
    go run main.go || exit
    cd ../..
fi

cd ../../..
