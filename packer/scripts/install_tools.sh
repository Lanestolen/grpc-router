#!/bin/bash -eux

## install wireguard
apt-get update -y
apt-get install wireguard -y
apt-get install wget -y

## configure port forwarding ...
modprobe wireguard
systemctl enable wg-quick@wg0
echo "net.ipv4.ip_forward=1" >> /etc/sysctl.conf
sysctl -p

# install net tools like ifconfig
apt-get install net-tools -y
apt-get install ifupdown -y

## install zip and unzip
apt-get install zip  -y
apt-get install unzip -y
apt-get install isc-dhcp-server -y

mkdir -p /home/jensmp/.config/grpc-router
mv /home/jensmp/uploads/config.yml /home/jensmp/.config/grpc-router/config.yml

wget https://github.com/Lanestolen/grpc-router/releases/download/v0.0.8/grpc-router_0.0.8_linux_amd64.tar.gz
mkdir /home/jensmp/grpc-router
tar -xvzf grpc-router_0.0.8_linux_amd64.tar.gz -C /home/jensmp/grpc-router
chmod +x /home/jensmp/grpc-router/grpc-router

cp /home/jensmp/uploads/grpc-router.service /etc/systemd/system/
systemctl daemon-reload
systemctl enable grpc-router.service
