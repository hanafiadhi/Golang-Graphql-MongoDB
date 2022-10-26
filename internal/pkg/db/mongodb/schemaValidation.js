db.createCollection("users",{
    validator:{
        $jsonSchema:{
            bsonType: "object",
            title: "users Object Validation",
            required:["username","email","password","age","created_at","updated_at"],
            properties:{
                username:{
                    bsonType:"string",
                    description:"'username' must be string"
                },
                email:{
                    bsonType: "string",
                    description: "'email' must be string"
                },
                age:{
                    bsonType: "int",
                    description:"'age' must be integer"
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