#!/bin/bash

# source <(curl -s https://raw.githubusercontent.com/Collieries/system/main/go)

ver="1.22.0"
if [ "$(go version)" != "go version $ver linux/amd64" ]; then \
wget "https://golang.org/dl/go$ver.linux-amd64.tar.gz" && \
sudo rm -rf /usr/local/go && \
sudo tar -C /usr/local -xzf "go$ver.linux-amd64.tar.gz" && \
rm "go$ver.linux-amd64.tar.gz" && \
echo "export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin" >> $HOME/.bash_profile && \
source $HOME/.bash_profile ; \
fi

go version
