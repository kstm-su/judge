CREATE DATABASE judgedb;

CREATE TABLE runners (
	id			SERIAL	PRIMARY KEY,
	host		varchar(30)
)

CREATE TABLE users (
	id			SESRIAL	PRIMARY KEY,
	userid		varchar(30),
	username	TEXT	
)

CREATE TABLE contests (
	id			SERIAL	PRIMARY	KEY,
	contestid	varchar(30),
	name		varchar(60),
	begintime	TIMESTAMP,
	endtime		TIMESTAMP
)

CREATE TABLE problems (
	id			SERIAL	PRIMARY KEY,
	cid			SERIAL,
	problemid	varchar(30),
	name		varchar(30),
	statement	TEXT,
	limit		TEXT,
	tid			SERIAL
)

CREATE TABLE tests (
	id			SERIAL	PRIMARY KEY,
	pid			SERIAL,
	input		TEXT,
	output		TEXT
)
