db.createCollection("photos",{
    validator:{
        $jsonSchema:{
            bsonType: "object",
            title: "Photo Object Validation",
            required:["title","caption","photoUrl","user_id","created_at","updated_at"],
            properties:{
                title:{
                    bsonType: "string",
                    description:"'title' must be string"
                },
                caption:{
                    bsonType: "string",
                     description:"'caption' must be string"
                },
                photoUrl :{
                    bsonType:"string",
                    description :"'photoUrl' musr be string"
                },
                user_id:{ "bsonType": "objectId" },
                comment:{
                    bsonType: "array",
                    properties:{
                        _id :{
                            bsonType: "objectId"
                        },
                    }
                },
                created_at:{
                    bsonType: "int",
                    description:"'created_at' must be integer"
                },
                updated_at:{
                    bsonType: "int",
                    description:"'updated_at' must be integer"
                }
            }

        }
    }
})
db.createCollection("comments",{
    validator:{
        $jsonSchema:{
            bsonType:"object",
            title:"comments Object Validation",
            required:["user_id",]
        }
    }
})
