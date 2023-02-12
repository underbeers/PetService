CREATE TABLE pet_type
(
    id       serial primary key,
    pet_type varchar(255) not null
);

CREATE TABLE breed
(
    id          serial primary key,
    pet_type_id int references pet_type (id) ON DELETE CASCADE,
    breed_name  varchar(255)
);

CREATE TABLE pet_card
(
    id            serial primary key,
    pet_type_id   int references pet_type (id) ON DELETE CASCADE,
    user_id       int,
    pet_name      varchar(255),
    breed_id      int references breed (id) ON DELETE CASCADE,
    photo         varchar(255),
    birth_date    timestamp,
    male          boolean,
    color         varchar(255),
    care          varchar(255),
    pet_character varchar(255),
    pedigree      varchar(255),
    sterilization boolean,
    vaccinations  boolean
);