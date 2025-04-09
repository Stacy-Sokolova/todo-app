CREATE TABLE tasks
(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    username varchar(255) NOT NULL unique,
    password varchar(255) NOT NULL
);

CREATE TABLE users_tasks
(
    id SERIAL PRIMARY KEY,
    user_id int references users (id) NOT NULL,
    task_id int references tasks (id) NOT NULL
);