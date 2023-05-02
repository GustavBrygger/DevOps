source ~/.bash_profile
printf $1
export IS_PRODUCTION="TRUE"
export IS_AZURE="FALSE"
export DB_PASSWORD=$1

cd /minitwit

# Configure ELK stack
printf "temp:$(openssl passwd -crypt temp)\n" > .htpasswd
sudo chown root filebeat.yml 
sudo chmod go-w filebeat.yml

docker-compose -f docker-compose.yml pull
docker-compose -f docker-compose.yml up -d
docker-compose restart prometheus
docker-compose restart grafana

docker pull $DOCKER_USERNAME/flagtoolimage:latest