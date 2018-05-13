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
	author varchar not null,
	feed_URL varchar not null unique,
	description varchar not null,
	image_URL varchar not null,
	pub_date datetime not null,
	created_at datetime not null,
	updated_at datetime not null,
	latest_fetch datetime
);

CREATE TABLE IF NOT EXISTS episodes
(
	id integer not null primary key,
	guid varchar,
	title varchar not null,
	url varchar not null,
	description varchar,
	image varchar,
	podcast_id integer not null,
	created_at datetime not null
);
CREATE INDEX IF NOT EXISTS episodes_podcast_id_idx ON episodes (podcast_id);
CREATE INDEX IF NOT EXISTS episodes_guid_podcast_id_idx ON episodes (guid, podcast_id);

CREATE TABLE IF NOT EXISTS subscriptions
(
	podcast_id integer not null,
	user_id integer not null,
	created_at datetime
);
CREATE INDEX IF NOT EXISTS subscriptions_podcast_id_user_id_idx ON subscriptions (user_id, podcast_id);
`

func sqlMigrate(db *sqlx.DB) {
	db.MustExec(sqlSchema)
}
