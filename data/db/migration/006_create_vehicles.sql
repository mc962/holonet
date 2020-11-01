CREATE TABLE vehicles(
    id BIGSERIAL primary key,
    name CHARACTER VARYING NOT NULL,
    -- based on name
    slug CHARACTER VARYING NOT NULL,
    model CHARACTER VARYING NOT NULL,
    manufacturer CHARACTER VARYING ARRAY NOT NULL,
    -- unknown translated to NULL
    cost_in_credits INTEGER,
    length FLOAT NOT NULL,
    max_atmosphering_speed INTEGER NOT NULL,
    crew INTEGER NOT NULL,
    passengers INTEGER NOT NULL,
    -- none translated to 0
    cargo_capacity INTEGER NOT NULL,
    consumables CHARACTER VARYING NOT NULL,
    vehicle_class CHARACTER VARYING NOT NULL,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    edited TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX "index_vehicles_on_name" ON vehicles USING btree(
    name
);
CREATE UNIQUE INDEX "index_vehicles_on_slug" ON vehicles USING btree(
    slug
);

---- create above / drop below ----

DROP TABLE vehicles;