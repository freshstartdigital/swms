CREATE TABLE organisation_types
(
  id SERIAL,
  type VARCHAR(255) NOT NULL,
  display_name VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE organisations
(
  id SERIAL,
  organisation_type_id INTEGER REFERENCES organisation_types(id) DEFAULT 2,
  name VARCHAR(255) NOT NULL,
  business_address VARCHAR(255) NOT NULL DEFAULT '',
  abn VARCHAR(255) NOT NULL DEFAULT '',
  business_phone VARCHAR(255) NOT NULL DEFAULT '',
  business_email VARCHAR(255) NOT NULL DEFAULT '',
  logo_file_name VARCHAR(255) NOT NULL DEFAULT '',
  stripe_customer_id VARCHAR(255) NOT NULL DEFAULT '',
  stripe_url VARCHAR(255) NOT NULL DEFAULT '',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  account_holder_id INTEGER,
  PRIMARY KEY (id)
);

CREATE TABLE users
(
  id SERIAL,
  organisation_id INTEGER REFERENCES organisations(id),
  email VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL DEFAULT '',
  password VARCHAR(255) NOT NULL, -- Store hashed passwords for security
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  UNIQUE(email)
);

CREATE TABLE swms
(
  id SERIAL,
  user_id INTEGER REFERENCES users(id),
  organisation_id INTEGER REFERENCES organisations(id),
  name VARCHAR(255) NOT NULL,
  swms_type VARCHAR(255) NOT NULL,
  generator_status VARCHAR(255) NOT NULL DEFAULT 'loading',
  swms_data JSONB NOT NULL DEFAULT '{}',
  file_name VARCHAR(255),
  file_path VARCHAR(255),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);

CREATE TABLE sessions
(
  id SERIAL,
  user_id INTEGER REFERENCES users(id),
  session_token VARCHAR(255) NOT NULL,
  ip_address VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  expires_at TIMESTAMPTZ NOT NULL,
  PRIMARY KEY (id),
  UNIQUE(session_token)
);

CREATE TABLE subscription_plans
(
  id SERIAL,
  stripe_plan_id VARCHAR(255),
  stripe_payment_link VARCHAR(255),
  description VARCHAR(255),
  price INTEGER NOT NULL DEFAULT 0,
  duration VARCHAR(255),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);

CREATE TABLE subscriptions
(
  id SERIAL,
  organisation_id INTEGER REFERENCES organisations(id),
  subscription_plan_id INTEGER REFERENCES subscription_plans(id),
  stripe_subscription_id VARCHAR(255),
  current_period_start INTEGER NOT NULL,
  current_period_end INTEGER NOT NULL,
  stripe_status VARCHAR(255),
  cancelled_at INTEGER,
  PRIMARY KEY (id)
);


CREATE TABLE login_attempts
(
  id SERIAL,
  ip_address VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);

CREATE TABLE banned_ips
(
  id SERIAL,
  ip_address VARCHAR(255) NOT NULL,
  expires_at TIMESTAMPTZ NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);

CREATE TABLE password_resets
(
  id SERIAL,
  user_id INTEGER REFERENCES users(id),
  token VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);


-- Insert organisation types
INSERT INTO organisation_types (type, display_name)
VALUES ('ADMIN', 'Administrator');

INSERT INTO organisation_types (type, display_name)
VALUES ('USER', 'User');

-- Insert organisations
INSERT INTO organisations (name, account_holder_id, organisation_type_id, business_address, abn, business_phone, business_email)  
VALUES ('Fresh Start Project', 1, 1, '2/15 Montague Street, Stones Corner, QLD 4120', '52 648 576 390', '1300 980 999', 'george.frilingos@freshstart.edu.au');


-- Insert users
INSERT INTO users (organisation_id, password, email, name)
VALUES (1, 'password', 'ryan.slater@droneanalytics.com.au', 'Ryan Slater');

INSERT INTO users (organisation_id, password, email, name)
VALUES (1, 'password', 'justin.keating@droneanalytics.com.au', 'Justin Keating');

INSERT INTO users (organisation_id, password, email, name)
VALUES (1, 'password', 'george.frilingos@freshstartprojects.com.au', 'George Frilingos');

-- Insert subscription plans
INSERT INTO subscription_plans (stripe_plan_id, stripe_payment_link, description, price, duration)
VALUES ('prod_PMNAB3HqlD1F8e', 'https://buy.stripe.com/test_3cs16Q6e4fGefVSbII', 'Monthly Subscription', 5000, 'month');

