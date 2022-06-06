package extended_models

import (
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
	"io/ioutil"
	"log"
)

type Location struct {
	models.Entity
}

func (l *Location) SetLocationId(countryId string, source string) *Location {
	l.SetAttribute("location_id", models.Value{
		ValueType:   ValueType.String,
		ValueString: countryId,
		Source:      source,
	})
	return l
}

func (l *Location) SetName(name string, source string) *Location {
	l.SetTitle(models.Value{
		ValueType:   ValueType.String,
		ValueString: name,
		Source:      source,
	})
	return l
}

func (l *Location) SetCentroid(centroid string, source string) *Location {
	l.SetAttribute("centroid", models.Value{
		ValueType:   ValueType.JSON,
		ValueString: centroid,
		Source:      source,
	})
	return l
}

func (l *Location) SetPopulation(population string, source string) *Location {
	l.SetAttribute("population", models.Value{
		ValueType:   ValueType.Number,
		ValueString: population,
		Source:      source,
	})
	return l
}

func (l *Location) SetGeoCoordinates(sourcePath string) *Location {
	file, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		log.Fatal("error loading coordinate file", sourcePath)
	}
	if err != nil {
		log.Fatal("error reading coordinate file", sourcePath)
	}

	l.SetAttribute("geo_boundary_coordinates",
		*new(models.Value).
			SetType(ValueType.JSON).SetValueString(string(file)))
	return l
}
