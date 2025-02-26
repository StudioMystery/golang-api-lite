-- foo_bars definition

CREATE TABLE foo_bars (
	ID INTEGER UNIQUE,
	Foo INTEGER,
	Bar TEXT, 
	created_at DATETIME, 
	updated_at DATETIME, 
	deleted_at DATETIME,
    CONSTRAINT PK_foo_bars PRIMARY KEY (ID)
);