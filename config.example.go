package GIG_Scripts

import "github.com/lsflk/gig-sdk/client"

/**
Set the GIG server API url here for crawlers
*/
var GigClient = client.GigClient{
	ApiUrl:                 "<ApiUrl>",
	ApiKey:                 "<ApiKey>",
	NerServerUrl:           "<NerServerUrl>",
	NormalizationServerUrl: "<NormalizationServerUrl>",
	OcrServerUrl:           "<OcrServerUrl>",
}
