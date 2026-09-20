create table public.accounts
(
    id   varchar(36)  not null
        constraint accounts_pk
            primary key,
    name varchar(200) not null
);
