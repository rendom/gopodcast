package main

import "github.com/jmoiron/sqlx"

var sqlSchema = `CREATE TABLE IF NOT EXISTS users
(
	id integer not null primary key,
	name varchar not null,
	email varchar not null unique,
	password varchar not null,
	createdAt datetime
);

CREATE TABLE IF NOT EXISTS podcasts
(
	id integer not null primary key,
	name varchar not null,
	author varchar,
	feed_URL varchar not null unique,
	description varchar,
	image_URL varchar,
	pub_date datetime,
	created_at datetime not null,
	updated_at datetime not null,
	latest_fetch datetime
);

CREATE TABLE IF NOT EXISTS episodes
(
	id integer not null primary key,
	title varchar not null,
	description varchar,
	image varchar,
	podcast_id integer not null,
	created_at datetime not null
);
CREATE INDEX episodes_podcast_id_idx ON episodes (podcast_id);
`

func sqlMigrate(db *sqlx.DB) {
	db.MustExec(sqlSchema)
}
