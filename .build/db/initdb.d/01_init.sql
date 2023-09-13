USE ca_db;

CREATE TABLE IF NOT EXISTS `user` (
`user_id` INT AUTO_INCREMENT NOT NULL,
`token` VARCHAR(255) NOT NULL,
`name` VARCHAR(8) NOT NULL,
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`user_id`));

CREATE TABLE IF NOT EXISTS `charactor` (
`charactor_id` INT AUTO_INCREMENT NOT NULL,
`name` VARCHAR(255) NOT NULL, PRIMARY KEY (`charactor_id`));


INSERT INTO `user` (`name`, `token`)
VALUES
  ('user1', 'aaaaa'),
  ('user2', 'bbbbb'),
  ('user3', 'ccccc'),
  ('user4', 'ddddd'),
  ('user5', 'eeeee');

INSERT INTO `charactor` (`name`)
VALUES
  ('Charactor A'),
  ('Charactor B'),
  ('Charactor C'),
  ('Charactor D'),
  ('Charactor E');