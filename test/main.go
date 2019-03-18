package main

import (
	"fmt"
	"log"

	sms "github.com/patomp3/smsservices"
)

func main() {
	fmt.Printf("hello")

	db := sms.Create("", "", "172.19.218.104", "27017", "tvscampaigndb", "campaign")
	log.Printf("%s", db.URL)

	//var result []Campaigns
	//log.Printf("Get all rows")
	//results := db.GetCampaigns()
	//log.Printf("%v", results)

	//log.Printf("Get by id")
	//result := db.GetCampaignByID("5c8a0d21837c11431a1cec42")
	//campaigns, _ := json.Marshal(result)
	///log.Printf("%v", result)

	log.Printf("Get by field Id")
	resultz := db.GetCampaign("campaignid", 2)
	log.Printf("%v", resultz)

	resultz[0].Status = "A"
	resultz[0].OfferAdd = nil
	log.Printf("update")
	xxx := db.UpdateCampaign(resultz[0])
	log.Printf("insert result %t", xxx)
	//log.Printf("Get by field")
	//resultzz := db.GetCampaign("campaignname", "test")
	//log.Printf("%v", resultzz)

	//log.Printf("insert ")
	//c := sms.Campaign{CampaignId: 4, CampaignName: "Campaign 4", Status: "A"}
	//xxx := db.InsertCampaign(c)
	//log.Printf("insert result %t", xxx)

	//log.Printf("remove")
	//xxx := db.RemoveCampaign("campaignid", "4")
	//log.Printf("insert result %t", xxx)

	// get campaign id= 3 and update status 'A'
	/*log.Printf("Get")
	camp := db.GetCampaign("campaignid", "3")
	camp[0].Status = "C"
	log.Printf("update")
	xxx := db.UpdateCampaign(camp[0])
	log.Printf("insert result %t", xxx)*/
	// end
}
