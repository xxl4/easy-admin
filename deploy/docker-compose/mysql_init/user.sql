CREATE DATABASE IF NOT EXISTS easy_admin;
CREATE USER easyadmin@'*' IDENTIFIED BY 'easyadmin123';
GRANT ALL ON easyadmin.* TO fate@'*';