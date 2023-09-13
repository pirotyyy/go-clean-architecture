USE ca_db;

CREATE TABLE IF NOT EXISTS `user` (
`user_id` VARCHAR(64) NOT NULL,
`token` VARCHAR(255) NOT NULL,
`name` VARCHAR(32) NOT NULL,
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`user_id`));

DELIMITER //
CREATE TRIGGER `user_before_insert` BEFORE INSERT ON `user`
FOR EACH ROW
BEGIN
  SET NEW.`user_id` = CONCAT('user-', UUID());
END;
//
DELIMITER ;

CREATE TABLE IF NOT EXISTS `game_character` (
`character_id` VARCHAR(64) NOT NULL,
`name` VARCHAR(255) NOT NULL,
`rarity` VARCHAR(8) NOT NULL, PRIMARY KEY (`character_id`));

DELIMITER //
CREATE TRIGGER `game_character_before_insert` BEFORE INSERT ON `game_character`
FOR EACH ROW
BEGIN
  SET NEW.`character_id` = CONCAT('character-', UUID());
END;
//
DELIMITER ;


CREATE TABLE IF NOT EXISTS `user_character` (
  `user_character_id` VARCHAR(64) PRIMARY KEY NOT NULL,
  `user_id` VARCHAR(64) NOT NULL,
  `character_id` VARCHAR(64) NOT NULL,
  `acquired_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`character_id`) REFERENCES `game_character` (`character_id`)
);

DELIMITER //
CREATE TRIGGER `user_character_before_insert` BEFORE INSERT ON `user_character`
FOR EACH ROW
BEGIN
  SET NEW.`user_character_id` = CONCAT('user-character-', UUID());
END;
//
DELIMITER ;



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