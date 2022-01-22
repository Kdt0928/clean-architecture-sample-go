-- // First migration.
-- Migration SQL that makes the change goes here.
CREATE TABLE products (
    id int auto_increment,
    product_code VARCHAR(10),
    product_name VARCHAR(50),
    start_date DATE,
    end_date DATE,
    index(id)
);

-- //@UNDO
-- SQL to undo the change goes here.


