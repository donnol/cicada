CREATE TABLE t_expense(
    id int not null PRIMARY KEY auto_increment,
    user_id int not null,
    pay DECIMAL(12, 2) not null,
    thing VARCHAR(256) not null,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    created_on varchar(256) not null
)engine=innodb default charset=utf8mb4;