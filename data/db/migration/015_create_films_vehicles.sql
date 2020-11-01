CREATE TABLE films_vehicles(
    film_id BIGINT REFERENCES films,
    vehicle_id BIGINT REFERENCES vehicles,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    edited TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX "index_films_vehicles_on_films_id" ON films_vehicles USING btree(
    film_id
);
CREATE UNIQUE INDEX "index_films_vehicles_on_vehicle_id" ON films_vehicles USING btree(
    vehicle_id
);

---- create above / drop below ----

DROP TABLE films_vehicles;