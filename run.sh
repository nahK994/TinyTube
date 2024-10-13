echo "--- AUTH ---"
echo "1) Start APP"
echo "2) Kill APP"
echo "3) Start DB"
echo "4) Kill DB"
echo "------------------"
echo "--- USER MANAGEMENT ---"
echo "5) Start APP"
echo "6) Kill APP"
echo "7) Start DB"
echo "8) Kill DB"
echo "------------------"

read -p "Type: " cmd

if [[ $cmd -le 4 ]]; then
    cd backend/services/auth/
elif [[ $cmd -gt 4 && $cmd -le 8 ]]; then
    cd backend/services/user-management/
else
    echo "Invalid option"
    exit 1
fi


if [[ $((cmd % 4)) == 1 ]]; then
    cd cmd/run/
    go run main.go
    cd ../..
elif [[ $((cmd % 4)) == 2 ]]; then
    cd cmd/kill/
    go run main.go
    cd ../..
elif [[ $((cmd % 4)) == 3 ]]; then
    docker compose -f db.yml up -d
elif [[ $((cmd % 4)) == 0 ]]; then
    docker compose -f db.yml down
fi

cd ../../..
