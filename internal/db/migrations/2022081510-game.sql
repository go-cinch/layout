-- +migrate Up
CREATE TABLE `game`
(
    `id`         BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'auto increment id' PRIMARY KEY,
    `created_at` DATETIME(3) NULL COMMENT 'create time',
    `updated_at` DATETIME(3) NULL COMMENT 'update time',
    `name`       VARCHAR(50) COMMENT 'name'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

-- +migrate Down
DROP TABLE `game`;