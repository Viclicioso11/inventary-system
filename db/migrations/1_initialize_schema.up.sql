CREATE TABLE day_of_week (
    id          SERIAL PRIMARY KEY,
    description VARCHAR(50) NOT NULL
);

CREATE TABLE catalog (
    id          SERIAL PRIMARY KEY,
    description VARCHAR(100) NOT NULL,
    code        VARCHAR(50) NOT NULL,
    enabled     BOOLEAN NOT NULL,
    created_at  TIMESTAMP NOT NULL
);

CREATE TABLE catalog_detail (
    id          SERIAL PRIMARY KEY,
    description VARCHAR(100) NOT NULL,
    code        VARCHAR(50) NOT NULL,
    enabled     BOOLEAN NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    catalog_id  INT NOT NULL,
    CONSTRAINT fk_catalog
      FOREIGN KEY(catalog_id) 
	  REFERENCES catalog(id)
);

CREATE TABLE store (
    id          VARCHAR(100) PRIMARY KEY,
    name        VARCHAR(150) NOT NULL,
    code        VARCHAR(50) NOT NULL,
    enabled     BOOLEAN NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    image_key   VARCHAR(255)
);

CREATE TABLE schedule (
    id              SERIAL PRIMARY KEY,
    day_of_week_id  INT,
    store_id        VARCHAR(100),
    open_at         TIME,
    close_at        TIME,
    CONSTRAINT fk_day_of_week
      FOREIGN KEY(day_of_week_id) 
	  REFERENCES day_of_week(id),
    CONSTRAINT fk_store
      FOREIGN KEY(store_id) 
	  REFERENCES store(id)
);