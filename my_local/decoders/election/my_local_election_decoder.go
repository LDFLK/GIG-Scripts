package election

import (
	"GIG-Scripts/my_local/decoders"
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalElectionDecoder struct {
	decoders.MyLocalDecoderInterface
	decoders.MyLocalDecoder
	ElectionType string
	ElectionYear string
}

func (d MyLocalElectionDecoder) DecodeToEntity(record []string, source string, headers []string) models.Entity {

	// PARLIAMENTARY - 	0-entity_id		1-valid		2-rejected	3-polled	4-electors	UNP	SLFP	ELJP	MEP	USA	INDI	TULF	DPLF	ACTC	SLMC	IND	USP	IND1	USF	IND2	IND 1	IND 2
	// PRESIDENTIAL  -	0-entity_id		1-valid		2-rejected	3-polled	4-electors	IND01	IND02	IND03	IND04	DUNF	OWORS	IND05	IND06	USP	RJA	DNM	NMPP	FSP	JSWP	IND07	OPPP	IND08	IND09	NDF	JSP	NSSP	NSU	IND10	IND11	SLPP	NUA	SLLP	IND12	IND13	SPSL	IND14	SEP	IND15	ONF	NPP

	d.ParentEntity = d.GetParentEntity()
	entity := d.MapToEntity()
	entity.AddCategory("ELECTION").AddCategory(d.ElectionType)
	decoder.AppendToParentEntity(entity)

	return entity
}
