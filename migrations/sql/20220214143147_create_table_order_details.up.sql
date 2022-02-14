CREATE TABLE IF NOT EXISTS `order_details` (
  `order_id` bigint(20) unsigned NOT NULL,
  `inventory_id` bigint(20) unsigned NOT NULL,
  `qty` decimal(10,2) unsigned NOT NULL,
  `wac` decimal(10,2) unsigned NOT NULL,
  `amount` decimal(10,2) unsigned NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`order_id`, `inventory_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
