-- apiVersion: v1
-- kind: ConfigMap
-- metadata:
--   name: home-idp-postgres-initdb
--   namespace: idp-system
-- data:
--   initdb.sql: |
CREATE DATABASE home-idp;

CREATE TABLE projects (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  creator VARCHAR(100) NOT NULL,
  create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE roles (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL
);

CREATE TABLE policies (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  policy JSON NOT NULL
);

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  role_id INTEGER DEFAULT 0,
  create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (role_id) REFERENCES roles(id)
);

CREATE TABLE userprojectmapping (
  user_id INTEGER NOT NULL,
  project_id INTEGER NOT NULL,
  PRIMARY KEY (user_id, project_id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (project_id) REFERENCES projects(id)
);

CREATE TABLE rolepolicymapping (
  role_id INTEGER NOT NULL,
  policy_id INTEGER NOT NULL,
  PRIMARY KEY (role_id, policy_id),
  FOREIGN KEY (role_id) REFERENCES roles(id),
  FOREIGN KEY (policy_id) REFERENCES policies(id)
);