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

read -p "Type: " cmd

if [[ $cmd -le 5 ]]; then
    cd backend/services/auth/
elif [[ $cmd -gt 6 && $cmd -le 10 ]]; then
    cd backend/services/user-management/
else
    echo "Invalid option"
    exit 1
fi


if [[ $((cmd % 5)) == 1 ]]; then
    cd cmd/run/
    go run main.go
    cd ../..
elif [[ $((cmd % 5)) == 2 ]]; then
    cd cmd/kill/
    go run main.go
    cd ../..
elif [[ $((cmd % 5)) == 3 ]]; then
    docker compose -f db.yml up -d
elif [[ $((cmd % 5)) == 4 ]]; then
    docker compose -f db.yml down
elif [[ $((cmd % 5)) == 0 ]]; then
    cd cmd/playground/
    go run main.go
    cd ../..
fi

cd ../../..
