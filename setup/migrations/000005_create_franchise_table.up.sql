CREATE TABLE IF NOT EXISTS franchise (
  id uuid,
  company_id uuid NOT NULL,
  name VARCHAR(64) NOT NULL,
  url VARCHAR(256) NOT NULL,
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
      REFERENCES locations(id),
  CONSTRAINT fk_company
    FOREIGN KEY(company_id) 
      REFERENCES company(id)
);