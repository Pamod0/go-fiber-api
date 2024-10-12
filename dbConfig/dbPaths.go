package dbConfig

import (
	mongo "go.mongodb.org/mongo-driver/mongo"
)

var DATABASE *mongo.Database

// const DATABASE_URL = "mongodb+srv://cgaas:rvyuMzkZXfLp52m7@cgaas.bbsin5h.mongodb.net/?retryWrites=true&w=majority"

const DATABASE_URL = "mongodb+srv://pamod:29StecdLbCivz7nR@pamod.inscxbr.mongodb.net/?retryWrites=true&w=majority&appName=Pamod"

const DATABASE_NAME = "gofiberapi"
