package reservations

import (
	"github.com/vsivarajah/AirlineReservation/domain/flights"
	"github.com/vsivarajah/AirlineReservation/domain/passengers"
)

type Reservation struct {
	Id         int                      `json:"id"`
	IsValid    bool                     `json:"paymentsuccessful"`
	Passenger  passengers.PassengerInfo `json:"passenger"`
	FlightInfo flights.Flight           `json:"flightinfo"`
}
type Reservations []*Reservation

var reservationDetails = []*Reservation{}

func CreateFlightDetails(reservation *Reservation) {
	reservationDetails = append(reservationDetails, reservation)
}

func GetReservationDetails() Reservations {
	return reservationDetails
}

func FindReservationById(id int) (*Reservation, int, error) {
	for i, v := range reservationDetails {
		if v.Id == id {
			//Found id
			return v, i, nil
		}
	}
	return nil, -1, nil
}
func UpdateReservation(id int, r *Reservation) error {
	value, pos, err := FindReservationById(id)
	if err != nil {
		return err
	}
	value.Id = id
	value.IsValid = true
	reservationDetails[pos] = value
	return nil
}