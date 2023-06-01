-- Clear db

DROP TABLE IF EXISTS `Users`;
DROP TABLE IF EXISTS `Timers`;
DROP TABLE IF EXISTS `EditAccess`;
DROP TABLE IF EXISTS `ReadAccess`;


-- Define schema

CREATE TABLE `Users` (
    `userID` text NOT NULL

    PRIMARY KEY (`userID`)
)

CREATE TABLE `Timers` (
    `id` uuid NOT NULL,
    `name` text,
    `initialSeconds` int,
    `remainingSeconds` int,
    `paused` bool,
    `userID` text,

    CONSTRAINT fk_user
        FOREIGN KEY(userID)
            REFERENCES Users(userID)
            ON DELETE CASCADE

    PRIMARY KEY (`id`)
);

CREATE TABLE `EditAccess` (
    `userID` text NOT NULL,
    `timerID` uuid NOT NULL

    CONSTRAINT fk_user
        FOREIGN KEY(userID)
            REFERENCES Users(userID)
            ON DELETE CASCADE

    CONSTRAINT fk_timer
        FOREIGN KEY(timerID)
            REFERENCES Timers(timerID)
            ON DELETE CASCADE
);

CREATE TABLE `ReadAccess` (
    `userID` text NOT NULL,
    `timerID` uuid NOT NULL

    CONSTRAINT fk_user
        FOREIGN KEY(userID)
            REFERENCES Users(userID)
            ON DELETE CASCADE

    CONSTRAINT fk_timer
        FOREIGN KEY(timerID)
            REFERENCES Timers(timerID)
            ON DELETE CASCADE
);