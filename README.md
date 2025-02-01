# keep track of payments

## requirements

- xampp php server, with mysql database

## known bugs

- xampp on linux possible bugs:
    - if errors are not reported, try setting display_errors=On in /opt/lampp/etc/php.ini
    - if reset-admin.php not work in linux and say something about using mysql_upgrade, try running: 
        `sudo /opt/lampp/bin/mysql_upgrade -u root -p`
