
SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
                            `id_account` bigint NOT NULL AUTO_INCREMENT,
                            `name` varchar(255) NOT NULL,
                            `email` varchar(255) NOT NULL,
                            `password` varchar(255) NOT NULL,
                            `status` varchar(255) NOT NULL,
                            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `contact_number` varchar(20) DEFAULT NULL,
                            `profile_image` blob,
                            PRIMARY KEY (`id_account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


DROP TABLE IF EXISTS `goals`;
CREATE TABLE `goals` (
                         `id_goal` bigint NOT NULL AUTO_INCREMENT,
                         `id_account` bigint NOT NULL,
                         `name` varchar(255) NOT NULL,
                         `icon_name` varchar(50) DEFAULT NULL,
                         `current_progress` decimal(10,2) NOT NULL,
                         `goal_number` decimal(10,2) NOT NULL,
                         `expected_date` datetime NOT NULL,
                         `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id_goal`),
                         KEY `id_account` (`id_account`),
                         CONSTRAINT `goals_ibfk_1` FOREIGN KEY (`id_account`) REFERENCES `accounts` (`id_account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions` (
                                `id_transaction` bigint NOT NULL AUTO_INCREMENT,
                                `id_origin_account` bigint DEFAULT NULL,
                                `transaction_type` varchar(255) NOT NULL,
                                `amount` decimal(10,2) NOT NULL,
                                `description` varchar(255) DEFAULT NULL,
                                `category` varchar(255) DEFAULT NULL,
                                `id_destination_account` bigint DEFAULT NULL,
                                PRIMARY KEY (`id_transaction`),
                                KEY `id_destination_account` (`id_destination_account`),
                                KEY `id_origin_account` (`id_origin_account`),
                                CONSTRAINT `transactions_ibfk_4` FOREIGN KEY (`id_destination_account`) REFERENCES `accounts` (`id_account`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                                CONSTRAINT `transactions_ibfk_5` FOREIGN KEY (`id_origin_account`) REFERENCES `accounts` (`id_account`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


DROP TABLE IF EXISTS `wallet`;
CREATE TABLE `wallet` (
                          `id_wallet` bigint NOT NULL AUTO_INCREMENT,
                          `id_account` bigint NOT NULL,
                          `current_balance` decimal(10,2) DEFAULT NULL,
                          PRIMARY KEY (`id_wallet`),
                          KEY `id_account` (`id_account`),
                          CONSTRAINT `wallet_ibfk_1` FOREIGN KEY (`id_account`) REFERENCES `accounts` (`id_account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

