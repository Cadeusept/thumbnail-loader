CREATE TABLE thumbnails_cache 
(
    id          integer not null unique,
    url_hash    varchar(255) not null,
    picture     varchar(255) not null,
    Primary key (id AUTOINCREMENT)
);

CREATE TABLE "thumbnails_cache" (
	"id"	INTEGER NOT NULL UNIQUE,
	"url_hash"	VARCHAR(255) NOT NULL,
	"picture"	VARCHAR(255) NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT)
);