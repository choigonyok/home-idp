apiVersion: v1
kind: ConfigMap
metadata:
  name: home-idp-postgres-initdb
  namespace: idp-system
data:
  initdb.sql: |
    CREATE DATABASE home-idp;

    CREATE TABLE projects (
      id VARCHAR(100) PRIMARY KEY,
      name VARCHAR(100) NOT NULL,
      creator_id VARCHAR(100) NOT NULL,
      create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE roles (
      id VARCHAR(100) PRIMARY KEY,
      name VARCHAR(100) NOT NULL,
      create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE policies (
      id VARCHAR(100) PRIMARY KEY,
      name VARCHAR(100) NOT NULL,
      policy JSON NOT NULL,
      create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE users (
      id FLOAT PRIMARY KEY,
      name VARCHAR(100) NOT NULL,
      role_id VARCHAR(100) NOT NULL,
      create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      FOREIGN KEY (role_id) REFERENCES roles(id)
    );

    CREATE TABLE userprojectmapping (
      user_id FLOAT NOT NULL,
      project_id VARCHAR(100) NOT NULL,
      PRIMARY KEY (user_id, project_id),
      FOREIGN KEY (user_id) REFERENCES users(id),
      FOREIGN KEY (project_id) REFERENCES projects(id)
    );

    CREATE TABLE rolepolicymapping (
      role_id VARCHAR(100) NOT NULL,
      policy_id VARCHAR(100) NOT NULL,
      PRIMARY KEY (role_id, policy_id),
      FOREIGN KEY (role_id) REFERENCES roles(id),
      FOREIGN KEY (policy_id) REFERENCES policies(id)
    );

    CREATE TABLE dockerfiles (
      id VARCHAR(100) PRIMARY KEY,
      image_name VARCHAR(100) NOT NULL,
      image_version VARCHAR(100) NOT NULL,
      creator_id FLOAT NOT NULL,
      trace_id VARCHAR(100) NOT NULL,
      repository VARCHAR(100) NOT NULL,
      content TEXT NOT NULL,
      FOREIGN KEY (creator_id) REFERENCES users(id)
    );

    CREATE TABLE spans (
      trace_id VARCHAR(100) NOT NULL,
      span_id VARCHAR(100) PRIMARY KEY,
      parent_span_id VARCHAR(100) DEFAULT '',
      start_time VARCHAR(100) NOT NULL,
      end_time VARCHAR(100) DEFAULT '',
      status VARCHAR(100) NOT NULL,
      create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );