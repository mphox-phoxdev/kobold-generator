drop database if exists kobolds;
create database kobolds;
use kobolds;

CREATE TABLE skills(
    id  INT(2) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(16) not null,
    stat enum("brawn","ego","extraneous","reflexes") not null,
    description VARCHAR(500),
    dangerous BOOLEAN,
    every_kobold BOOLEAN,
    extra varchar(500) default ""
);

INSERT INTO skills (id, name, stat, description, dangerous, every_kobold) values
    (1, "bully", "brawn", "", false, false),
    (2, "duel", "brawn", "", true, false),
    (3, "heft", "brawn", "", true, false),
    (4, "lift", "brawn", "", false, false),
    (5, "sport", "brawn", "", false, false),
    (6, "swim", "brawn", "", false, false),
    (7, "wrassle", "brawn", "", false, false),
    (8, "fear", "ego", "", true, false),
    (9, "lacky", "ego", "", true, false),
    (10, "sage", "ego", "", true, false),
    (11, "shoot", "ego", "", false, false),
    (12, "speak_human", "ego", "", false, false),
    (13, "trap", "ego", "", false, false),
    (14, "tinker", "ego", "", true, false),
    (15, "bard", "extraneous", "", false, false),
    (16, "cook", "extraneous", "", false, true),
    (17, "dungoen", "extraneous", "", false, false),
    (18, "nature", "extraneous", "", false, false),
    (19, "perform", "extraneous", "", false, false),
    (20, "speak_critter", "extraneous", "", false, false),
    (21, "track", "extraneous", "", false, false),
    (22, "trade", "extraneous", "", true, false),
    (23, "fast", "reflexes", "", false, false),
    (24, "hide", "reflexes", "", false, false),
    (25, "nurture", "reflexes", "", false, false),
    (26, "ride", "reflexes", "", false, false),
    (27, "sneak", "reflexes", "", false, false),
    (28, "steal", "reflexes", "", false, false),
    (29, "wiggle", "reflexes", "", false, false),
    (30, "MAYORS CHOICE!!!", "extraneous", "", false, false);

CREATE TABLE roles(
    name VARCHAR(8) PRIMARY Key
);

INSERT INTO roles (name) values 
    ("blazer"),
    ("caster"),
    ("fryer"),
    ("scrapper"),
    ("taker"),
    ("weirder");

CREATE TABLE role_skills(
    d6 int(1) UNSIGNED not null,
    role VARCHAR(8) not null,
    brawn_skill int(2) UNSIGNED not null,
    ego_skill int(2) UNSIGNED not null,
    extraneous_skill int(2) UNSIGNED not null,
    reflex_skill int(2) UNSIGNED not null,

    FOREIGN KEY (role) references roles(name),
    FOREIGN KEY (brawn_skill) references skills(id),
    FOREIGN KEY (ego_skill) references skills(id),
    FOREIGN KEY (extraneous_skill) references skills(id),
    FOREIGN KEY (reflex_skill) references skills(id)
);

INSERT INTO role_skills (d6, role, brawn_skill, ego_skill, extraneous_skill, reflex_skill) values
    (1, "blazer", 3, 9, 17, 23),
    (2, "blazer", 4, 10, 18, 24),
    (3, "blazer", 5, 10, 18, 24),
    (4, "blazer", 5, 10, 20, 26),
    (5, "blazer", 6, 12, 21, 27),
    (6, "blazer", 7, 13, 21, 27),

    (1, "caster", 3, 9, 15, 23),
    (2, "caster", 4, 9, 17, 24),
    (3, "caster", 5, 9, 18, 25),
    (4, "caster", 5, 9, 20, 25),
    (5, "caster", 6, 10, 22, 28),
    (6, "caster", 7, 14, 22, 29),

    (1, "fryer", 1, 8, 15, 24),
    (2, "fryer", 3, 10, 17, 25),
    (3, "fryer", 4, 12, 18, 25),
    (4, "fryer", 4, 13, 18, 27),
    (5, "fryer", 6, 13, 20, 28),
    (6, "fryer", 7, 14, 22, 28),

    (1, "scrapper", 1, 8, 17, 23),
    (2, "scrapper", 2, 11, 17, 26),
    (3, "scrapper", 3, 11, 18, 26),
    (4, "scrapper", 3, 12, 20, 28),
    (5, "scrapper", 7, 13, 21, 29),
    (6, "scrapper", 7, 13, 21, 29),

    (1, "taker", 2, 8, 15, 24),
    (2, "taker", 4, 11, 17, 24),
    (3, "taker", 5, 12, 17, 27),
    (4, "taker", 5, 13, 21, 27),
    (5, "taker", 6, 13, 21, 28),
    (6, "taker", 6, 14, 22, 28);

CREATE TABLE random_skills(
    first_d6 int(1) UNSIGNED not null,
    second_d6 int(1) UNSIGNED not null,
    skill_id int(2) UNSIGNED not null,

    FOREIGN KEY (skill_id) references skills(id)
);

INSERT INTO random_skills (first_d6, second_d6, skill_id) values
    (1, 1, 15),
    (1, 2, 1),
    (1, 3, 16),
    (1, 4, 2),
    (1, 5, 17),
    (1, 6, 23),

    (2, 1, 8),
    (2, 2, 3),
    (2, 3, 24),
    (2, 4, 9),
    (2, 5, 4),
    (2, 6, 18),

    (3, 1, 18),
    (3, 2, 19),
    (3, 3, 26),
    (3, 4, 10),
    (3, 5, 11),
    (3, 6, 27),

    (4, 1, 12),
    (4, 2, 20),
    (4, 3, 5),
    (4, 4, 28),
    (4, 5, 6),
    (4, 6, 13),

    (5, 1, 22),
    (5, 2, 13),
    (5, 3, 14),
    (5, 4, 29),
    (5, 5, 7),
    (5, 6, 2),

    (6, 1, 8),
    (6, 2, 3),
    (6, 3, 9),
    (6, 4, 10),
    (6, 5, 22),
    (6, 6, 14);

