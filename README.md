# ADMIN FRANCHISE

## Description

## How to run the project

## Flow Diagram

## Datamodel

## C2

## API Documentation

Graphql Documentation

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

NOTE: logger configuration, and error handling. Update mutation