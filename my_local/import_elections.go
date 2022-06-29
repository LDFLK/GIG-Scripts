package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	ElectionDataPath = "gig-data-master/elections/"
)

func main() {
	exit := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	// open file
	// PARLIAMENTARY - 	0-entity_id		1-valid		2-rejected	3-polled	4-electors	UNP	SLFP	ELJP	MEP	USA	INDI	TULF	DPLF	ACTC	SLMC	IND	USP	IND1	USF	IND2	IND 1	IND 2
	// PRESIDENTIAL  -	0-entity_id		1-valid	r	2-rejected	3-polled	4-electors	IND01	IND02	IND03	IND04	DUNF	OWORS	IND05	IND06	USP	RJA	DNM	NMPP	FSP	JSWP	IND07	OPPP	IND08	IND09	NDF	JSP	NSSP	NSU	IND10	IND11	SLPP	NUA	SLLP	IND12	IND13	SPSL	IND14	SEP	IND15	ONF	NPP

	//Needs to run decoder in the exact order to allow connecting with parents
	//helpers.AddDecodedData(countrySource, decoders.MyLocalCountryDecoder{}, exit)

	log.Println("Completed importing MyLocal election data.")
}
