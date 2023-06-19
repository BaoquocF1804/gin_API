CREATE TABLE accounts (
                          id int PRIMARY KEY,
                          owner varchar(255) NOT NULL,
                          balance int NOT NULL,
                          currency varchar(10) NOT NULL,
                          created_at timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE entries (
                         id int PRIMARY KEY,
                         accounts_id int,
                         amount int NOT NULL COMMENT 'Can be nagative',
                         created_at timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE transfer (
                          id int PRIMARY KEY,
                          from_account_id int,
                          to_account_id int,
                          amount int NOT NULL COMMENT 'Must be nagative',
                          created_at timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX accounts_index_0 ON accounts (owner);

CREATE INDEX posts_index_1 ON entries (accounts_id);

CREATE INDEX transfer_index_2 ON transfer (from_account_id);

CREATE INDEX transfer_index_3 ON transfer (to_account_id);

CREATE INDEX transfer_index_4 ON transfer (from_account_id, to_account_id);

ALTER TABLE entries ADD FOREIGN KEY (accounts_id) REFERENCES accounts (id);

ALTER TABLE transfer ADD FOREIGN KEY (from_account_id) REFERENCES accounts (id);

ALTER TABLE transfer ADD FOREIGN KEY (to_account_id) REFERENCES accounts (id);