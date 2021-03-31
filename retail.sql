gen --connstr "root@tcp(127.0.0.1:3306)/retail_store_management?&parseTime=True" --database retail_store_management --json --gorm --guregu --rest
CREATE TABLE `user` (
`id` bigint(11) NOT NULL AUTO_INCREMENT,
`real_name` varchar(255) NOT NULL,
`user_name` varchar(255) NOT NULL,
`password` varchar(255) NOT NULL,
`shop_id` bigint(11) NOT NULL,
`token` varchar(255) NULL,
`phone` varchar(255) NOT NULL,
`create_time` datetime NULL,
`updata_time` datetime NULL ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`) 
);
CREATE TABLE `shop` (
`id` bigint(11) NOT NULL AUTO_INCREMENT,
`name` varchar(255) NOT NULL,
`address` varchar(255) NOT NULL,
`owner` varchar(255) NOT NULL,
`create_time` datetime NULL,
`update_time` datetime NULL ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`) 
);
CREATE TABLE `rule` (
`id` bigint(11) NOT NULL AUTO_INCREMENT,
`v0` varchar(255) NOT NULL,
`v1` varchar(255) NOT NULL,
`create_time` datetime NULL,
`update_time` datetime NULL ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`) 
);
CREATE TABLE `role` (
`id` bigint(11) NOT NULL AUTO_INCREMENT,
`name` varchar(255) NOT NULL,
`shop_id` bigint(11) NOT NULL,
`is_ customer` int(255) NULL,
`create_time` datetime NULL,
`update_time` datetime NULL ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`) 
);
CREATE TABLE `permission` (
`id` bigint(11) NOT NULL AUTO_INCREMENT,
`name` varchar(255) NOT NULL,
`url` varchar(255) NOT NULL,
`create_time` datetime NULL,
`update_time` datetime NULL ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`) 
);
CREATE TABLE `spu` (
`id` bigint(11) NOT NULL AUTO_INCREMENT,
`name` varchar(255) NOT NULL,
`shop_id` bigint(11) NOT NULL,
`img_url` varchar(255) NULL,
`category_id` bigint(11) NOT NULL,
`original_price` decimal(10,2) NOT NULL,
`member_price` decimal(10,2) NULL,
`spec_values` varchar(255) NOT NULL,
`create_time` datetime NULL,
`update_time` datetime NULL ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`) 
);
CREATE TABLE `sku` (
`id` bigint(11) NOT NULL AUTO_INCREMENT,
`name` varchar(255) NOT NULL,
`original_price` decimal(10,2) NOT NULL,
`member_price` decimal(10,2) NOT NULL,
`shop_id` bigint(11) NOT NULL,
`category_id` bigint(11) NOT NULL,
`img_url` varchar(255) NULL,
`barcode` varchar(255) NOT NULL,
`create_time` datetime NULL,
`update_time` datetime NULL ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`) 
);
CREATE TABLE `category` (
`id` bigint(11) NOT NULL AUTO_INCREMENT,
`name` varchar(255) NOT NULL,
`parent_id` bigint(11) NOT NULL,
`shop_id` bigint(11) NOT NULL,
`level` varchar(255) NOT NULL,
`create_time` datetime NULL,
`update_time` datetime NULL ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`) 
);
CREATE TABLE `stock` (
`id` bigint(11) NOT NULL AUTO_INCREMENT,
`shop_id` bigint(11) NOT NULL,
`num` bigint NULL,
`sale_num` bigint NULL,
`add_up_num` bigint NULL,
`barcode` varchar(255) NOT NULL,
`sku_id` bigint(11) NOT NULL,
`create_time` datetime NULL,
`update_time` datetime NULL ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`) 
);
CREATE TABLE `order` (
`id` bigint(11) NOT NULL AUTO_INCREMENT,
`shop_id` bigint(11) NOT NULL,
`opreator` bigint(255) NOT NULL,
`count` decimal NOT NULL,
`amount` int(11) NOT NULL,
`create_time` datetime NULL,
`update_time` datetime NULL ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`) 
);
CREATE TABLE `order_sku` (
`id` bigint(11) NOT NULL AUTO_INCREMENT,
`sku_id` bigint(11) NOT NULL,
`order_id` bigint(11) NOT NULL,
`sale_price` decimal(10,2) NOT NULL,
`create_time` datetime NULL,
`update_time` datetime NULL,
PRIMARY KEY (`id`) 
);
CREATE TABLE `table_1` (
);
