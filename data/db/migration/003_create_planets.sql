CREATE TABLE planets(
    id BIGSERIAL PRIMARY KEY,
    name CHARACTER VARYING NOT NULL,
    -- based on name
    slug CHARACTER VARYING NOT NULL,
    rotation_period INTEGER NOT NULL,
    orbital_period INTEGER NOT NULL,
    diameter INTEGER NOT NULL,
    climate CHARACTER VARYING NOT NULL,
    gravity CHARACTER VARYING NOT NULL,
    terrain CHARACTER VARYING NOT NULL ,
    -- unknown translated to NULL
    surface_water INTEGER,
    -- unknown translated to NULL
    population INTEGER,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    edited TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX "index_planets_on_name" ON planets USING btree(
    name
);
CREATE UNIQUE INDEX "index_planets_on_slug" ON planets USING btree(
    slug
);

---- create above / drop below ----

DROP TABLE planets;