# Use the official Ubuntu 22.04 image
FROM ubuntu:22.04

ENV DEBIAN_FRONTEND noninteractive

# Update the package lists
RUN apt-get update \
    && apt-get install -y curl \
    && apt-get -y autoclean

RUN mkdir -p /usr/local/nvm

ENV NODE_VERSION v18.17.0
ENV NVM_DIR /usr/local/nvm

# Install Git
RUN apt-get install git -y

# Install Node.js and npm
RUN touch ~/.bash_profile
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
RUN /bin/bash -c "source $NVM_DIR/nvm.sh && nvm install $NODE_VERSION && nvm use --delete-prefix $NODE_VERSION"

ENV NODE_PATH $NVM_DIR/versions/node/$NODE_VERSION/bin
ENV PATH $NODE_PATH:$PATH

WORKDIR /app
CMD ["/bin/bash"]