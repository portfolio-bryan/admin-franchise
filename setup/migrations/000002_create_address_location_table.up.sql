CREATE TABLE IF NOT EXISTS address_location (
  id uuid,
  address VARCHAR(64) NOT NULL,
  zip_code VARCHAR(64) NOT NULL,
  location_id uuid NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_location
    FOREIGN KEY(location_id) 
	    REFERENCES locations(id)
);