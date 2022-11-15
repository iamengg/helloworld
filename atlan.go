About product - 
Collect -   
    is a data collection platform.
    that is being used by customers in 50+ countries.
    250 organizations
    11 million responses

    features, 
        team management
        multilingual forms 
        offline data collection

    Problem - 
        Submission of a response
        Post Submission business logic 
        1. Search for slangs in local language for an answer
        2. Validate responses coming in against set of business rules
        3. Want all data in google sheets
            Connect to CRM, Graphs, Charts offered by sheets 
        Each response of form becomes row in sheet 
        question in form become column 
        4. Client want to sent an SMS to the customer whose details are collected in response as soon as ingestion was complete reliably.

        Each new use case can just be plugged in & does not need an overhaul on the backend.
        OPtimize for latency & having an unified interface acting as middleman.

        Design sample schematic for how you would store forms with questions
        response, answes 

        Make assumption where needed.
        must be failsafe, ----------
            Once identifiable data is entered then on or after entering data but if not submitted it saves to browser cache.
        Load balancer, to make it fail safe , we should have redundant load balancer.
        DB - Consider we are supporting text only data, So here SQL based DB will suffice.
        
        We should have Master-Master DB, where One of master will be focused on outside form submissions while other to support analytics service.

        Kafka Queue - For notifications to be sent we should use kafka where on successful injestion of data the notification services reads message from kafka & sends it to user who entered data.

        eventually recover from power/internet outages --------

        scale to cases like millions of respnoes
        How would you benchmark, setup logs, monitor a system health, alerts .

        Design specs.
        brief but comprehensive explanaiton for the approaches
        codebase.


        How google form works
        How google sheet works
        support & extension for them with our problem 



    How to store form
        form link
        Form data
            Customer info
                phone, email, name, address
            Questions
                answer options 1. sinlgle line, 2, Multi line, 3. Option based
            Response store
                Question , answer
    
    Schematic for form
       Design flow
    PSeudo code 

Collect -   
    is a data collection platform.
    that is being used by customers in 50+ countries.
    250 organizations
    11 million responses

    features, 
        team management
        multilingual forms 
        offline data collection

    Problem - 
        Submission of a response
        Post Submission business logic 
        1. Search for slangs in local language for an answer
        2. Validate responses coming in against set of business rules
        3. Want all data in google sheets
            Connect to CRM, Graphs, Charts offered by sheets 
        Each response of form becomes row in sheet 
        question in form become column 
        4. Client want to sent an SMS to the customer whose details are collected in response as soon as ingestion was complete reliably.

        Each new use case can just be plugged in & does not need an overhaul on the backend.
        OPtimize for latency & having an unified interface acting as middleman.



Functional requirement - 

    1. Form creation
    2. Submission of response 
    3. Search engine for slang
    4. Response validator 
    5. Data export option to google sheets
    6. Notification service

Non functional requirement - 
    High availability for forms
    Low latency for form viewing
    Scale - 
        millions of surveys
            Thousands of responses for each survey

Assumptions
        Considering HLD for this MVP
        Consider only once response is getting submitted by any user for any form
        Customer enters valid identifiable info (email, name etc.)
        No limit on creating questions.
        Option to configure payments based on number of questions in survey.
        One time send , No edit once submitted 

Components selection & explanation :
1. Admin 
    Admin Creates survey,
    Gets link for it.
    Shares link with users.

2. Data form
    This is UI where clients submits there survey.
    Clients can submit their survey.
    Response gets' Saved in DB

3. Cache Redis
    When every user hits link cache serves it.

4. DB - SQL This will be structured data
    For each response DB saves it's against survey 
    When ever link is hits db , it' retuns response.

5. Loadbalancer       
    Distributes load across form services.
6. Analytics service :
    This service is PLatform where All users scenarios were accessed.
    1. Search service (Slang search)    
    2. Google sheet interface
    3. Data validations
7. Notification Queue - 
    Once data injestion is done, then analytics sends event to Notification Queue, which sends notification.


