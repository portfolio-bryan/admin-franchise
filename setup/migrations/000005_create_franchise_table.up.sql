CREATE TABLE IF NOT EXISTS franchise (
  id uuid,
  company_id uuid NOT NULL,
  title VARCHAR(128) NOT NULL,
  site_name VARCHAR(128) NOT NULL,
  description VARCHAR(512) NOT NULL,
  image VARCHAR(256) NOT NULL,
  url VARCHAR(256) NOT NULL,
  protocol VARCHAR(16) NOT NULL,
  domain_jumps SMALLINT NOT NULL,
  server_names TEXT[] NOT NULL,
  domain_creation_date TIMESTAMPTZ NOT NULL,
  domain_expiration_date TIMESTAMPTZ NOT NULL,
  registrant_name VARCHAR(128) NOT NULL,
  registrant_email VARCHAR(128) NOT NULL,
  location_id uuid NOT NULL,
  address_location_id uuid NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  deleted_at TIMESTAMPTZ,
  PRIMARY KEY (id),
  CONSTRAINT fk_location
    FOREIGN KEY(location_id) 
	    REFERENCES locations(id),
  CONSTRAINT fk_address_location
    FOREIGN KEY(address_location_id) 
      REFERENCES address_location(id),
  CONSTRAINT fk_company
    FOREIGN KEY(company_id) 
      REFERENCES company(id)
);