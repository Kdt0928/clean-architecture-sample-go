-- // create_user_table
-- Migration SQL that makes the change goes here.
CREATE TABLE users (
    id int auto_increment,
    first_name VARCHAR(20),
    last_name VARCHAR(20),
    index(id)
);

-- //@UNDO
-- SQL to undo the change goes here.
DROP TABLE users;
