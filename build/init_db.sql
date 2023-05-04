CREATE DATABASE IF NOT EXISTS `interview` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;
use `interview`;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`               int                                     NOT NULL AUTO_INCREMENT,
    `username`         varchar(180) COLLATE utf8mb4_unicode_ci NOT NULL,
    `password`         varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `first_name`       varchar(255) COLLATE utf8mb4_unicode_ci          DEFAULT NULL,
    `last_name`        varchar(255) COLLATE utf8mb4_unicode_ci          DEFAULT NULL,
    `email`            varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `phone_no`         varchar(11)  COLLATE utf8mb4_unicode_ci NOT NULL,
    `roles`            JSON,
    `verification_ode` varchar(6) COLLATE utf8mb4_unicode_ci            DEFAULT NULL,
    `is_verified`      BOOLEAN                                 NOT NULL DEFAULT false,
    `is_active`        BOOLEAN                                 NOT NULL DEFAULT false,
    `created_at`       datetime                                         DEFAULT CURRENT_TIMESTAMP,
    `updated_at`       datetime                                         DEFAULT NULL,
    `last_login`       datetime                                         DEFAULT NULL,
    `deleted_at`       datetime                                         DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `UNIQ_8D93D64992FC23A8` (`username`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

INSERT INTO `user`
VALUES (1,
        'user1',
        '$2a$13$idXK61n71KvV26xsXOt8F.zHlQjquyfLdVHPlBMDeb.3DdpL.aRDO',
        'username1',
        'lastname1',
        'username1@domain1.com',
        '123456789'
        '{
          "roles": "ROLE_ADMIN"
        }',
        '624813',
        1,
        1,
        '2023-05-01 09:16:12',
        '2023-05-02 22:14:09',
        NULL,
        NULL),
       (2,
        'another_user',
        '$2a$13$idXK61n71KvV26xsXOt8F.zHlQjquyfLdVHPlBMDeb.3DdpL.aRDO',
        'another_user_firstname',
        'another_lastname',
        'user@user.com',
        '987654321'
        '{
          "roles": "ROLE_USER"
        }',
        '176278',
        1,
        1,
        '2023-05-01 22:14:09',
        NULL,
        '2023-05-02 22:14:09',
        NULL),
       (3,
        'bad_user',
        '$2a$13$TQWml9TszLLVuZ62rMjbxeXSkJWeiYN51tAp4Dnb3zq6Q/c4OM196',
        'bad, bad,',
        'very bad user',
        'trickortreat@strange.com',
        '123333777',
        '{
          "roles": "ROLE_USER"
        }',
        '347816',
        1,
        0,
        '2023-05-01 04:33:20',
        '2023-05-01 19:11:34',
        '2023-05-02 15:47:04',
        '2023-05-02 17:12:48'),
       (5,
        'vip',
        '$2a$13$TQWml9TszLLVuZ62rMjbxeXSkJWeiYN51tAp4Dnb3zq6Q/c4OM196',
        'very important',
        'user',
        'memberplus@vip.com',
        '777333444',
        '{
          "roles": "ROLE_VIP"
        }',
        '642103',
        1,
        1,
        '2023-05-01 04:33:20',
        NULL,
        '2023-05-02 15:47:04',
        NULL);