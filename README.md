# yamltoenv
Yaml to bash envs converter. 

## Quick start:

```
# Install golang:
cd /tmp
sudo wget https://dl.google.com/go/go1.14.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.14.3.linux-amd64.tar.gz
sudo echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile

echo "export PATH=$HOME/go/bin:$PATH" >> ~/.profile
echo "export GOPATH=$HOME/go" >> ~/.profile

# Compile:
go get github.com/shalb/yamltoenv
cd ${GOPATH}/src/github.com/shalb/yamltoenv
go install

# test util
yamltoenv -f test.yaml

# use util in bash script example
./test.sh 
```
