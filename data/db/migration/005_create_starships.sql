CREATE TABLE starships(
    id BIGSERIAL primary key,
    name CHARACTER VARYING NOT NULL,
    -- based on name
    slug CHARACTER VARYING NOT NULL,
    model CHARACTER VARYING NOT NULL,
    manufacturer CHARACTER VARYING NOT NULL,
    -- unknown translated to NULL
    cost_in_credits INTEGER,
    length FLOAT NOT NULL,
    -- n/a translated to NULL
    max_atmosphering_speed CHARACTER VARYING,
    crew CHARACTER VARYING NOT NULL,
    -- n/a translated to NULL
    passengers INTEGER,
    cargo_capacity INTEGER NOT NULL,
    consumables CHARACTER VARYING NOT NULL,
    hyperdrive_rating FLOAT NOT NULL,
    mglt INTEGER NOT NULL,
    starship_class CHARACTER VARYING NOT NULL,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    edited TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX "index_starships_on_name" ON starships USING btree(
    name
);
CREATE UNIQUE INDEX "index_starships_on_slug" ON starships USING btree(
    slug
);

---- create above / drop below ----

DROP TABLE starships;