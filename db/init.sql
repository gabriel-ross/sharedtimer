-- Clear db

DROP TABLE IF EXISTS `Timers`;


-- Define schema

CREATE TABLE `Timers` (
    `id` uuid NOT NULL,
    `name` text
    `initialSeconds` int
    `remainingSeconds` int
    `paused` bool

    PRIMARY KEY (`id`)
);