package election

import (
	"GIG-Scripts/my_local/decoders"
	"encoding/json"
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

type MyLocalElectionDecoder struct {
	decoders.MyLocalDecoderInterface
	decoders.MyLocalDecoder
	ElectionType string
	ElectionYear string
	Headers      []string
	VoteCounts   []string
}

func (d MyLocalElectionDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {

	// PARLIAMENTARY - 	0-entity_id		1-valid		2-rejected	3-polled	4-electors	UNP	SLFP	ELJP	MEP	USA	INDI	TULF	DPLF	ACTC	SLMC	IND	USP	IND1	USF	IND2	IND 1	IND 2
	// PRESIDENTIAL  -	0-entity_id		1-valid		2-rejected	3-polled	4-electors	IND01	IND02	IND03	IND04	DUNF	OWORS	IND05	IND06	USP	RJA	DNM	NMPP	FSP	JSWP	IND07	OPPP	IND08	IND09	NDF	JSP	NSSP	NSU	IND10	IND11	SLPP	NUA	SLLP	IND12	IND13	SPSL	IND14	SEP	IND15	ONF	NPP
	d.Name = d.ElectionType + " Election - " + d.ElectionYear
	d.Category = d.ElectionType
	d.Attribute = strings.ToLower(d.ElectionType) + "_elections"
	d.ParentId = record[0]
	d.Source = source
	d.ParentAttribute = "location_id"
	d.Headers = headers[1:]
	d.VoteCounts = record[1:]

	entity := d.GetParentEntity()
	entity.AddCategory("ELECTION").AddCategory(d.ElectionType)
	entity = d.SetPartyWiseResults(entity)
	return entity
}

func (d MyLocalElectionDecoder) SetPartyWiseResults(entity models.Entity) models.Entity {
	if len(d.VoteCounts) != len(d.Headers) {
		log.Println("result count and party count mismatch")
		return entity
	}
	partyWiseResults := make(map[string]string)

	for i := range d.VoteCounts {
		partyWiseResults[d.Headers[i]] = d.roundOffStringValue(d.VoteCounts[i])
	}
	resultBytes, err := json.Marshal(partyWiseResults)
	if err != nil {
		log.Println("error marshaling party results")
		return entity
	}

	sourceDate, err := time.Parse("2006", d.ElectionYear)
	if err != nil {
		log.Fatal("error parsing election year")
	}

	entity.SetAttribute(strings.ToLower(d.ElectionType)+"_election_results", models.Value{
		ValueType:   ValueType.JSON,
		ValueString: string(resultBytes),
		Source:      d.Source,
		Date:        sourceDate,
	})
	return entity
}

func (d MyLocalElectionDecoder) roundOffStringValue(value string) string {
	floatVal, err := strconv.ParseFloat(value, 32)
	if err != nil {
		log.Println("error converting to int: ", err)
		return value
	}
	return strconv.Itoa(int(math.Round(floatVal)))
}
