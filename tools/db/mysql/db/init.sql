CREATE DATABASE sample_db;
CREATE USER 'sample_root'@'%' identified by 'password';
CREATE USER 'sample_dev'@'%' identified by 'developer';

GRANT ALL ON *.* TO 'root'@'%' WITH GRANT OPTION;
GRANT ALL ON sample_db.* TO 'sample_root'@'%';
GRANT SELECT,INSERT,UPDATE,DELETE ON sample_db.* TO 'sample_dev'@'%';
FLUSH PRIVILEGES;
