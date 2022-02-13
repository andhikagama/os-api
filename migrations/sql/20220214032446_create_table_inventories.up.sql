CREATE TABLE IF NOT EXISTS `inventories` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `sku_name` varchar(255) NOT NULL DEFAULT '',
  `sku_code` varchar(255) NOT NULL DEFAULT '',
  `wac` decimal(10,2) NOT NULL,
  `total_qty` decimal(10,2) NOT NULL,
  `available_qty` decimal(10,2) NOT NULL,
  `quarantine_qty` decimal(10,2) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_sku_code` (`sku_code`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;