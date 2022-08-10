sudo systemctl start mongod
sudo service nginx start

gnome-terminal --tab --working-directory=/home/dimitrije/Desktop/drinkstore/frontend -- /bin/bash -c 'npm run serve-proxy; exec /bin/bash -i'
gnome-terminal --tab --working-directory=/home/dimitrije/Desktop/drinkstore/user-service -- /bin/bash -c 'go run main.go; exec /bin/bash -i'
gnome-terminal --tab --working-directory=/home/dimitrije/Desktop/drinkstore/drink-service -- /bin/bash -c 'go run main.go; exec /bin/bash -i'
gnome-terminal --tab --working-directory=/home/dimitrije/Desktop/drinkstore/comment-service -- /bin/bash -c 'go run main.go; exec /bin/bash -i'
gnome-terminal --tab --working-directory=/home/dimitrije/Desktop/drinkstore/purchase-service -- /bin/bash -c 'cargo run; exec /bin/bash -i'

