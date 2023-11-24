# ADMIN FRANCHISE

## Description

This is an application to manage the franchises of a company.

## How to run the project

The easy way to run the project is whit the command

```bash
make docker-compose-up
```

But first you will need to setup the `.env` file, use the variables in the `.env.example` file as a guide.

You need to have installed docker and docker-compose in your machine.

Then you will need to run the migrations

```bash
migrate -path ./setup/migrations -database "postgres://user:password@localhost:5433/admin_franchise?sslmode=disable" up
```

Use the right envs for the previous command.

## Flow Diagram

The next diagram is a representation of the flow of the application for the creation and update franchise use cases.

![Flow Diagram](./docs/images/ClubHub-tech-test-Flows.jpg)

The use cases to get franchise and get company are direct queries to the database, so they are not represented in the diagram.

In the next image you can see a more specific flow for the creation request.

![Flow Diagram](./docs/images/ClubHub-tech-test-Flows-Creation.jpg)

It is used an async process to handle the scrapping process in another goroutine, so the request can be finished and the user can get a response. The scrapping process is handled by the strategy ChannelOwner/ChannelUtilizer, so we can have a better control of the process and we can handle the errors in a better way.

## Datamodel

The next diagram is a representation of the datamodel of the application.

![Datamodel](./docs/images/ClubHub-tech-test-DataModel.jpg)

The two tables `transaction_log_trailing` and `incomplete_franchise` are used to help to edge cases of the application. The first one is used to handle async process in memory with the strategy ChannelOwner/ChannelUtilizer, so we can know if a event was processed or not. The second one is used to handle the case when a franchise has dirty data and it is not possible to save like a franchise, so we save it in this table and we can handle it later.

## API Documentation

The API was built with graphql. The following are the schemas, queries and mutations that can be used to interact with the API.

### Schemas

```graphql
type Franchise {
  id: ID!
  title: String!
  siteName: String!
  url: String!
  company: Company!
  location: Location!
  addressLocation: AddressLocation!
}

type OwnerCompany {
  id: ID!
  name: String!
  firstName: String!
  lastName: String!
  email: String!
  phone: String!
  location: Location!
  addressLocation: AddressLocation!
}

type Company {
  id: ID!
  name: String!
  owner: OwnerCompany!
  tax_number: String!
  location: Location!
  addressLocation: AddressLocation!
  franchises: [Franchise!]!
}

type Location {
  id: ID!
  city: String!
  state: String!
  country: String!
}

type AddressLocation {
  id: ID!
  location: Location!
  address: String!
  zipCode: String!
}

input FranchiseCriteria {
  name: String
}

input LocationCriteria {
  city: String
  state: String
  country: String
}

input AddressLocationCriteria {
  address: String
  zipCode: String
}

input FranchisesCriteria {
  name: String
  company: ID
  location: LocationCriteria
}

input CompanyCriteria {
  name: String
}

type Query {
  # Supporting only getting franchises by name
  getFranchise(criteria: FranchiseCriteria): Franchise!
  getCompany(criteria: CompanyCriteria): Company!
  getFranchises(criteria: FranchisesCriteria): [Franchise!]!
}

input NewFranchise {
  url: String!
}

input UpdateFranchise {
  url: String
  name: String
  company: ID
  location: LocationCriteria
  addressLocation: AddressLocationCriteria
}

type Mutation {
  createFranchise(input: NewFranchise!): Franchise!
  updateFranchise(id: ID!, input: UpdateFranchise!): Franchise!
}
```

You can use the next operations when you are interacting with the API:

### Queries

```graphql
query franchiseQuery($criteria: FranchiseCriteria!) {
  getFranchise(criteria: $criteria) {
    id
    name
    company {
      id
    }
    location {
      id
    }
  }
}

{
  "criteria": {
    "name": "Marriott Bonvoy"
  }
}

query companyQuery($input: CompanyCriteria!) {
  getCompany(criteria: $input) {
    id
    name
    owner {
      id
    }
  }
}

{
  "input": {
    "name": "Marriott International, Inc."
  }
}
```

### Mutations

```graphql
mutation createFranchiseMutation($input: NewFranchise!) {
  createFranchise(input: $input) {
    id
    url
  }
}

{
  "input": {
    "url": "https://marriott.com/"
  }
}
```

NOTE: logger configuration, and error handling. Update mutation