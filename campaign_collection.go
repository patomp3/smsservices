package sms

import (
	"strings"

	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

//MongoDBInfo for ..
type MongoDBInfo struct {
	user       string
	password   string
	host       string
	port       string
	database   string
	collection string
	URL        string
}

// Campaign ...
type Campaign struct {
	Id           bson.ObjectId `bson:"_id"`
	CampaignId   int           `json:"campaignid"`
	CampaignName string        `json:"campaignName"`
	Description  string        `json:"description"`
	StartDate    string        `json:"startdate"`
	EndDate      string        `json:"enddate"`
	PeriodBy     string        `json:"periodby"`
	Schedule     struct {
		Type    string `json:"type"`
		Execute string `json:"execute"`
	} `json:"schedule"`
	Status       string `json:"status"`
	ProfileAllow []int  `json:"profileallow,omitempty"`
	PackageAllow []int  `json:"packageallow,omitempty"`
	KeywordAllow []int  `json:"keywordallow,omitempty"`
	ProductAdd   []struct {
		ProductNr int    `json:"productnr"`
		DayAdd    int    `json:"dayadd"`
		EndDate   string `json:"enddate"`
	} `json:"productadd,omitempty"`
	OfferAdd []struct {
		OfferNr     int    `json:"offernr"`
		DayAdd      int    `json:"dayadd"`
		DayFormular int    `json:"dayformular"`
		MonthAdd    int    `json:"monthadd"`
		EndDate     string `json:"enddate"`
	} `json:"offeradd,omitempty"`
}

// Create for create dbinfo
func Create(user string, pass string, host string, port string, database string, collection string) *MongoDBInfo {
	db := &MongoDBInfo{user: user, password: pass, host: host, port: port, database: database, collection: collection}

	var url string
	if user == "" || pass == "" {
		url = "mongodb://$host:$port/$db"
	} else {
		url = "mongodb://$user:$pass@$host:$port/$db"
	}
	url = strings.Replace(url, "$user", user, -1)
	url = strings.Replace(url, "$pass", pass, -1)
	url = strings.Replace(url, "$host", host, -1)
	url = strings.Replace(url, "$port", port, -1)
	url = strings.Replace(url, "$db", database, -1)

	db.URL = url

	return db
}

// GetCampaigns for using get all documents in collection
func (db *MongoDBInfo) GetCampaigns() []Campaign {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(db.database).C(db.collection)
	result := []Campaign{}
	err = c.Find(nil).All(&result)
	if err != nil {
		return nil
		//log.Fatal(err)
	}

	return result
}

// GetCampaignByID for get document by OjectId
func (db *MongoDBInfo) GetCampaignByID(id string) Campaign {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(db.database).C(db.collection)
	result := Campaign{}
	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	if err != nil {
		return Campaign{}
		//log.Fatal(err)
	}

	return result
}

// GetCampaign for get documents that match field with value
func (db *MongoDBInfo) GetCampaign(field string, value interface{}) []Campaign {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(db.database).C(db.collection)
	result := []Campaign{}
	err = c.Find(bson.M{field: value}).All(&result)
	if err != nil {
		return nil
		//log.Fatal(err)
	}

	return result
}

// InsertCampaign for insert new document
func (db *MongoDBInfo) InsertCampaign(camp Campaign) bool {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(db.database).C(db.collection)
	camp.Id = bson.NewObjectId()
	err = c.Insert(camp)
	if err != nil {
		return false
		//log.Fatal(err)
	}

	return true
}

// RemoveCampaignByID for remove document that match with object id
func (db *MongoDBInfo) RemoveCampaignByID(id string) bool {
	session, err := mgo.Dial(db.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(db.database).C(db.collection)

	err = c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		return false
		//log.Fatal(err)
	}

	return true
}

// RemoveCampaign for remove document that math field & value. (one document)
func (db *MongoDBInfo) RemoveCampaign(field string, value interface{}) bool {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(db.database).C(db.collection)

	err = c.Remove(bson.M{field: value})
	if err != nil {
		return false
		//log.Fatal(err)
	}

	return true
}

// UpdateCampaign for update document by objectId
func (db *MongoDBInfo) UpdateCampaign(camp Campaign) bool {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(db.database).C(db.collection)

	err = c.Update(bson.M{"_id": camp.Id}, camp)
	if err != nil {
		return false
		//log.Fatal(err)
	}

	return true
}
