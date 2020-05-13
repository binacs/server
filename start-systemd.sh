#!/bin/bash

sudo cp ./bin/server /usr/sbin/binacsserver

cat <<EOF | sudo tee /etc/systemd/system/binacsserver.service
[Unit]
Description=BinacsServer
Documentation=https://github.com/BinacsLee/server
[Service]
ExecStart=/usr/sbin/binacsserver start \\
  --configFile=${configfile} 
Restart=on-failure
RestartSec=5
[Install]
WantedBy=multi-user.target
EOF

sudo systemctl stop binacsserver
sudo systemctl daemon-reload
sudo systemctl enable binacsserver
sudo systemctl start binacsserver