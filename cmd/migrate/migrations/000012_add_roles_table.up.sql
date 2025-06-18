CREATE TABLE IF NOT EXISTS roles (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    level int NOT NULL DEFAULT 0,
    description TEXT 
);

insert into
    roles (name, description, level)
VALUES
    (
        'user',
        'A user can create posts and comments',
        1
    );


insert into
    roles (name, description, level)
VALUES
    (
        'moderator',
        'A moderator can update other posts',
        2
    );


insert into
    roles (name, description, level)
VALUES
    (
        'admin',
        'A admin can update and delete other users posts',
        3
    );


