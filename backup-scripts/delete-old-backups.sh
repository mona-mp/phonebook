#!/bin/sh

file_num_2=$(ls /opt/db-backup | wc -l)
if [ $file_num_2 -ge 2 ]
then
        echo "delete "
        shopt -s extglob
        cd /opt/db-backup
        pwd
        file_name=$(ls   /opt/db-backup/ | tail -n 1)
        echo $file_name
        sudo rm -- !("$file_name")
        cd
fi
