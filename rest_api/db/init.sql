create table userdata(
    id serial unique not null,
    login text not null,
    password text not null,
    is_logged BIT not null default B'0',
    time_logged timestamp without time zone
);

create table message(
    id serial unique not null,
    text_message text not null,
    user_id int references userdata (id)
);