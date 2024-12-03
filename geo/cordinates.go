package geo

import (
	"errors"
	"unicode/utf8"
)

type Coordinates struct {
	latitude  float64
	longitude float64
}

type Landmark struct {
	name string
	Coordinates
}

func (l *Landmark) Name() string {
	return l.name
}

func (l *Landmark) SetName(name string) error {
	if utf8.RuneCountInString(name) > 30 {
		return errors.New("invalid name")
	}
	l.name = name
	return nil
}

func (c *Coordinates) Latitude() float64 {
	return c.latitude
}

func (c *Coordinates) SetLatitude(latitude float64) error {
	if latitude < -90 || latitude > 90 {
		return errors.New("invalid latitude")
	}
	c.latitude = latitude
	return nil
}

func (c *Coordinates) Longitude() float64 {
	return c.longitude
}

func (c *Coordinates) SetLongitude(longitude float64) error {
	if longitude < -180 || longitude > 180 {
		return errors.New("invalid longitude")
	}
	c.longitude = longitude
	return nil
}
