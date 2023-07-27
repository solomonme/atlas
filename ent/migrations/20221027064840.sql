-- modify "organizations" table
ALTER TABLE `organizations` ADD COLUMN `security_score` double NOT NULL DEFAULT 50;
-- modify "audits" table
ALTER TABLE `audits` ADD COLUMN `org_id` bigint NOT NULL, ADD CONSTRAINT `audits_organizations_audits` FOREIGN KEY (`org_id`) REFERENCES `organizations` (`id`) ON DELETE RESTRICT;
