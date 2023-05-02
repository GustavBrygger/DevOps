source ~/.bash_profile

cd /minitwit
echo "$3" > .env

# Configure ELK stack
printf "$1:$(openssl passwd -crypt $2)\n" > .htpasswd
sudo chown root filebeat.yml 
sudo chmod go-w filebeat.yml

docker-compose -f docker-compose.yml pull
docker-compose -f docker-compose.yml up -d
docker-compose restart prometheus
docker-compose restart grafana

docker pull $DOCKER_USERNAME/flagtoolimage:latest