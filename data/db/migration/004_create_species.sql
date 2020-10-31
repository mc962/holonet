CREATE TABLE species(
    id BIGSERIAL primary key,
    name CHARACTER VARYING NOT NULL,
    -- based on name
    slug CHARACTER VARYING NOT NULL,
    classification CHARACTER VARYING NOT NULL,
    designation CHARACTER VARYING NOT NULL,
    -- n/a translated to NULL
    average_height INTEGER,
    -- float to support infinity, unknown translated to NULL
    average_lifespan FLOAT,
    -- n/a translated to NULL
    hair_colors CHARACTER VARYING ARRAY,
    skin_colors CHARACTER VARYING ARRAY,
    eye_colors CHARACTER VARYING ARRAY,
    language CHARACTER VARYING,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    edited TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX "index_species_on_name" ON species USING btree(
    name
);
CREATE UNIQUE INDEX "index_species_on_slug" ON species USING btree(
    slug
);

---- create above / drop below ----

DROP TABLE species;