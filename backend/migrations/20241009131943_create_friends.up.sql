CREATE TABLE friends (
	id INTEGER PRIMARY_KEY,
	user_a INTEGER,
	user_b INTEGER, 
	FOREIGN KEY(user_a) REFERENCES users(ID), 
	FOREIGN KEY(user_b) REFERENCES users(ID)
);
