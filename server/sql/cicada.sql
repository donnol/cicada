USE cicada;

DROP TABLE IF EXISTS t_phone_code;

CREATE TABLE t_phone_code(
    id bigserial NOT NULL PRIMARY KEY,
    phone char(11) NOT NULL,
    code char(6) NOT NULL,
    used BOOLEAN NOT NULL DEFAULT FALSE
);

DROP TABLE IF EXISTS t_user;

CREATE TABLE t_user(
    id bigserial NOT NULL PRIMARY KEY,
    phone char(11) NOT NULL,
    password VARCHAR(256) NOT NULL,
    name VARCHAR(256) NOT NULL
);

DROP TABLE IF EXISTS t_expense;

CREATE TABLE t_expense(
    id bigserial not null PRIMARY KEY,
    user_id int not null,
    pay DECIMAL(12, 2) not null,
    thing VARCHAR(256) not null,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT clock_timestamp(),
    created_on varchar(256) not null
);

DROP TABLE IF EXISTS t_note;

CREATE TABLE t_note(
    id bigserial not null PRIMARY KEY,
    user_id int not null,
    title text not null default '',
    detail text not null,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT clock_timestamp()
);