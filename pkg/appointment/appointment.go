package appointment

import (
	"fmt"
	"time"
)

// Appointment es la estructura que representa una cita
type Appointment struct {
	DateTime time.Time
	Patient  string
	Doctor   string
}

// Schedule es la estructura que representa el horario de un médico
type Schedule struct {
	Doctor       string
	Appointments []Appointment
}

// NewAppointment crea una nueva cita
func NewAppointment(dateTime time.Time, patient string, doctor string) Appointment {
	return Appointment{DateTime: dateTime, Patient: patient, Doctor: doctor}
}

// NewSchedule crea un horario para un médico
func NewSchedule(doctor string) Schedule {
	return Schedule{Doctor: doctor, Appointments: []Appointment{}}
}

// AddAppointment agrega una cita a un horario
func (s *Schedule) AddAppointment(appointment Appointment) error {
	for _, a := range s.Appointments {
		if a.DateTime == appointment.DateTime {
			return fmt.Errorf("Ya hay una cita programada para esa fecha y hora")
		}
	}
	s.Appointments = append(s.Appointments, appointment)
	return nil
}

func main() {
	schedule := NewSchedule("Dr. Smith")
	appointment1 := NewAppointment(time.Date(2023, 02, 12, 14, 0, 0, 0, time.UTC), "Jane Doe", "Dr. Smith")
	err := schedule.AddAppointment(appointment)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Cita agendada con éxito")
	}
}
