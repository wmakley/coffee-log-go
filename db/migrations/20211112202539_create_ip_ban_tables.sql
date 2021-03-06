-- migrate:up
CREATE TABLE login_attempts
(
	ip_address VARCHAR(255) NOT NULL PRIMARY KEY,
	attempts   INTEGER      NOT NULL DEFAULT 1,
	created_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE banned_ips
(
	ip_address VARCHAR(255) NOT NULL PRIMARY KEY,
	created_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- migrate:down

DROP TABLE IF EXISTS banned_ips;
DROP TABLE IF EXISTS login_attempts;
