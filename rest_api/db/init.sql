create table post(
    id serial unique not null,
    title varchar(128),
    content text,
    primary key (id)
);

insert into post(title, content) values('hello', 'world');