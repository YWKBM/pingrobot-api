CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    registered_at TIMESTAMPTZ NOT NULL,
    last_visit_at TIMESTAMPTZ NOT NULL
);

CREATE TYPE service_status as enum('NOT_STATED', 'SUCCESS', 'ERROR');

CREATE TABLE web_services
(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    name VARCHAR(255) NOT NULL,
    link VARCHAR(255) NOT NULL,
    port INT NOT NULL,
    status service_status DEFAULT ('NOT_STATED'),
);