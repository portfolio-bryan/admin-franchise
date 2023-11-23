CREATE TABLE IF NOT EXISTS incomplete_franchise (
  id uuid,
  data json NOT NULL,
  was_verified BOOLEAN DEFAULT FALSE,
  url VARCHAR(256) NOT NULL,
  location_id uuid NOT NULL,
  address_location_id uuid NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);