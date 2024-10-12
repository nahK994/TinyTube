echo "1) Start auth"
echo "2) Kill auth"
echo "3) Start auth db"
echo "4) Kill auth db"

read -p "Type: " cmd
if [[ $cmd == 1 ]]; then
    cd backend/services/auth/cmd/
    go run main.go
    cd ../../../..
elif [[ $cmd == 2 ]]; then
    sudo kill -9 $(sudo lsof -t -i:8000)
elif [[ $cmd == 3 ]]; then
    cd backend/
    docker compose -f docker-compose.db.yml up -d
    cd ..
elif [[ $cmd == 4 ]]; then
    cd backend/
    docker compose -f docker-compose.db.yml down
    cd ..
fi
