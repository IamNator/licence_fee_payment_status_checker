# licence_fee_payment_status_checker
This is part of a packing lot management software, It is an API (application Program Interface) that exposes endpoints allowing users to add new subscribers (cars), identiy and flag defaulters (non-subscribers using the parking lot).  
  
# Endpoints  
POSTMAN DOCUMENTATION = https://www.getpostman.com/collections/133f448fed78751ce2ff    

## 1. Get all Subscribers    

Get all subscribers (both paid and defaulter) in a parking lot.  
Base URL = https://mawakif.herokuapp.com
```json
method - GET
url - https://mawakif.herokuapp.com/api/subscribers

an example of response returned 
{
    "status": true,
    "code": 200,
    "data": [
        {
            "id": 184,
            "plate_number": "67P-78PL",
            "start_time": "2020-03-09T22:18:27Z",
            "status": false
        },
        {
            "id": 594,
            "plate_number": "67P-908PL",
            "start_time": "2021-05-17T18:18:27Z",
            "status": false
        }
    ]
}
```

| field                 | description                                  |
|-----------------------|----------------------------------------------|
| status                | returns true if successful                   |
|code                   | standard http status code                    |
|data                   | this contains the data/resource requested for|



   
## 2. Endpoint for the Pi robot   
- send the car plate number, packing space id, current time.
- returns status ok
```
/api/plate
POST 
{
    plate_number: "45-TY-90",
    packing_space_id: 5674, 
    current_time: 23:00:45
    is_empty: true/false
}

returns 
STATUS -  200 OK
{}
```

## 3. Endpoint for employees   
- sends car plate number, status, current time.
-  returns status created if non exists, or status ok if it's updates a vehicle.

```
/api/subscribers/add
POST
{
    plate_number: "45-GH-46L",
    status: true,
    current_time: 21:50:59
}

status - 200/201
{
    message: "user registered"/ "user registration updated"
}
```

## 4. Endpoint for web interface  
- request for list of cars 
- returns a list of cars recorded by the employees and robot and their status
```
/api/subscribers
GET

returns 
[
    {
        plate_number: "34-534L-34"
        status: true/false,    (indicates if car is paid or unpaid)
        packing_space_id: 4233, 
        message: "upaid"/"paid"/"expired"
    },
    ...
]
```

## 5. Endpoint for admin   
-  sets ricket expiration duration 

```
/api/admin/ticket-duration
PUT
{
    ticket-duration: 24:00:00
}

returns 
STATUS 200 OK if successful.
```
Explanation:

When an employee records a paid user(car owner), the current time is designated as start time and status is 
true. If a car is detected by the robot and it's status is false, it does not exist in records or presence exceeds ticket duration (reset),  
the car is said to be unpaid thus this is recorded  or updated in database with a default start time of zero and a false status.

## 6. An endpoint that returns Packing space Checks    
```
GET 
/api/checks

returns 
STATUS 200 OK
[
    {
        packing_space_id: 10,
        plate_number: "34L-24K-43",
        is_empty: true/false,
        created_at: "23:45:89"

    }
]
```

## 7. An endpoint that add and another that returns a list of packing space and their designation  

```
GET

/api/space 

returns
[
    {
        id: 12,
        designation: "The President"
    }, 
    ...
]
```

```
POST

/api/space 

returns
[
    {
        id: 12,
        designation: "The President"
    }, 
    ...
]
```

## TABLE 1  - subscribers
contains id, plate_number, packing_space_id, start_time, status. It keeps a record all the subscribers (both paid and defaulters).  

## TABLE 2 - checks   
contains id, is_empty, plate_number, packing_space_id, created_at. Data uploaded from the robot is logged here.

## TABLE 3 -  packing_space
contains id, designation. It keeps a record of special designation of each packing space.

## TABLE 4 - config
contains the id, name, value e.g TICKET_DURATION
 
  ```
    
URL = https://mawakif.herokuapp.com   
POSTMAN DOCUMENTATION = https://www.getpostman.com/collections/133f448fed78751ce2ff



        
