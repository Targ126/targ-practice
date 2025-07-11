create table if not exists space (
    id bigserial primary key,

    name varchar not null default '',

    created_at timestamptz default now() not null,
    updated_at timestamptz default now() not null
);

create index if not exists space_name_idx on space (name);