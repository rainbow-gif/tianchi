CREATE TABLE IF NOT EXISTS `room`(
                                     `id` INT UNSIGNED AUTO_INCREMENT,
                                     `name` VARCHAR(100) UNIQUE NOT NULL,
                                     PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE IF NOT EXISTS `user`(
                                     `id` INT UNSIGNED AUTO_INCREMENT,
                                     `username` VARCHAR(100) UNIQUE NOT NULL,
                                     `fristname` VARCHAR(100) NOT NULL,
                                     `lastname` VARCHAR(100) NOT NULL,
                                     `email` VARCHAR(100) UNIQUE NOT NULL,
                                     `password` VARCHAR(100) NOT NULL,
                                     `phone` VARCHAR(20) UNIQUE NOT NULL,
                                     `roomid` INT UNSIGNED,
                                     PRIMARY KEY ( `id` ),
                                     FOREIGN KEY (`roomid`) REFERENCES `room` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE IF NOT EXISTS `time`(
                                     `id` VARCHAR(100) NOT NULL,
                                     `text` VARCHAR(100) NOT NULL,
                                     `time` INT UNSIGNED
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
