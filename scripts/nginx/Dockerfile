FROM nginx

# Install SSH server
RUN apt-get update && \
    apt-get install -y openssh-server && \
    apt-get install -y unzip && \
    mkdir /var/run/sshd

# Set root password (CHANGE THIS TO A SECURE PASSWORD)
RUN echo 'root:teste123@' | chpasswd
RUN sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config
RUN mkdir -p /etc/nginx/www

# Expose SSH port
EXPOSE 22
EXPOSE 80
EXPOSE 443

# Start SSH and Nginx services
CMD service nginx start  && /usr/sbin/sshd -D