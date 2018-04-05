USE cicada;

DROP TABLE IF EXISTS t_phone_code;

CREATE TABLE t_phone_code(
    id int NOT NULL PRIMARY KEY auto_increment,
    phone char(11) NOT NULL,
    code char(6) NOT NULL,
    used BOOLEAN NOT NULL DEFAULT FALSE
)engine=innodb DEFAULT charset=utf8mb4;

DROP TABLE IF EXISTS t_user;

CREATE TABLE t_user(
    id int NOT NULL PRIMARY KEY auto_increment,
    phone char(11) NOT NULL,
    password VARCHAR(256) NOT NULL,
    name VARCHAR(256) NOT NULL
)engine=innodb DEFAULT charset=utf8mb4;

DROP TABLE IF EXISTS t_expense;

CREATE TABLE t_expense(
    id int not null PRIMARY KEY auto_increment,
    user_id int not null,
    pay DECIMAL(12, 2) not null,
    thing VARCHAR(256) not null,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    created_on varchar(256) not null
)engine=innodb default charset=utf8mb4;