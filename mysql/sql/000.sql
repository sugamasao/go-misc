CREATE TABLE IF NOT EXISTS `blog` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `published_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

CREATE TABLE IF NOT EXISTS `entry`(
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `blog_id` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `body` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

INSERT INTO `blog` (id, title, created_at, updated_at) VALUES
  (1, 'sample 1', NOW(), NOW()),
  (2, 'sample 2', NOW(), NOW()),
  (3, 'sample 3', NOW(), NOW())
;

INSERT INTO `entry` (id, blog_id, title, body, created_at, updated_at) VALUES
  (1, 1, 'sample 1-1', 'foo', NOW(), NOW()),
  (2, 2, 'sample 2-1', 'foo', NOW(), NOW()),
  (3, 2, 'sample 2-2', 'foo', NOW(), NOW()),
  (4, 2, 'sample 2-3', 'foo', NOW(), NOW())
;