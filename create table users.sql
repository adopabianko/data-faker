-- The first thing you'll need to do is enable the extension
-- create extension "uuid-ossp";

-- To test that it's working
-- select uuid_generate_v1();

CREATE TABLE users(
    -- id uuid DEFAULT uuid_generate_v4() primary key not null,
    id bigserial primary key not null,
    lat float not null,
    long float not null,
    cc_number varchar(255) not null,
    cc_type varchar(255) not null,
    email varchar(255) not null,
    domain_name varchar(255) not null,
    ipv4 varchar(255) not null,
    ipv6 varchar(255) not null,
    password varchar(255) not null,
    jwt text not null,
    phone_number varchar(255) not null,
    mac_address varchar(255) not null,
    url varchar(255) not null,
    username varchar(255) not null,
    toll_free_number varchar(255) not null,
    e_164_phone_number varchar(255) not null,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    unix_time int not null,
    date varchar(255) not null,
    time varchar(255) not null,
    month_name varchar(255) not null,
    year varchar(255) not null,
    day_of_week varchar(255) not null,
    day_of_month varchar(255) not null,
    timestamp varchar(255) not null,
    century varchar(255) not null,
    timezone varchar(255) not null,
    time_period varchar(255) not null,
    word varchar(255) not null,
    sentence text not null,
    paragraph text not null,
    currency varchar(255) not null,
    amount float not null,
    amount_with_currency varchar(255) not null,
    uuid_hyphenated varchar(255) not null,
    uuid_digit varchar(255) not null
)