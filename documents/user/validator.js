db.createCollection("user",
    {
        "validator": {
            "$jsonSchema": {
                "bsonType": "object",
                "required": [
                    "username",
                    "password",
                    "email",
                    "role"
                ],
                "properties": {
                    "username": {
                        "bsonType": "string",
                        "description": "must be a string and is required"
                    },
                    "password": {
                        "bsonType": "string",
                        "minLength": 8,
                        "description": "must be an string with atlease 8 characters and is required"
                    },
                    "email": {
                        "bsonType": "string",
                        "description": "Email required"
                    },
                    "role": {
                        "bsonType": "string",
                        "description": "must be a role of the user"
                    },
                    "CreatedAt": {
                        "bsonType": "date",
                        "description": "must be created date and required"
                    },
                    "UpdatedAt": {
                        "bsonType": "date",
                        "description": "must be updated date and required"
                    }
                }
            }
        }
    }
)