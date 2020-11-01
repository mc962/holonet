CREATE TABLE people_species(
    person_id BIGINT REFERENCES people,
    species_id BIGINT REFERENCES species,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    edited TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX "index_people_species_on_people_id" ON people_species USING btree(
    person_id
);
CREATE UNIQUE INDEX "index_people_species_on_species_id" ON people_species USING btree(
    species_id
);

---- create above / drop below ----

DROP TABLE people_species;