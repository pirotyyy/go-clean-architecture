USE ca_db;

CREATE TABLE IF NOT EXISTS `user` (
`id` INT NOT NULL AUTO_INCREMENT,
`token` VARCHAR(255) NOT NULL,
`name` VARCHAR(32) NOT NULL,
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`));

CREATE TABLE IF NOT EXISTS `game_character` (
`id` INT NOT NULL AUTO_INCREMENT,
`name` VARCHAR(255) NOT NULL,
`rarity` VARCHAR(8) NOT NULL, PRIMARY KEY (`id`));

CREATE TABLE IF NOT EXISTS `user_character` (
`id` INT NOT NULL AUTO_INCREMENT,
`user_id` INT NOT NULL,
`character_id` INT NOT NULL,
`acquired_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (`id`),
FOREIGN KEY (`character_id`) REFERENCES `game_character` (`id`));

INSERT INTO `user` (`name`, `token`)
VALUES
  ('user1', 'aaaaa'),
  ('user2', 'bbbbb'),
  ('user3', 'ccccc'),
  ('user4', 'ddddd'),
  ('user5', 'eeeee');

INSERT INTO `game_character` (`name`, `rarity`)
VALUES
  ('Character A1', 'N'),
  ('Character A2', 'N'),
  ('Character A3', 'N'),
  ('Character A4', 'N'),
  ('Character A5', 'N'),
  ('Character A6', 'N'),
  ('Character B1', 'R'),
  ('Character B2', 'R'),
  ('Character C1', 'R'),
  ('Character C2', 'R'),
  ('Character D', 'SR'),
  ('Character E', 'SSR');