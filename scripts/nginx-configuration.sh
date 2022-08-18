REQUIRED_PKG="nginx"
PKG_OK=$(dpkg-query -W --showformat='${Status}\n' $REQUIRED_PKG|grep "install ok installed")
echo Checking for $REQUIRED_PKG: $PKG_OK
if [ "" = "$PKG_OK" ]; then
  echo "$REQUIRED_PKG not installed. Install it first, then proceed with setup."
else
    echo "Copying conf files..."
    sudo cp -f ../nginx-setup-local/nginx.conf /etc/nginx
    sudo cp -f ../nginx-setup-local/api_json_errors.conf /etc/nginx
    sudo cp -rf ../nginx-setup-local/drinkstore_conf.d /etc/nginx

    echo "Removing logs..."
    sudo rm -rf /var/log/nginx/user_service /var/log/nginx/drink_service /var/log/nginx/comment_service /var/log/nginx/purchase_service

    sudo mkdir /var/log/nginx/user_service /var/log/nginx/drink_service /var/log/nginx/comment_service /var/log/nginx/purchase_service

    echo "Creating log folders..."
    sudo mkdir /var/log/nginx/user_service/access /var/log/nginx/user_service/error
    sudo mkdir /var/log/nginx/drink_service/access /var/log/nginx/drink_service/error
    sudo mkdir /var/log/nginx/comment_service/access /var/log/nginx/comment_service/error
    sudo mkdir /var/log/nginx/purchase_service/access /var/log/nginx/purchase_service/error

fi
