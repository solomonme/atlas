-- create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `email` varchar(255) NOT NULL, `balance` double NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `name` (`name`), UNIQUE INDEX `email` (`email`)) CHARSET utf8mb4 COLLATE utf8mb4_bin AUTO_INCREMENT 4294967296;
-- create "ent_types" table
CREATE TABLE `ent_types` (`id` bigint unsigned NOT NULL AUTO_INCREMENT, `type` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `type` (`type`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "audits" table
CREATE TABLE `audits` (`id` bigint NOT NULL AUTO_INCREMENT, `identity` varchar(255) NOT NULL, `timestamp` timestamp NOT NULL, `balance` double NOT NULL, `description` varchar(255) NOT NULL, `user_audits` bigint NULL, PRIMARY KEY (`id`), CONSTRAINT `audits_users_audits` FOREIGN KEY (`user_audits`) REFERENCES `users` (`id`) ON DELETE SET NULL) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- add pk ranges for ('audits'),('users') tables
INSERT INTO `ent_types` (`type`) VALUES ('audits'), ('users');
