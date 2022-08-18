echo "Deleting containers..."
sudo docker system prune

echo "Deleting images..."
sudo docker images --format="{{.Repository}} {{.ID}}" | grep "^drinkstore_" | cut -d' ' -f2 | xargs sudo docker rmi