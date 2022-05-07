#!/bin/bash -eux

ufw --force enable
ufw allow ssh
ufw allow 5353/tcp
ufw allow 5181/udp
ufw allow 5182/udp

#allow elastic-agent
ufw allow 8220/tcp
ufw allow 8220/udp

