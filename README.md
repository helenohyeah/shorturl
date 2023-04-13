## Short url

A link shortener service built with Golang and React

Table of contents

- [Demo](https://github.com/helenohyeah/shorturl#demo)
- [Data model](https://github.com/helenohyeah/shorturl#data-model)
- [Technologies & tools](https://github.com/helenohyeah/shorturl#technologies-and-tools-used--would-use)
- [Technical complexities](https://github.com/helenohyeah/shorturl#technical-complexities)
- [Traffic spike scenario](https://github.com/helenohyeah/shorturl#traffic-spike-scenario)
- [How to run locally](https://github.com/helenohyeah/shorturl#how-to-run-locally)

---

## Demo

![url shortener demo](./docs/demo.gif)

## Data model

Here's a look at what I have in mind for the database schema. We first start with the users and urls table and later look at adding in separate table to track analytics for each url

![url shortener erd](./docs/erd.png)

Note: Postgresql BIGSERIAL has a range of 1 to 9223372036854775807

---

## Technologies and tools used / would use

Frontend:

- React was used because its the frontend framework I'm most comfortable using

Backend:

- Golang was chosen as the backend language because it is designed to handle concurrency and can manage multiple concurrent http requests efficiently

Database:

- Relational DB was chosen over a NoSQL DB because it is easier to ensure data consistency and perform complex queries for analytic features
- Schema can be used to enforce uniqueness of our short URL
- There are options to increase read performance that is covered later
- Use a tool like goose to handle migrations

Infrastructure / Ops

- Security:
  - Hash passwords and secure authorized pages using token based authentication
- Monitoring:
  - Record data on CPU usage, memory usage, and latency to identify bottlenecks in performance using tools like Sentry or DataDog
  - Record logs for easy debugging
- Cloud hosting:
  - Use a cloud provider like AWS so its easy to scale up / down resources to accomodate different traffic patterns and handle traffic spikes
- Deployment:
  - Dockerize services so they can be run in any environment (independent of local environment)
  - Use a tool like Travis CI or Circl CI to automate testing, build, and deployment process

---

## Technical complexities

### URL uniqueness

Short URLs need to be unique to guarantee that they will redirect to the expected long URL. They should also be short, otherwize what's the point of shortening them. Here are a couple of approaches

- Base62 encoding (implemented)

  - Convert base 10 to base 62 to get a unique alphanumeric string
  - As long as our base 10 integers are unique, then our encoded string will be unique
  - It's easy to encode and decode however if our integer is sequential then the short urls will not appear random. We can solve that by deterministically map a sequential integer to a non-sequential integer
  - With a 7 character short URL, we have ~3.5 billion unique different URLs

- Hashing

  - Use an algorithim like MD5 to generate a hash from the long URL with additional data like the current time and userID
  - Take the resulting 128-bit value and encode using base62 and take the first 7 digits to get a 7 digit short URL
  - This solution will not scale, chances of collision increases as our database gets larger

- Standalone key generation service

  - Generate our unique keys ahead of time and store it in a standalone service
  - When a short URL is created we grab a key and mark is as used
  - It will need to ensure concurrent requests do not end up using the same key by using methods that lock a resource

### Low latency DB reads at scale

Our short url service will experience much higher read requests than write so it will become increasingly challenging to ensure efficient data retrieval as our database grows in size and complexity

There are a couple of approaches we can take:

- Caching

  - Store frequently accessed data in memory so they can be read faster

- Database partitioning

  - Split up our DB into smaller multiple subsets of the original database to allow queries to be run in parallel

- Load balancing
  - If we have multiple servers / instances we should distribute incoming requests evenly and based on the load or availability of the server

### Other challenges to consider

- Prevent bad actors from abusing or using the service maliciously
- Prevent creating short urls for websites that don't exist
- Handling circular redirection (redirect to a redirect)

---

## Traffic spike scenario

> Imagine that a celebrity begins using your link shortener, sharing your short links on their widely-followed social media accounts. Their followers begin eagerly clicking on the short links, driving a huge amount of traffic. Which parts of the application do you think will feel the stress of the increased traffic? How might you mitigate this stress?

Effects on the server

- it can result in slow response times as the server tries to respond to each request, potentially server crash and downtime
- users will experience this as slow or failed redirections

Mitigation options

- Caching

  - store frequently accessed data (like this celebrity's short link) in memory so it can be accessed without needing to query the database
  - reduces amount of resource required for each request

- Scale up resources:

  - add more resources to handle increase in traffic (e.g. add more servers)
  - with additional servers, load balancers should be used to distribute the load evenly
  - multiple servers increase availability and resiliency in case one server crashes

- Database / application optimizations:
  - optimize database schema, indexes, or queries
  - identify and fix bottlenecks in the application

# How to run locally

You will need two terminals to run the frontend and backend

**Frontend**

```
$ cd ui
$ npm install
$ npm start
```

**Backend**

Download postgresql, start it locally and create a database

```
$ cd server
$ cp config.yml.template config.yml
```

Fill in database variables with what you set up locally

```
$ go run main.go
```

If you want to run it with hot reloading

```
$ go run github.com/cosmtrek/air
```

Note: the server will regenerate the tables on each reload. To prevent it, comment out `dB.SeedDB()` in `main.go`
