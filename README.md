# GoMail
A simple client to send and receive emails.

![Alt text](assets/GoMail.drawio.svg?raw=true "Title")

## Backend
* Golang with gin
* MongoDB (remote DB based in mLab)
* Kafka

## Setting up
```bash
cd email-client/src
make build
make run
```

## File Structure
```
Model - For modelling entities
Routes - Controller
Handler - Service
Middleware - interceptor
Helper - Utils
```

## Requirements
### Fetch Inbox

The user should be able to fetch their inbox emails. The task is to create an endpoint “localhost:8080/user/mail/inbox” to fetch the inbox.

 

a. Choose appropriate HTTP methods (GET,POST,PUT,DELETE) to hit this URL.

 

b. There should be a key-value pair of “Authorization: <JWT_TOKEN>” & if it's incorrect or not present, the API should throw a Bad Request Exception with 401 Unauthorized code.

The username can be identified from the JWT_TOKEN.

 

c. In case the user name exists in the database we should fetch the inbox details sorted by date (descending order of date), as a response with appropriate HTTP status.

### Send Email API

The user should be able to send an email to another user.The task is to create an endpoint localhost:8080/user/mail/send to send an email to a user

 

a. Choose appropriate HTTP methods (GET,POST,PUT,DELETE) to hit this URL.

b. The user needs to provide the details of the email like (sender email address, list of receiver’s email address, subject line, email body) in the request body.

 

Basic validation on the attributes of the email needs to be performed with the following rules :

• Receiver Email address: It should be a valid email ID format
• Subject: It should be a minimum of 2 characters and a maximum of 100 characters
• Email Body: It should be a minimum of 10 chars and a maximum of 1500 characters
• Date: It should be sent as the current date
 

c. In case a user provides wrong input, they should get a message saying “Invalid Details” with the appropriate HTTP status.

 

d. In case the user provides the right set of inputs, the data should get stored in the database and we should get a response with a message “ Email sent successfully” with appropriate HTTP status.

### Fetch Outbox

The task is to create an endpoint “localhost:8080/user/mail/outbox” to fetch the outbox of a user.

 

a. Choose appropriate HTTP methods (GET, POST, PUT, DELETE) to hit this URL.

 

b. There should be a key-value pair of “Authorization: <JWT_TOKEN>” & if it's incorrect or not present, the API should throw a Bad Request Exception with 401 Unauthorized code.

 

c. The username of the user can be identified from the JWT_TOKEN.

 

d. In case the user name exists in the database we should fetch the outbox details sorted by date (descending order of date), as a response with appropriate HTTP status.