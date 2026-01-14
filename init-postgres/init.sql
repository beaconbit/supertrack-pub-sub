CREATE TABLE general (
    id SERIAL PRIMARY KEY,
    timestamp BIGINT NOT NULL,
    source_mac VARCHAR NOT NULL,
    source_ip VARCHAR NOT NULL,
    value INTEGER DEFAULT 1,
    factor REAL,
    reading REAL,
    measurement REAL,
    metric INTEGER,
    data_field_index INTEGER
);

CREATE TABLE livestatuscbw1 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    delta_seconds INTEGER,
    cbw INTEGER,
    pocket INTEGER,
    product INTEGER,
);
CREATE TABLE livestatuscbw2 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    delta_seconds INTEGER,
    cbw INTEGER,
    pocket INTEGER,
    product INTEGER,
);
CREATE TABLE livestatuscbw3 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    delta_seconds INTEGER,
    cbw INTEGER,
    pocket INTEGER,
    product INTEGER,
);

CREATE TABLE cbw1pocket1 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket2 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket3 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket4 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket5 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket6 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket7 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket8 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket9 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket10 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket11 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket12 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket13 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket14 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket15 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket16 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket17 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket18 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket19 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket20 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket21 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket22 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket23 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket24 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket25 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket26 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket27 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket28 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket29 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket30 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket31 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw1pocket32 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);

CREATE TABLE cbw2pocket1 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket2 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket3 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket4 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket5 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket6 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket7 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket8 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket9 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket10 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket11 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket12 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket13 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket14 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket15 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket16 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket17 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket18 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket19 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket20 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket21 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket22 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket23 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket24 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket25 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket26 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket27 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket28 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket29 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket30 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket31 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw2pocket32 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);

CREATE TABLE cbw3pocket1 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket2 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket3 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket4 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket5 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket6 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket7 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket8 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket9 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket10 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket11 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket12 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket13 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket14 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket15 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket16 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket17 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket18 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket19 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket20 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket21 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket22 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket23 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket24 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket25 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket26 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket27 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket28 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket29 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket30 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket31 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
CREATE TABLE cbw3pocket32 (
    id uuid PRIMARY KEY,
    entered BIGINT NOT NULL,
    exited BIGINT NOT NULL,
    delta_seconds INTEGER
    product INTEGER,
);
