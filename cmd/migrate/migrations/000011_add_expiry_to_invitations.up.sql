alter table
    user_invitations
ADD
    column expiry TIMESTAMP(0) with time zone not null;