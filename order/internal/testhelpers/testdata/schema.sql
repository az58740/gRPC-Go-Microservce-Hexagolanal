CREATE TABLE IF NOT EXISTS Order (
    id serial NOT NULL ,
    customer_id INTEGER NOT NULL,
    status VARCHAR(50) ,
    	CreatedAt INTEGER ,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS OrderItem (
    id INTEGER NOT NULL AUTO_INCREMENT,
    product_code VARCHAR(50) NOT NULL,
    unit_price FLOAT,
    quantity MEDIUMINT,
    PRIMARY KEY (id)
);
