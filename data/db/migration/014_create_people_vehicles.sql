CREATE TABLE people_vehicles(
    person_id BIGINT REFERENCES people,
    vehicle_id BIGINT REFERENCES vehicles,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    edited TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX "index_people_vehicles_on_people_id" ON people_vehicles USING btree(
    person_id
);
CREATE UNIQUE INDEX "index_people_vehicles_on_vehicle_id" ON people_vehicles USING btree(
    vehicle_id
);

---- create above / drop below ----

DROP TABLE people_vehicles;