CREATE TABLE users (
	user_id INTEGER PRIMARY KEY,
	verified_at INTEGER,
	username TEXT NOT NULL,
	email TEXT NOT NULL,
	password_hash TEXT NOT NULL,
	created_at TEXT NOT NULL,
	updated_at TEXT,
	deleted_at TEXT,
)
