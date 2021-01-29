package db

import (
	"context"
	"github.com/jackc/pgx/v4"
	"holonet/data/resource"
)

type Vehicle struct {
	DB *pgx.Conn
}

func FindVehicle(id int) (resource.Vehicle, error) {
	var vehicle resource.Vehicle
	row := DB.QueryRow(context.Background(), "SELECT * FROM vehicles WHERE id = $1;", id)

	err := row.Scan(
		&vehicle.Name,
		&vehicle.Model,
		&vehicle.Manufacturer,
		&vehicle.CostInCredits,
		&vehicle.Length,
		&vehicle.MaxAtmospheringSpeed,
		&vehicle.Crew,
		&vehicle.Passengers,
		&vehicle.CargoCapacity,
		&vehicle.Consumables,
		&vehicle.VehicleClass,
		&vehicle.Pilots,
		&vehicle.Films,
	)

	if err == nil {
		return vehicle, nil
	} else {
		return resource.Vehicle{}, err
	}
}

func AllVehicles() ([]resource.Vehicle, error) {
	rows, err := DB.Query(context.Background(), "SELECT * FROM vehicles;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicles []resource.Vehicle

	for rows.Next() {
		var vehicle resource.Vehicle

		err := rows.Scan(
			&vehicle.Name,
			&vehicle.Model,
			&vehicle.Manufacturer,
			&vehicle.CostInCredits,
			&vehicle.Length,
			&vehicle.MaxAtmospheringSpeed,
			&vehicle.Crew,
			&vehicle.Passengers,
			&vehicle.CargoCapacity,
			&vehicle.Consumables,
			&vehicle.VehicleClass,
			&vehicle.Pilots,
			&vehicle.Films,
		)
		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, vehicle)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return vehicles, nil
}
