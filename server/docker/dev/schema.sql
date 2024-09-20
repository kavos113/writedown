create table pages (
    id int(11) not null primary key auto_increment,
    parent_id int(11) not null,
    title text not null,
    body text not null,
    path text not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    creator_name text not null
);
create table tags (
    id int(11) not null primary key auto_increment,
    page_id int(11) not null,
    name text not null
    foreign key (page_id) references pages(id)
);

insert into pages (parent_id, title, body, path, creator_name) values (1, 'Home', 'Welcome to the home page', '/', 'admin');