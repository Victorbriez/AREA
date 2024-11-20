CREATE TYPE action_type AS ENUM ('TRIGGER', 'ACTION');
CREATE TYPE action_method AS ENUM ('GET', 'POST', 'PUT', 'DELETE', 'OPTION');

CREATE TABLE IF NOT EXISTS users (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email VARCHAR UNIQUE,
    password VARCHAR,
    username VARCHAR,
    created_at TIMESTAMP DEFAULT NOW(),
    admin BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS providers (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    client_id VARCHAR(255) NOT NULL,
	client_secret VARCHAR(255) NOT NULL,
	redirect_url VARCHAR(255) NOT NULL,
	auth_endpoint VARCHAR(255) NOT NULL,
	token_endpoint VARCHAR(255) NOT NULL,
	device_code_endpoint VARCHAR(255) NOT NULL,
	user_info_endpoint VARCHAR(255) NOT NULL,
	user_id_field VARCHAR(255) NOT NULL,
	user_email_field VARCHAR(255) NOT NULL,
	user_name_field VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS scopes (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    scope VARCHAR(255) NOT NULL,
    provider_id INT NOT NULL,
    required BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (provider_id) REFERENCES providers(id) ON DELETE CASCADE,
    UNIQUE (scope, provider_id)
);

CREATE TABLE IF NOT EXISTS user_scopes (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    scope_id INT NOT NULL,
    provider_id INT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (provider_id) REFERENCES providers(id) ON DELETE CASCADE,
    FOREIGN KEY (scope_id) REFERENCES scopes(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE (scope_id, user_id)
);

CREATE TABLE IF NOT EXISTS tokens (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    access_token TEXT NOT NULL,
    refresh_token TEXT NOT NULL,
    expiry TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS user_providers (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    external_account_id TEXT NOT NULL,
    external_account_name TEXT NOT NULL,
    provider_id INT NOT NULL,
    user_id INT NOT NULL,
    token_id INT NOT NULL,
    FOREIGN KEY (provider_id) REFERENCES providers(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (token_id) REFERENCES tokens(id) ON DELETE CASCADE,
    UNIQUE (external_account_id, provider_id)
);

CREATE TABLE IF NOT EXISTS logs (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    event VARCHAR NOT NULL,
    log_time TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS actions (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR NOT NULL,
    description TEXT NOT NULL,
    type action_type NOT NULL,
    method action_method NOT NULL,
    url VARCHAR NOT NULL,
    body VARCHAR,
    scope_id INT,
    FOREIGN KEY (scope_id) REFERENCES scopes(id)
);

CREATE TABLE IF NOT EXISTS action_fields (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR NOT NULL,
    json_path VARCHAR NOT NULL,
    is_input BOOLEAN NOT NULL,
    action_id INT,
    FOREIGN KEY (action_id) REFERENCES actions(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS flow_steps (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    previous_step INT,
    next_step INT,
    action_id INT NOT NULL,
    FOREIGN KEY (previous_step) REFERENCES flow_steps(id) ON DELETE SET NULL,
    FOREIGN KEY (next_step) REFERENCES flow_steps(id) ON DELETE SET NULL,
    FOREIGN KEY (action_id) REFERENCES actions(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS flows (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    first_step INT NOT NULL,
    run_every INT NOT NULL,
    next_run_at TIMESTAMP NOT NULL,
    active BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (first_step) REFERENCES flow_steps(id) ON DELETE CASCADE
);

ALTER TABLE flow_steps
ADD COLUMN flow_id INT default NULL,
ADD CONSTRAINT fk_flow
FOREIGN KEY (flow_id) REFERENCES flows(id) ON DELETE CASCADE;


CREATE TABLE IF NOT EXISTS flow_runs (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    flow_id INT NOT NULL,
    executed_at TIMESTAMP NOT NULL DEFAULT NOW(),
    logs TEXT,
    successful BOOLEAN,
    FOREIGN KEY (flow_id) REFERENCES flows(id) ON DELETE CASCADE
);
