CREATE TABLE IF NOT EXISTS address_location (
  id uuid,
  `address` VARCHAR(64) NOT NULL,
  zip_code VARCHAR(64) NOT NULL,
  PRIMARY KEY (id)
  CONSTRAINT fk_user
    FOREIGN KEY(user_id) 
	    REFERENCES users(id)
);