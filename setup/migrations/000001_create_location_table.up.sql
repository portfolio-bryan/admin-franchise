CREATE TABLE IF NOT EXISTS locations (
  id uuid,
  country VARCHAR(64) NOT NULL,
  `state` VARCHAR(64) NOT NULL,
  city VARCHAR(64) NOT NULL,
  PRIMARY KEY (id)
);