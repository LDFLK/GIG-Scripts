package extended_models

import (
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
)

type Election struct {
	models.Entity
}

func (e *Election) SetElectionType(electionType string, source string) *Election {
	e.SetAttribute("election_type", models.Value{
		ValueType:   ValueType.String,
		ValueString: electionType,
		Source:      source,
	})
	return e
}

func (e *Election) SetElectionYear(electionYear string, source string) *Election {
	e.SetAttribute("election_year", models.Value{
		ValueType:   ValueType.Number,
		ValueString: electionYear,
		Source:      source,
	})
	return e
}
