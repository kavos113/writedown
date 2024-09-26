create table users (
    id int(11) not null primary key auto_increment,
    username text not null,
    password text not null
);
create table pages (
    id int(11) not null primary key auto_increment,
    parent_id int(11) not null,
    name text not null,
    body text,
    path text not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    creator_user_id int(11),
    foreign key (creator_user_id) references users(id)
);
create table tags (
    id int(11) not null primary key auto_increment,
    page_id int(11) not null,
    name text not null,
    foreign key (page_id) references pages(id) on delete cascade
);

insert into pages (parent_id, name, path) values (0, 'Home', '/');