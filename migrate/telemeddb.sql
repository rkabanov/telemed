drop table if exists public.doctors;

create table public.doctors (
    rowid serial primary key,
    id text unique not null,
    name text not null,
    email text not null,
    role text not null,
    speciality text not null
);

drop table if exists public.patients;

create table public.patients (
    rowid serial primary key,
    id text unique not null,
    name text not null,
    age integer not null,
    external boolean not null
);
