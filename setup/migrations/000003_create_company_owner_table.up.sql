CREATE TABLE IF NOT EXISTS company_owner (
  id uuid,
  first_name VARCHAR(64) NOT NULL,
  last_name VARCHAR(64) NOT NULL,
  email VARCHAR(64) NOT NULL,
  phone VARCHAR(64) NOT NULL,
  location_id uuid NOT NULL,
  address_location_id uuid NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  CONSTRAINT fk_location
    FOREIGN KEY(location_id) 
	    REFERENCES locations(id),
  CONSTRAINT fk_address_location
    FOREIGN KEY(address_location_id) 
      REFERENCES locations(id)
);