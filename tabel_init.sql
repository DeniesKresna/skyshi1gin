SET time_zone = "+07:00";
CREATE TABLE `orders` (
  `id` bigint NOT NULL,
  `payment_id` bigint NOT NULL,
  `product_id` bigint NOT NULL,
  `amount` bigint DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `orders` (`id`, `payment_id`, `product_id`, `amount`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 2, 1, 2, '2024-04-04 17:23:28', '2024-04-04 17:23:28', NULL);

CREATE TABLE `payments` (
  `id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `price` bigint DEFAULT NULL,
  `code` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `paid_off` tinyint NOT NULL DEFAULT '0',
  `channel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `fail_reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `expired_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `payments` (`id`, `user_id`, `price`, `code`, `paid_off`, `channel`, `fail_reason`, `expired_at`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 0, 0, 'QSMKAqcbkBinNwwECFvY', 0, '', '', '2024-04-04 18:06:33', '2024-04-04 17:06:33', '2024-04-04 17:06:33', NULL),
(2, 1, 0, 'fGielipYbWOVdyBKKQof', 0, '', '', '2024-04-04 18:23:28', '2024-04-04 17:23:28', '2024-04-04 17:23:28', NULL);

CREATE TABLE `products` (
  `id` bigint NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `price` bigint NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


INSERT INTO `products` (`id`, `name`, `code`, `price`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Sabun Giv', 'SOAP-GIV', 0, '2024-04-03 12:02:23', '2024-04-03 12:02:23', NULL),
(2, 'Sabun Lux', 'SOAP-LUX', 5000, '2024-04-04 08:33:03', '2024-04-04 08:33:03', NULL);


CREATE TABLE `roles` (
  `id` bigint NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


INSERT INTO `roles` (`id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'administrator', '2023-11-22 06:02:08', '2023-11-22 06:02:08', NULL),
(2, 'user', '2023-11-22 06:02:08', '2023-11-22 06:02:08', NULL);


CREATE TABLE `users` (
  `id` bigint NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `role_id` bigint NOT NULL DEFAULT '2',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `users` (`id`, `name`, `email`, `phone`, `password`, `role_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Admin Denies', 'denieskresna@gmail.com', '081357006008', '$2a$10$MIvH94BmhUsq1b4c4IvvgOntx9EuwfcXqC2XodFT.w6KSiKNF/9xu', 1, '2024-04-03 11:11:26', '2024-04-03 11:11:26', NULL);


CREATE TABLE `warehouses` (
  `id` bigint NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `active` tinyint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `warehouses` (`id`, `name`, `active`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Warehouse A', 1, '2024-04-04 07:39:19', '2024-04-04 16:58:53', NULL),
(2, 'Warehouse B', 1, '2024-04-04 13:58:23', '2024-04-04 13:58:23', NULL);

CREATE TABLE `warehouse_product` (
  `id` bigint NOT NULL,
  `warehouse_id` bigint NOT NULL,
  `product_id` bigint NOT NULL,
  `amount` bigint DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


INSERT INTO `warehouse_product` (`id`, `warehouse_id`, `product_id`, `amount`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 1, 1, 2, '2024-04-04 09:57:02', '2024-04-04 14:00:46', NULL),
(2, 1, 2, 6, '2024-04-04 09:57:02', '2024-04-04 12:27:46', NULL),
(4, 2, 1, 2, '2024-04-04 14:00:46', '2024-04-04 14:00:46', NULL);

ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `payments`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`);

ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `role_id` (`role_id`);

ALTER TABLE `warehouses`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `warehouse_product`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `orders`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

ALTER TABLE `payments`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

ALTER TABLE `products`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

ALTER TABLE `roles`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

ALTER TABLE `users`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

ALTER TABLE `warehouses`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

ALTER TABLE `warehouse_product`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

