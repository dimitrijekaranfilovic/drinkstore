

while getopts s: flag
do
    case "${flag}" in
        s) start=${OPTARG};;
    esac
done

if [[ "$start" == "local" ]];then 
    echo "Starting system localy..."

    sudo systemctl start mongod
    sudo service nginx start

    gnome-terminal --tab --working-directory=/home/dimitrije/Desktop/drinkstore/frontend -- npm run serve-proxy
    gnome-terminal --tab --working-directory=/home/dimitrije/Desktop/drinkstore/user-service -- go run main.go
    gnome-terminal --tab --working-directory=/home/dimitrije/Desktop/drinkstore/drink-service -- go run main.go
    gnome-terminal --tab --working-directory=/home/dimitrije/Desktop/drinkstore/comment-service -- go run main.go
    gnome-terminal --tab --working-directory=/home/dimitrije/Desktop/drinkstore/purchase-service -- cargo run


else 
    echo "Starting system with docker containers..."
    sudo service nginx stop
    gnome-terminal --tab  -- sudo docker start -a user-service-database
    gnome-terminal --tab  -- sudo docker start -a drink-service-database
    gnome-terminal --tab  -- sudo docker start -a comment-service-database
    gnome-terminal --tab  -- sudo docker start -a purchase-service-database
    
    gnome-terminal --tab  -- sudo docker start -a user-service
    gnome-terminal --tab  -- sudo docker start -a drink-service
    gnome-terminal --tab  -- sudo docker start -a comment-service
    gnome-terminal --tab  -- sudo docker start -a purchase-service

    gnome-terminal --tab  -- sudo docker start -a frontend
    gnome-terminal --tab  -- sudo docker start -a api-gateway

fi

