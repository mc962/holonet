CREATE TABLE films_species(
    film_id BIGINT REFERENCES films,
    species_id BIGINT REFERENCES species,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    edited TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX "index_films_species_on_films_id" ON films_species USING btree(
    film_id
);
CREATE UNIQUE INDEX "index_films_species_on_species_id" ON films_species USING btree(
    species_id
);

---- create above / drop below ----

DROP TABLE films_species;