-- SQL dump generated using DBML (dbml-lang.org)
-- Database: MySQL
-- Generated at: 2023-06-20T03:38:54.893Z

-- SQL dump generated using DBML (dbml-lang.org)
-- Database: MySQL
-- Generated at: 2023-06-20T03:38:54.893Z

CREATE TABLE users (
                       username varchar(255) PRIMARY KEY,
                       hashed_password varchar(255) NOT NULL,
                       full_name varchar(255) NOT NULL,
                       email varchar(255) UNIQUE NOT NULL,
                       is_email_verified bool NOT NULL DEFAULT false,
                       password_changed_at timestamp NOT NULL ,
                       created_at timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE verify_emails (
                               id int PRIMARY KEY,
                               username varchar(255) NOT NULL,
                               email varchar(255) NOT NULL,
                               secret_code varchar(255) NOT NULL,
                               is_used bool NOT NULL DEFAULT false,
                               created_at timestamp NOT NULL DEFAULT (now()),
                               expired_at timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE accounts (
                          id int PRIMARY KEY AUTO_INCREMENT,
                          owner varchar(255) NOT NULL,
                          balance int NOT NULL,
                          currency varchar(10) NOT NULL,
                          created_at timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE entries (
                         id int PRIMARY KEY AUTO_INCREMENT,
                         account_id int NOT NULL,
                         amount int NOT NULL COMMENT 'can be negative or positive',
                         created_at timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE transfers (
                           id int PRIMARY KEY AUTO_INCREMENT,
                           from_account_id int NOT NULL,
                           to_account_id int NOT NULL,
                           amount int NOT NULL COMMENT 'must be positive',
                           created_at timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE sessions (
                          id int PRIMARY KEY AUTO_INCREMENT,
                          username varchar(255) NOT NULL,
                          refresh_token varchar(255) NOT NULL,
                          user_agent varchar(255) NOT NULL,
                          client_ip varchar(255) NOT NULL,
                          is_blocked boolean NOT NULL DEFAULT false,
                          expires_at timestamp NOT NULL,
                          created_at timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX accounts_index_0 ON accounts (owner);

CREATE UNIQUE INDEX accounts_index_1 ON accounts (owner, currency);

CREATE INDEX entries_index_2 ON entries (account_id);

CREATE INDEX transfers_index_3 ON transfers (from_account_id);

CREATE INDEX `transfers_index_4` ON `transfers` (`to_account_id`);

CREATE INDEX `transfers_index_5` ON `transfers` (`from_account_id`, `to_account_id`);

ALTER TABLE `verify_emails` ADD FOREIGN KEY (`username`) REFERENCES `users` (`username`);

ALTER TABLE `accounts` ADD FOREIGN KEY (`owner`) REFERENCES `users` (`username`);

ALTER TABLE `entries` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`from_account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`to_account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `sessions` ADD FOREIGN KEY (`username`) REFERENCES `users` (`username`);
