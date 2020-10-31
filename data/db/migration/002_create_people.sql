CREATE TABLE people(
    id BIGSERIAL PRIMARY KEY,
    name CHARACTER VARYING NOT NULL,
    -- based on name
    slug CHARACTER VARYING NOT NULL,
    height INTEGER NOT NULL,
    mass INTEGER NOT NULL,
    hair_color CHARACTER VARYING NOT NULL,
    skin_color CHARACTER VARYING NOT NULL,
    eye_color CHARACTER VARYING NOT NULL,
    birth_year CHARACTER VARYING NOT NULL,
    gender CHARACTER VARYING NOT NULL,
    homeworld CHARACTER VARYING NOT NULL,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    edited TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX "index_people_on_name" ON people USING btree(
     name
);
CREATE UNIQUE INDEX "index_people_on_slug" ON people USING btree(
    slug
);

---- create above / drop below ----

DROP TABLE people;
