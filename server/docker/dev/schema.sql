create table pages (
    id int(11) not null primary key auto_increment,
    parent_id int(11) not null,
    title text not null,
    body text not null,
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