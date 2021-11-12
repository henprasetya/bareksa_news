CREATE TABLE `news`.`topic` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(100) NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `news`.`news` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(100) NULL,
  `status` VARCHAR(10) NULL,
  `topic_id` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_topic_id_idx` (`topic_id` ASC) VISIBLE,
  CONSTRAINT `fk_topic_id`
    FOREIGN KEY (`topic_id`)
    REFERENCES `news`.`topic` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);

CREATE TABLE `news`.`tags` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(45) NULL,
  `description` VARCHAR(150) NULL,
  PRIMARY KEY (`id`),
INDEX `fk_news_id_idx` (`news_id` ASC) VISIBLE,
CONSTRAINT `fk_tags_news`
  FOREIGN KEY (`news_id`)
  REFERENCES `news`.`news` (`id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION;
);
