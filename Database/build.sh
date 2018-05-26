#!/bin/bash
sudo mysql -u root | GRANT ALL PRIVILEGES ON *.* TO 'tom'@'localhost' IDENTIFIED BY 'pwd123';
mysql -u tom -p -t < database.sql
