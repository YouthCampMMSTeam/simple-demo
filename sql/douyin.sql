CREATE TABLE `User` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) UNIQUE NOT NULL,
  `password` varchar(255) NOT NULL,
  `follow_count` integer NOT NULL,
  `follower_count` integer NOT NULL
);

CREATE TABLE `Video` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime,
  `play_url` varchar(255) NOT NULL,
  `cover_url` varchar(255) NOT NULL,
  `favorite_count` integer NOT NULL,
  `comment_count` integer NOT NULL,
  `title` varchar(255) NOT NULL,
  `author_id` integer NOT NULL
);

CREATE TABLE `Favorite` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime,
  `video_id` integer NOT NULL,
  `user_id` integer NOT NULL
);

CREATE TABLE `Relation` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `follow_id` integer NOT NULL,
  `follower_id` integer NOT NULL
);

CREATE TABLE `Comment` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime,
  `video_id` integer NOT NULL,
  `user_id` integer NOT NULL,
  `content` varchar(255)
);

CREATE TABLE `Message` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `from_id` integer NOT NULL,
  `to_id` integer NOT NULL,
  `content` varchar(255),
  `create_time` datetime
);

CREATE INDEX `User_index_0` ON `User` (`name`);

CREATE INDEX `Video_index_1` ON `Video` (`author_id`);

CREATE INDEX `Favorite_index_2` ON `Favorite` (`video_id`);

CREATE INDEX `Favorite_index_3` ON `Favorite` (`user_id`);

CREATE INDEX `Relation_index_4` ON `Relation` (`follow_id`);

CREATE INDEX `Relation_index_5` ON `Relation` (`follower_id`);

CREATE INDEX `Comment_index_6` ON `Comment` (`video_id`);

CREATE INDEX `Comment_index_7` ON `Comment` (`user_id`);

CREATE INDEX `Message_index_8` ON `Message` (`from_id`);

CREATE INDEX `Message_index_9` ON `Message` (`to_id`);

ALTER TABLE `Video` ADD FOREIGN KEY (`author_id`) REFERENCES `User` (`id`);

ALTER TABLE `Favorite` ADD FOREIGN KEY (`video_id`) REFERENCES `Video` (`id`);

ALTER TABLE `Favorite` ADD FOREIGN KEY (`user_id`) REFERENCES `User` (`id`);

ALTER TABLE `Relation` ADD FOREIGN KEY (`follow_id`) REFERENCES `User` (`id`);

ALTER TABLE `Relation` ADD FOREIGN KEY (`follower_id`) REFERENCES `User` (`id`);

ALTER TABLE `Comment` ADD FOREIGN KEY (`video_id`) REFERENCES `Video` (`id`);

ALTER TABLE `Comment` ADD FOREIGN KEY (`user_id`) REFERENCES `User` (`id`);

ALTER TABLE `Message` ADD FOREIGN KEY (`from_id`) REFERENCES `User` (`id`);

ALTER TABLE `Message` ADD FOREIGN KEY (`to_id`) REFERENCES `User` (`id`);
