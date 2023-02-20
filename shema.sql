create table product_price (
    id serial primary key,
    product_id int unique not null, 
    price float not null
);