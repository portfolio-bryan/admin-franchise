CREATE TABLE IF NOT EXISTS company (
  id uuid,
  company_owner_id uuid NOT NULL,
  name VARCHAR(64) NOT NULL,
  tax_number VARCHAR(64) NOT NULL,
  location_id uuid NOT NULL,
  address_location_id uuid NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_location
    FOREIGN KEY(location_id) 
	    REFERENCES locations(id),
  CONSTRAINT fk_address_location
    FOREIGN KEY(address_location_id) 
      REFERENCES locations(id),
  CONSTRAINT fk_company_owner
    FOREIGN KEY(company_owner_id) 
      REFERENCES company_owner(id)
);