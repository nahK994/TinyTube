echo "1) Start auth"
echo "2) Kill auth"

read -p "Type: " cmd
if [[ $cmd == 1 ]]; then
    cd services/auth/
    go run .
    cd ../..
elif [[ $cmd == 2 ]]; then
    sudo kill -9 $(sudo lsof -t -i:8000)
fi
