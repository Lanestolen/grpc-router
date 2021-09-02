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

wget https://github.com/Lanestolen/grpc-router/releases/download/v0.0.1/grpc-router_0.0.1_linux_amd64.tar.gz
mkdir /home/jensmp/grpc-router
tar -xvzf grpc-router_0.0.1_linux_amd64.tar.gz -C /home/jensmp/grpc-router
chmod +x /home/jensmp/grpc-router/grpc-router

cp /home/jensmp/uploads/grpc-router.service /etc/systemd/system/
systemctl daemon-reload
systemctl enable grpc-router.service

## install netman service to manage down network interfaces
## pop up version if required
# mkdir /home/vagrant/netman && cd /home/vagrant/netman
# wget https://github.com/mrturkmenhub/netman/releases/download/1.0.4/netman_1.0.4_linux_64-bit.zip
# unzip netman_1.0.4_linux_64-bit.zip && mv netman_1.0.4_linux_64-bit/* /home/vagrant/netman/
# chmod +x /home/vagrant/netman/netman
# cp /home/vagrant/uploads/netman.service /etc/systemd/system/
# systemctl daemon-reload
# systemctl enable netman.service

## install git
# apt-get install git-all -y
#
# ## install wireguard gRPC service
# cd /home/vagrant
# wget https://github.com/aau-network-security/gwireguard/releases/download/v1.0.3/gwireguard_1.0.3_linux_64-bit.zip
# unzip gwireguard_1.0.3_linux_64-bit.zip && mv gwireguard_1.0.3_linux_64-bit/wgsservice /home/vagrant/wg-service
# chmod +x /home/vagrant/wg-service
# rm -rf gwireguard*
# wget -P /home/vagrant/ https://raw.githubusercontent.com/aau-network-security/gwireguard/master/config/config.yml
#
# ## enable wg-service in system daemon
# cp /home/vagrant/uploads/wg-service.service /etc/systemd/system/wg-service.service
# sudo chmod 644  /etc/systemd/system/wg-service.service
# sudo systemctl start wg-service
# sudo systemctl enable wg-service
