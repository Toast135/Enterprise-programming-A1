

CREATE TABLE author (
	id         INT AUTO_INCREMENT NOT NULL,
	title      VARCHAR(128) NOT NULL,
	artist     VARCHAR(255) NOT NULL,
	price      DECIMAL(5,2) NOT NULL,
	PRIMARY KEY (`id`)
  );

  CREATE TABLE notes (
	id         INT AUTO_INCREMENT NOT NULL,
	title      VARCHAR(128) NOT NULL,
	artist     VARCHAR(255) NOT NULL,
	price      DECIMAL(5,2) NOT NULL,
	PRIMARY KEY (`id`)
  );

  
  CREATE TABLE associations (
	id         INT AUTO_INCREMENT NOT NULL,
	title      VARCHAR(128) NOT NULL,
	artist     VARCHAR(255) NOT NULL,
	price      DECIMAL(5,2) NOT NULL,
	PRIMARY KEY (`id`)
  );

  
  CREATE TABLE shared (
	id         INT AUTO_INCREMENT NOT NULL,
	title      VARCHAR(128) NOT NULL,
	artist     VARCHAR(255) NOT NULL,
	price      DECIMAL(5,2) NOT NULL,
	PRIMARY KEY (`id`)
  );
  