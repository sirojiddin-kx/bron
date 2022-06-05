create table company (
    guid uuid not null primary key,
    name varchar(200) not null,
    logo varchar,
    login varchar not null,
    password varchar not null,
    address varchar,
    service_type_id varchar,
    created_at timestamp default current_timestamp,
    updated_at timestamp
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
    services uuid[],
    start_time timestamp,
    end_time timestamp,
    create_at timestamp default current_timestamp,
    updated_at timestamp
);








