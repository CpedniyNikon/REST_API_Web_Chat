create table userdata(
    id serial unique not null,
    login text not null,
    password text not null);

create table message(
    id serial unique not null,
    text_message text not null,
    user_id int references userdata (id));

