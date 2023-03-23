-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE game
(
    `id`   BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'auto increment id' PRIMARY KEY,
    `name` VARCHAR(50) NULL COMMENT 'name',
    `age`  INTEGER     NOT NULL COMMENT 'code'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

CREATE UNIQUE INDEX idx_name ON game (`name`);