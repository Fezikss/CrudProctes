create table category(
    id uuid primary key,
    name varchar not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

create table product(
  Id uuid primary key not null ,
  name varchar,
  price int,
  category_id uuid references category(id) ON DELETE CASCADE,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);

