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

NOTE: It lacks test configuration, logger configuration, and error handling. One Query and Update mutation