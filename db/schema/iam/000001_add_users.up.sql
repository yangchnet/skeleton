CREATE TABLE `user` (
  `id` varchar(50) PRIMARY KEY,
  `username` varchar(50) UNIQUE NOT NULL,
  `email` varchar(50) UNIQUE NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

