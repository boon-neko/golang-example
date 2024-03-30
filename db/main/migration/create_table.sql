create table books
(
    id         varchar(50)                         not null
        primary key,
    title      varchar(255)                        not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint books_title_uk
        unique (title)
);

create table users
(
    id         varchar(50)                         not null
        primary key,
    name       varchar(30)                         not null,
    role       int                                 not null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    constraint name_unique
        unique (name)
);

create table user_book_subscription
(
    id         int auto_increment
        primary key,
    user_id    varchar(50)                         not null,
    book_id    varchar(50)                         not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    constraint user_book_subscription_uk
        unique (user_id, book_id),
    constraint user_book_subscription_books_id_fk
        foreign key (book_id) references books (id),
    constraint user_book_subscription_users_id_fk
        foreign key (user_id) references users (id)
);

