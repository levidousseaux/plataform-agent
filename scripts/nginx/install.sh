docker run -d --name nginx_ssh -p 0.0.0.0:80:80 nginx_ssh
docker inspect -f "{{ .NetworkSettings.IPAddress }}" nginx_ssh
ping 172.17.0.2

# TODO: Change to ssh key approach
sshpass -p teste123@ ssh root@172.17.0.2
