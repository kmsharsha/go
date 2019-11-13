echo "Downloading go packages..."
wget https://dl.google.com/go/go1.11.4.linux-amd64.tar.gz

echo "Unpacking go packages..."
sudo tar -C /usr/local -xzvf go1.11.4.linux-amd64.tar.gz

echo "Updating go path..."
echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOROOT/bin' >> ~/.bashrc

echo "Removing downloaded tar file..."
rm go1.11.4.linux-amd64.tar.gz

echo "Checking go version..."
go version
