
use goapp;

create table users(
  `id`         int auto_increment primary key,
  `name`       varchar(255) not null unique,
  `email`      varchar(255) not null unique,
  `password`   varchar(255) not null,  
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create table todoList (
  `id` int auto_increment primary key, 
  `user_id` varchar(255) not null,  
  `todo` varchar(255) not null, 
  `finish` boolean not null,   
  `created_at` timestamp not null DEFAULT CURRENT_TIMESTAMP,

  foreign key (user_id) references todomanager.users(name) ON DELETE CASCADE
);

-- create table session (

-- );

-- insert default user 
-- INSERT INTO users(name, email, password) VALUES("admin", "sample@sample.com", "xxxxxxxx");