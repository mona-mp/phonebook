#!/bin/sh


	
#backup_exist=$(ls /opt/db-backup | wc -l)
#if [ $backup_exist != 0 ]
#then
#	sudo rm -rf /opt/db-backup/*
#fi

file_num=$(s3cmd ls -r s3://mona-db-backup | wc -l)
echo $file_num

if [ $file_num -lt 3 ]
then
	 echo "les than 3"
	 new_backup=$(ls /opt/db-backup | tail -n 1 )
	 echo $new_backup
	sudo s3cmd put /opt/db-backup/$new_backup s3://mona-db-backup
fi

if [ $file_num -ge 3 ]
then
	echo "more than 3"
        file_name=$(sudo s3cmd ls s3://mona-db-backup | tail -n 1 | awk '{print $4}')
        sudo s3cmd del $file_name
        new_backup=$(ls /opt/db-backup | tail -n 1 )
        echo $new_backup
        sudo s3cmd put /opt/db-backup/$new_backup s3://mona-db-backup
fi

file_num_2=$(ls /opt/db-backup | wc -l)

#if [ $file_num_2 -ge 2 ]
#then 
#	echo "delete "
#	shopt -s extglob
#	cd /opt/db-backup
#       pwd
#	file_name=$(ls   /opt/db-backup/ | tail -n 1)
#	echo $file_name
#	sudo rm -- !("$file_name")
#	cd	
#fi	
