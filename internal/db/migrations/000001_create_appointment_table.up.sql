
CREATE TABLE appointments (
    id serial PRIMARY KEY,
    start_time timestamptz NOT NULL,
    end_time timestamptz NOT NULL
);


CREATE INDEX ON "appointments" ("id");