# gRPC router
This repository contains the router used for the Defatt platform developed at Aalborg University, which is created as an A/D CTF platform focusing on the Blue team.

The router is designed to be ran on an Ubuntu machine, hence the packer configuration found in the packer folder, the routing is done using simple ip forwarding and IPtables.
The gRPC interface allows the user configure DHCP on the connected interfaces by sending an array of subnet details.
Moreover it also has a gRPC interface which allows for the creation of wireguard servers and adding/removing participants to these servers.
