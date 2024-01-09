CREATE TABLE users
(
  id SERIAL,
  password VARCHAR(255) NOT NULL, -- Store hashed passwords for security
  email VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL,
  PRIMARY KEY (id),
  UNIQUE(email)
);


CREATE TABLE swms
(
  id SERIAL,
  user_id INTEGER REFERENCES users(id),
  name VARCHAR(255) NOT NULL,
  swms_type VARCHAR(255) NOT NULL,
  generator_status VARCHAR(255) NOT NULL DEFAULT 'loading',
  file_name VARCHAR(255),
  file_path VARCHAR(255),
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE swms_data
(
  id SERIAL,
  swms_id INTEGER REFERENCES swms(id),
  data JSONB,
  version INTEGER NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE sessions
(
  id SERIAL,
  user_id INTEGER REFERENCES users(id),
  session_token VARCHAR(255) NOT NULL,
  expires_at TIMESTAMPTZ NOT NULL,
  PRIMARY KEY (id),
  UNIQUE(session_token)
);


-- Insert users
INSERT INTO users (password, email, created_at, updated_at)
VALUES ('hashed_password1', 'user1@example.com', NOW(), NOW());

INSERT INTO users (password, email, created_at, updated_at)
VALUES ('hashed_password2', 'user2@example.com', NOW(), NOW());

