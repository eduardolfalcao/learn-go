# Install Go Environment

```bash
sudo apt-get update
sudo apt-get -y upgrade

# check the latest stable version here: https://golang.org/dl/
curl -O https://storage.googleapis.com/golang/go1.15.2.linux-amd64.tar.gz
tar -xvf go1.15.2.linux-amd64.tar.gz
sudo mv go /usr/local

# set up Go language environment variables GOROOT , GOPATH and PATH.
echo -e "export GOROOT=/usr/local/go\nexport GOPATH=$HOME/learn-go/go\nexport PATH=$GOPATH/bin:$GOROOT/bin:$PATH" >> ~/.profile
# update current shell session
source ~/.profile

# check installation
go version
go env
```
