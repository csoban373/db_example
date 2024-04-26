CREATE TABLE person
(
  id              INT unsigned NOT NULL AUTO_INCREMENT, # Unique ID for the record
  name            VARCHAR(150) NOT NULL,                # Name of the person
  age             INT unsigned NOT NULL,                # Age of the person
  PRIMARY KEY     (id)                                  # Make the id the primary key
);

INSERT INTO person (name, age) VALUES
  ( 'Sandy', 55 ),
  ( 'Cookie', 12 ),
  ( 'Charlie', 34 );