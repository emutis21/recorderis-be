CREATE TABLE `refresh_tokens` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `token_id` uuid UNIQUE NOT NULL DEFAULT 'uuid_generate_v4()',
  `user_id` uuid NOT NULL,
  `token_hash` varchar(255) NOT NULL,
  `device_type` ENUM ('web', 'mobile', 'tablet') NOT NULL,
  `device_name` varchar(255),
  `is_remember_me` boolean NOT NULL DEFAULT false,
  `expires_at` timestamp NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT 'now()',
  `last_used_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `active_sessions` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `session_id` uuid UNIQUE NOT NULL DEFAULT 'uuid_generate_v4()',
  `user_id` uuid NOT NULL,
  `refresh_token_id` uuid NOT NULL,
  `ip_address` varchar(45),
  `user_agent` varchar(255),
  `last_activity` timestamp NOT NULL DEFAULT 'now()',
  `created_at` timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE `users` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `user_id` uuid UNIQUE NOT NULL DEFAULT 'uuid_generate_v4()',
  `username` varchar(255) NOT NULL,
  `display_name` varchar(255) NOT NULL,
  `avatar_url` varchar(255),
  `email` varchar(255) UNIQUE NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `role` ENUM ('admin', 'user') NOT NULL DEFAULT 'user',
  `created_at` timestamp NOT NULL DEFAULT 'now()',
  `deleted_at` timestamp,
  `updated_at` timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE `memories` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `memory_id` uuid UNIQUE NOT NULL DEFAULT 'uuid_generate_v4()',
  `user_id` uuid NOT NULL,
  `title` varchar(255) NOT NULL,
  `date` timestamp NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT 'now()',
  `deleted_at` timestamp,
  `is_public` boolean NOT NULL DEFAULT false
);

CREATE TABLE `descriptions` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `description_id` uuid UNIQUE NOT NULL DEFAULT 'uuid_generate_v4()',
  `memory_id` uuid NOT NULL,
  `text` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT 'now()',
  `deleted_at` timestamp,
  `version` integer NOT NULL DEFAULT 1
);

CREATE TABLE `photos` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `photo_id` uuid UNIQUE NOT NULL DEFAULT 'uuid_generate_v4()',
  `memory_id` uuid NOT NULL,
  `filename` varchar(255) NOT NULL,
  `url` varchar(255) NOT NULL,
  `display_order` integer NOT NULL DEFAULT 0,
  `uploaded_at` timestamp NOT NULL DEFAULT 'now()',
  `deleted_at` timestamp
);

CREATE TABLE `tags` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `tag_id` uuid UNIQUE NOT NULL DEFAULT 'uuid_generate_v4()',
  `name` varchar(255) UNIQUE NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT 'now()',
  `deleted_at` timestamp
);

CREATE TABLE `memory_tags` (
  `memory_id` uuid NOT NULL,
  `tag_id` uuid NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT 'now()',
  `deleted_at` timestamp
);

CREATE TABLE `locations` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `location_id` uuid UNIQUE NOT NULL DEFAULT 'uuid_generate_v4()',
  `location` varchar(255) NOT NULL,
  `longitude` decimal(10,7) NOT NULL,
  `latitude` decimal(10,7) NOT NULL,
  `city` varchar(255) NOT NULL,
  `country` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT 'now()',
  `deleted_at` timestamp
);

CREATE TABLE `memory_locations` (
  `memory_id` uuid NOT NULL,
  `location_id` uuid NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT 'now()',
  `deleted_at` timestamp
);

CREATE TABLE `memory_shares` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `memory_id` uuid NOT NULL,
  `shared_with_user_id` uuid NOT NULL,
  `permission_level` ENUM ('read', 'write', 'admin') NOT NULL DEFAULT 'read',
  `created_at` timestamp NOT NULL DEFAULT 'now()',
  `deleted_at` timestamp
);

CREATE UNIQUE INDEX `refresh_tokens_index_0` ON `refresh_tokens` (`user_id`, `token_hash`);

CREATE UNIQUE INDEX `active_sessions_index_1` ON `active_sessions` (`user_id`, `refresh_token_id`);

CREATE UNIQUE INDEX `memory_tags_index_2` ON `memory_tags` (`memory_id`, `tag_id`);

CREATE UNIQUE INDEX `memory_locations_index_3` ON `memory_locations` (`memory_id`, `location_id`);

CREATE UNIQUE INDEX `memory_shares_index_4` ON `memory_shares` (`memory_id`, `shared_with_user_id`);

ALTER TABLE `refresh_tokens` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`);

ALTER TABLE `active_sessions` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`);

ALTER TABLE `active_sessions` ADD FOREIGN KEY (`refresh_token_id`) REFERENCES `refresh_tokens` (`token_id`);

ALTER TABLE `memories` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`);

ALTER TABLE `descriptions` ADD FOREIGN KEY (`memory_id`) REFERENCES `memories` (`memory_id`);

ALTER TABLE `photos` ADD FOREIGN KEY (`memory_id`) REFERENCES `memories` (`memory_id`);

ALTER TABLE `memory_tags` ADD FOREIGN KEY (`memory_id`) REFERENCES `memories` (`memory_id`);

ALTER TABLE `memory_tags` ADD FOREIGN KEY (`tag_id`) REFERENCES `tags` (`tag_id`);

ALTER TABLE `memory_locations` ADD FOREIGN KEY (`memory_id`) REFERENCES `memories` (`memory_id`);

ALTER TABLE `memory_locations` ADD FOREIGN KEY (`location_id`) REFERENCES `locations` (`location_id`);

ALTER TABLE `memory_shares` ADD FOREIGN KEY (`memory_id`) REFERENCES `memories` (`memory_id`);

ALTER TABLE `memory_shares` ADD FOREIGN KEY (`shared_with_user_id`) REFERENCES `users` (`user_id`);
