create table if not exists followers (
    user_id bigint not null,
    follower_id bigint not null,
    created_at timestamp(0) with time zone not null default now(),

    primary key (user_id, follower_id),
    foreign key (user_id) references users (id) ON delete cascade,
    foreign key (follower_id) references users (id) ON delete cascade
)