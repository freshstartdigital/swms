
CREATE TABLE swms
(
  id SERIAL,
  name VARCHAR(255) NOT NULL,
  swms_type VARCHAR(255) NOT NULL,
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


INSERT INTO swms (name, swms_type, file_name, file_path, created_at, updated_at)
VALUES ('SWMS 1', 'swms', 'swms1.pdf', 'swms1.pdf', NOW(), NOW());

INSERT INTO swms (name, swms_type, file_name, file_path, created_at, updated_at)
VALUES ('SWMS 2', 'swms', 'swms2.pdf', 'swms2.pdf', NOW(), NOW());

INSERT INTO swms_data (swms_id, data, version)
VALUES (1, '{"name": "SWMS 1", "swms_type": "swms", "file_name": "swms1.pdf", "file_path": "swms1.pdf"}', 1);

INSERT INTO swms_data (swms_id, data, version)
VALUES (2, '{"name": "SWMS 2", "swms_type": "swms", "file_name": "swms2.pdf", "file_path": "swms2.pdf"}', 1);