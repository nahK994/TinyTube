echo "1) Start auth"
echo "2) Kill auth"
echo "2) Start auth db"
echo "2) Kill auth db"

read -p "Type: " cmd
if [[ $cmd == 1 ]]; then
    cd backend/services/auth/
    go run .
    cd ../../..
elif [[ $cmd == 2 ]]; then
    cd backend/services/auth/
    sudo kill -9 $(sudo lsof -t -i:8000)
    cd ../../..
elif [[ $cmd == 3 ]]; then
    cd backend/services/auth/
    docker compose -f docker-compose.db.yml up -d
    cd ../../..
elif [[ $cmd == 4 ]]; then
    cd backend/services/auth/
    docker compose -f docker-compose.db.yml down
    cd ../../..
fi
