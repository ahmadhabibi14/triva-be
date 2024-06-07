CREATE TABLE Users (
  id          VARCHAR(36)   PRIMARY KEY,
  username    VARCHAR(100)  UNIQUE NOT NULL,
  fullName    VARCHAR(255)  NOT NULL,
  email       VARCHAR(100)  UNIQUE,
  password    VARCHAR(100)  NOT NULL,
  avatar      VARCHAR(255)  NOT NULL DEFAULT '/img/avatars/default.png',
  googleId    VARCHAR(255)  DEFAULT '',
  facebookId  VARCHAR(255)  DEFAULT '',
  githubId    VARCHAR(255)  DEFAULT '',
  createdAt   TIMESTAMP     WITHOUT TIME ZONE DEFAULT now(),
  updatedAt   TIMESTAMP     WITHOUT TIME ZONE DEFAULT now(),
  deletedAt   TIMESTAMP     WITHOUT TIME ZONE
);