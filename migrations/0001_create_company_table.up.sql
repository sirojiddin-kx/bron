create table company (
    guid uuid not null primary key,
    name varchar(200) not null,
    logo varchar,
    login varchar not null,
    password varchar not null,
    address varchar,
    service_type varchar,
    created_at timestamp default current_timestamp,
    updated_at timestamp
);

INSERT INTO company(
    guid, 
    name, 
    logo, 
    login, 
    password, 
    address, 
    service_type, 
    created_at
    ) values (
        'b7ab5845-70e9-4632-84eb-80b59319cf8d',
        'Tapor',
        'loge',
        'tapor_admin',
        'tapor_admin1234',
        'Shevchenko ko''chasi',
        'barber shop',
        current_timestamp
        );

create table company_service (  
    guid uuid not null primary key,
    title varchar(200) not null,
    duration integer not null,
    price int not null,
    company_id uuid references company(guid) on delete set null,
    created_at timestamp default current_timestamp,
    updated_at timestamp
);

create table employee (
    guid uuid not null primary key,
    first_name varchar(200) not null,
    last_name varchar(200) not null,
    description varchar not null,
    start_time time not null,
    end_time time not null,
    company_id uuid references company(guid) on delete set null,
    create_at timestamp default current_timestamp,
    updated_at timestamp
);

create table employee_service (
    guid uuid not null primary key,
    employee_id uuid references employee(guid) on delete set null,
    company_service_id uuid references company_service(guid) on delete set null,
    created_at timestamp default current_timestamp,
    updated_at timestamp

);

create table client (
    guid uuid not null primary key,
    first_name varchar(200) not null,
    last_name varchar(200) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp
);

create table orders (
    guid uuid not null primary key,
    employee_id uuid references employee(guid),
    client_id uuid references client(guid),
    services uuid[],
    start_time timestamp,
    end_time timestamp,
    create_at timestamp default current_timestamp,
    updated_at timestamp
);

alter table employee
add column position varchar(32);

alter table client
add column phone_number varchar(20);


alter table client
add column company_id uuid;

alter table orders
add column company_id uuid;

alter table orders
add column order_date date;

alter table orders
add column total_price integer;

