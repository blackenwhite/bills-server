package repo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"log"
)

// Bill

type Bill struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Year  int    `json:"year"`
	Month string `json:"month"`
	Value int    `json:"value"`
}

type BillsCollection struct {
	Data []Bill `json:"data"`
}

type BillResource struct {
	Data Bill `json:"data"`
}

type BillRepo struct {
	Coll *mgo.Collection
}


func (r *BillRepo) All(year string) (BillsCollection, error) {
	billsC := []BillsCollection{}
	yearInt, _ := strconv.Atoi(year)

	pipe := r.Coll.Pipe([]bson.M{
		{"$match": bson.M{"_id": "user1"}},
		{
			"$project": bson.M{
				"data": bson.M{
					"$filter": bson.M{
						"input": "$bills",
						"as":    "bill",
						"cond": bson.M{
							"$eq": []interface{}{
								"$$bill.year",
								yearInt,
							},
						},
					},
				},
				"_id": 0,
			},
		},
	})


	if err := pipe.All(&billsC); err != nil {
		return billsC[0], err
	}

	return billsC[0], nil
}

func (r *BillRepo) Delete(id string) error {
	log.Println("Delete bill", id)
	query := bson.M{
		"_id": "user1",
	}

	delete := bson.M{
		"$pull": bson.M{
			"bills": bson.M{
				"_id": bson.ObjectIdHex(id),
			},
		},
	}

	if err := r.Coll.Update(query, delete); err != nil {
		return err
	}

	return nil
}