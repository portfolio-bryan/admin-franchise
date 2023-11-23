CREATE TABLE IF NOT EXISTS transaction_log_trailing (
  event_id uuid,
  data json NOT NULL,
  was_read BOOLEAN DEFAULT FALSE,
  event_type VARCHAR(128) NOT NULL,
  occurred_on TIMESTAMPTZ NOT NULL,
  PRIMARY KEY (event_id)
);