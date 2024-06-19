CREATE TABLE users (
  id serial primary KEY,
  first_name varchar(255) NOT NULL,
  middle_name varchar(255) DEFAULT NULL,
  last_name varchar(255) DEFAULT NULL,
  username varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  status varchar(255) NOT NULL,
  phone varchar(255) DEFAULT NULL,
  phone_verified_at timestamp NULL DEFAULT NULL,
  email varchar(255) DEFAULT NULL,
  email_verified_at timestamp NULL DEFAULT NULL,
  last_active timestamp DEFAULT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(username)
);