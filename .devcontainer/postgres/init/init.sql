/* create table */

create table user_tbl( 
    id varchar (50) primary key
    , last_name varchar (50) not null
    , first_name varchar (50) not null
    , birth_date date
    , gender varchar (10) CHECK (gender IN ('male', 'female', 'other')) not null
    , created_at timestamp default current_timestamp
);
