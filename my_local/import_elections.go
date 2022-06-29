package main

import (
	"GIG-Scripts/my_local/decoders"
	"GIG-Scripts/my_local/helpers"
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
	// country_id	province_id	district_id	dsd_id	gnd_id	ed_id	pd_id	lg_name		lg_id	moh_id
	countrySource := DataPath + "country.tsv"
	provinceSource := DataPath + "province.tsv"
	districtSource := DataPath + "district.tsv"
	dsdSource := DataPath + "dsd.tsv"
	gndSource := DataPath + "gnd.tsv"
	edSource := DataPath + "ed.tsv"
	pdSource := DataPath + "pd.tsv"
	lgSource := DataPath + "lg.tsv"
	mohSource := DataPath + "moh.tsv"

	//Needs to run decoder in the exact order to allow connecting with parents
	helpers.AddDecodedData(countrySource, decoders.MyLocalCountryDecoder{}, exit)
	helpers.AddDecodedData(provinceSource, decoders.MyLocalProvinceDecoder{}, exit)
	helpers.AddDecodedData(districtSource, decoders.MyLocalDistrictDecoder{}, exit)
	helpers.AddDecodedData(dsdSource, decoders.MyLocalDSDDecoder{}, exit)
	helpers.AddDecodedData(gndSource, decoders.MyLocalGNDDecoder{}, exit)
	helpers.AddDecodedData(edSource, decoders.MyLocalEDDecoder{}, exit)
	helpers.AddDecodedData(pdSource, decoders.MyLocalPDDecoder{}, exit)
	helpers.AddDecodedData(lgSource, decoders.MyLocalLGDecoder{}, exit)
	helpers.AddDecodedData(mohSource, decoders.MyLocalMOHDecoder{}, exit)

	log.Println("Completed importing MyLocal data.")
}
