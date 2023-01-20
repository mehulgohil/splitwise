create table if not exists people
(
    id int auto_increment
        primary key,
    name varchar(255) charset utf8mb3 not null,
    mobile_number varchar(255) charset utf8mb3 not null,
    constraint people_mobile_number_uindex
        unique (mobile_number)
);

create table if not exists transactions_per_place
(
    id int auto_increment,
    total_amount int not null,
    place varchar(256) charset utf8mb3 not null,
    date varchar(256) not null,
    constraint transactions_per_place_id_uindex
        unique (id)
);

alter table transactions_per_place
    add primary key (id);

create table if not exists transactions
(
    id int auto_increment
        primary key,
    transaction_place_id int not null,
    spent_by int not null,
    owed_to int null,
    amount int not null,
    status varchar(256) null,
    constraint transactions_people_id_fk
        foreign key (spent_by) references people (id),
    constraint transactions_people_id_fk_2
        foreign key (owed_to) references people (id),
    constraint transactions_transactions_per_place_id_fk
        foreign key (transaction_place_id) references transactions_per_place (id)
);

