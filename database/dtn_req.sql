/*
 Navicat Premium Data Transfer

 Source Server         : My Localhost
 Source Server Type    : MySQL
 Source Server Version : 100428 (10.4.28-MariaDB)
 Source Host           : localhost:3306
 Source Schema         : dtn_req

 Target Server Type    : MySQL
 Target Server Version : 100428 (10.4.28-MariaDB)
 File Encoding         : 65001

 Date: 31/10/2024 17:05:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for departments
-- ----------------------------
DROP TABLE IF EXISTS `departments`;
CREATE TABLE `departments` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `code` varchar(10) NOT NULL,
  `level` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`),
  CONSTRAINT `departments_ibfk_1` FOREIGN KEY (`parent_id`) REFERENCES `departments` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of departments
-- ----------------------------
BEGIN;
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (1, NULL, 'Board Of Director', 'BOD0001', 1);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (2, 1, 'Information of Technology', 'DVS00001', 2);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (3, 1, 'Marketing And Sales', 'DVS00003', 2);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (4, 1, 'Purchasing', 'PURC', 2);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (5, 1, 'Finance', 'DVS00015', 2);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (6, 1, 'Special Project', 'SPRTJ0001', 2);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (7, 2, 'ERP Development', 'ITSDERP', 3);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (8, 2, 'Tech Development', 'DVS00005', 3);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (9, 2, ' Software Maintenance', 'DVS00008', 3);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (10, 2, 'Quality Assurance', 'DVS00012', 3);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (11, 2, 'IT Support', 'DVS00007', 3);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (12, 2, 'Implementation - AT', 'IMPOVS', 3);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (13, 2, 'Human Resource', 'DVS00014', 3);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (14, 2, 'HR Development', 'ITQA', 3);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (15, 2, 'General Affairs', 'DVS00013', 3);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (16, 2, 'ERP Implementation', 'ITSFERP', 3);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (17, 2, 'Training', 'DVS00016', 3);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (18, 3, 'Sales', 'DVS00017', 3);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (19, 6, 'PM Jakarta', 'PMJ0001', 3);
INSERT INTO `departments` (`id`, `parent_id`, `name`, `code`, `level`) VALUES (20, 6, 'Dekorey', 'DKR0001', 3);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
