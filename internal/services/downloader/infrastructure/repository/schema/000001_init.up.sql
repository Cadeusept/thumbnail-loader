CREATE TABLE thumbnails_cache 
(
    id          serial not null unique,
    url_hash    varchar(255) not null,
    picture     varchar(255),
    Primary key (id)
);