Consider below are the API services for survey,

POST /survey    -----create survey
GET /survey/id  -----Read survey
PUT /survey/id   ----UPdate survey



Let's discuss approach for POST /survey API

POST /survey    -----create survey
parameters - 
    {survey_id, survey_name, form_json}
user creates survey, with name
User selects questions & their types with options . (text, multiline txt, mcq etc)    
user saves form
Saved form gets uploaded along with survey
Creates link for form ,
User get's link to share
Default update status as survey is active
Default show count of responses


Nosql DB's
Survey [survey_id, survey_name, status, responseCount]
surveyData [survey_id, form_json, Answer_json, user_id]
userInfo [user_id, user_name, user_email, user_phone, user_address]

Requirement :
What problem you are trying to solve
how users be interacting with this system 

MVP
    Inventory
        search
        list
    Delivery    
        Notification thr email,sms
        Order tracking
    Payment
    Cart 
        use Add (Inc or dec quantity)
            Remove

Find how user using
FInd priority of items
Pick items which impact usrs more 

Monolith or Microservices
Design & develop in such a way that it's decoupled & easy to grow

Elastic search data store
    Txt search engine, designed as nosql engine
    tokenizing string & txt search algos.
Index on where clause

How could monitor & debug 
Vision for system 

What could be problems while designing

FR
    Search
    Cart/Wishlist
    Delivery
    Payment
    View orders
NFR
    Highly available
    High consistency
    Low latency


Ask whether it's LLD or HLD    
ON knowing problem see what all want to design
    what specific pts, funtional requriement are needed for design 
Clarifiy FR & scale (for 100 users or billions of users), throughput, datasize

Time your solution pretty well, 
Ask interviewer what all funtionalities we want to do

If you don't know if something, then ask how it's & say if it's you you will do it in this way.
Ask if you are going in right direction.

Choice of DB
    Query pattern
    Type of data
    Amount of scale to handle

    cache
        quering multiple times
        taking long time then go for cach
        key - query-pmrs, value response
        Redis, etcd, memcache

    blob storage - images, video
        To serve data as it is.
        Amazon S3
    CDN - Distributing image to lot's of locations
        ex. image is stored in S3, YOu want to serve that image across world
    Txt searching capability - 
        Apache lucene -> Elastic search/solar
        Fuzzy search - 
        They are not primary DB
    Application metric tracking system -
        Latency, througput, 
    Time Series DB - Extension of Relational db, but some extensions
        ability to query sequentital update, 
        append & read mode
        query to last few minutes , hourse
        ex. implz db, openTSDB        

Structured data
    Need ACID
    rdbms
        MySQL, Oracle, SQL Server, Postgres
Non structured data    
    Data types, queries
        Document db
            mongo, couchbase
    Ever increasing data  & finite queries
        Columner db
            casandra, hbase


Hotel service 
POST /hotels
GET /hotels/id 
PUT /hotels/id 
PUT /hotels/id/rooms/id 

HOtel DB 
hotel [id, locality_id, description, name, images, is_active]
hotel_facilities[id, hotel_id, facility_id, is_active]
rooms [id, hotel_id, display_name, is_Active, quantity, price_min, price_max]
room_facilities [id, hotel_id, facility_id, is_active]
facilities[id, display_name]
locality [id, city_id, state_id, country_id, zipcode, is_active]

POST /book
{
    user_id, hotel_id, room_id, quantity, start_date, end_date
}
check in available rooms, if yes then block rooms , reduce count of availalble rooms
put in redis with TTL
put in kafka 
redirect to payment

Booking DB
Available_rooms [roomID, date, initial_qty, available_qty]
booking[booking_id, room_id, user_id, startdatae, enddate, noofrooms, invoce id]
status [reserved, cancelled, booked]
available rooms , room_id, date, initialqty, availbleqty
booking bid, rid, uid, stdate, enddate, noforooms, status, invoice id


invoce - invoceid, price, paymentStatus

Payment status -
    failed,
    success
    TTL Redis



