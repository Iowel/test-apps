create table if not exists user_invitations (
    token bytea PRIMARY KEY,
    user_id bigint NOT NULL
);