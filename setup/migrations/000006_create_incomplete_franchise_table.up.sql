CREATE TABLE IF NOT EXISTS incomplete_franchise (
  id uuid,
  name VARCHAR(128) NOT NULL,
  data json NOT NULL,
  was_verified BOOLEAN DEFAULT FALSE,
  url VARCHAR(256) NOT NULL,
  location_id uuid NOT NULL,
  address_location_id uuid NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  deleted_at TIMESTAMPTZ,
  PRIMARY KEY (id)
);