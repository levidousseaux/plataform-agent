# Use the official Ubuntu 22.04 image
FROM ubuntu:22.04

# Set environment variables
ENV DEBIAN_FRONTEND noninteractive

# Update the package lists
RUN apt-get update

# Install Git
RUN apt-get install -y git

# Install Node.js and npm
RUN apt-get install -y nodejs
RUN apt-get install -y npm

# Set the default command to start a bash shell
CMD ["/bin/bash"]