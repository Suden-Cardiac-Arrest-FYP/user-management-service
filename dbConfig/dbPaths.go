package dbConfig

import (
	mongo "go.mongodb.org/mongo-driver/mongo"
)

var DATABASE *mongo.Database

// const DATABASE_URL = "mongodb+srv://evolza-finance-mgt:Asd1234@cluster0.nrv6p.mongodb.net/"

// const DATABASE_NAME ="User-Mgt-APP1514"
const DATABASE_URL = "mongodb+srv://mihishitharushika:CDQXteRVBbjApzMw@sca-cluster.wbptaob.mongodb.net/?retryWrites=true&w=majority&appName=SCA-Cluster"

const DATABASE_NAME = "User-Mgt-Service"
