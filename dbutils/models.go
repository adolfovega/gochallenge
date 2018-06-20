package dbutils

/*
const assignee = `
			 CREATE TABLE IF NOT EXISTS assignee (
			 	ID 			SERIAL	PRIMARY KEY,
			 	NAME 		VARCHAR	NOT NULL,
			 	LASTNAME 	VARCHAR	NOT NULL
			 	)
`*/

const task = `
			 CREATE TABLE IF NOT EXISTS tasks (
			 	ID 			SERIAL		PRIMARY KEY,
			 	NAME 		VARCHAR(25)	NOT NULL,
			 	PRIORITY 	INT,
			 	DESCRIPTION	VARCHAR(50),
			 	START_TIME	TIME,
			 	END_TIME 	TIME
			 	)
`