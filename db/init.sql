CREATE TABLE IF NOT EXISTS accounts
(
  id SERIAL PRIMARY KEY,
  name varchar(32) UNIQUE NOT NULL,
  balance int
);
