CREATE TABLE IF NOT EXISTS users ( id varchar PRIMARY KEY, firstname varchar(50) NOT NULL, lastname varchar(50)NOT NULL, email varchar(50) UNIQUE,password varchar NOT NULL, username varchar UNIQUE, usertype varchar NOT NULL);

