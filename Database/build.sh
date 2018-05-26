#!/bin/bash
sudo mysql -u root < New_User.sql
mysql -u tom -p -t < database.sql
