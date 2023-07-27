-- create "organizations" table
CREATE TABLE `organizations` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `revenues` double NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin AUTO_INCREMENT 8589934592;
-- modify "users" table
ALTER TABLE `users` ADD COLUMN `org_id` bigint NOT NULL, ADD CONSTRAINT `users_organizations_users` FOREIGN KEY (`org_id`) REFERENCES `organizations` (`id`) ON DELETE CASCADE;
-- add pk ranges for ('organizations') tables
INSERT INTO `ent_types` (`type`) VALUES ('organizations');
