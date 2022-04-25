#!/bin/sh

backup_exist=$(ls /home/backups | wc -l)
if [ $backup_exist != 0 ]
then
        sudo rm -rf /home/backups/*
fi
day="$(date +'%y':'%m':'%d'-'%T')"
db_backup="mydb_${day}.sql"
sudo mysqldump  -uroot -proot --no-tablespaces phonebook  >/home/backups/${db_backup}
sudo scp -i /home/ubuntu/.ssh/transfer-key /home/backups/${db_backup}  ubuntu@192.168.0.13:/opt/db-backup
