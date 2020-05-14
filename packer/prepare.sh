#!/usr/bin/env bash
set -e

# Install necessary packages
apt-get update
apt-get dist-upgrade --yes
apt-get install wget

# Download and extract the server
wget --no-verbose --output-document teamspeak.tar.bz2 https://files.teamspeak-services.com/releases/server/3.12.1/teamspeak3-server_linux_amd64-3.12.1.tar.bz2
tar -xf teamspeak.tar.bz2
mv teamspeak3-server_linux_amd64 teamspeak
