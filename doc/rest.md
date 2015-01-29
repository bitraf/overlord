
# Overlord Rest API

This document describes Bitraf's Overlord Rest API. It's work in progress and not everything in this document is implemented.

## Overview

The API described here is generally restful and returns the result in json format. We use a simple versioning of the
rest services described in the URL itself. For v1 of this API, the base URL is:

    http://localhost:8080/v1

## Error

If an error occurs it will return the reason for the error in a standardized json message. It will also include the
right HTTP status-code in the response. For example, if a member was not found, it will return the following
response:

    $ curl -i http://localhost:8080/v1/members/199

    HTTP/1.1 404 OK
    Content-Type: application/json; charset=utf-8

    {
        "status": 404,
        "message": "Member 199 not found"
    }

## Member Resource

A member object holds information about one single member at Bitraf.

### List members

List members in the database. Members can be filtered by basic query parameter. The result is paged.

    GET /v1/members

#### Parameters

Name | Default | Description
--- | --- | ---
*page* | 0 | Page number to view.
*count* | 20 | Number of items per page.
*query* | | Simple "contains" query on certain fields.

#### Result

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8

    {
        "page: "0"
        "count": "1",
        "total": "4",
        "result": [
            {
                "id": 199,
                "fullName": "John Doe",
                "email": "john@doe.com",
                "joinDate", "2014-11-12",
                "organization": "Bertha Foundation",
                "price": 500,
                "flag": "m_office",
                "recurrence: "1 mon"
            }
        ]
    }


### Get member

Returns a single member.

    GET /v1/members/{id}

#### Result

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8

    {
        "id": 199,
        "fullName": "John Doe",
        "email": "john@doe.com",
        "joinDate", "2014-11-12",
        "organization": "Bertha Foundation",
        "price": 500,
        "flag": "m_office",
        "recurrence: "1 mon"
    }


### Delete member

Deletes a single member.

    DELETE /v1/members/{id}

#### Result

    HTTP/1.1 204 No Content


### Update member

Updates an existing member.

    POST /v1/members/{id}

    {
        "fullName": "John Doe",
        "email": "john@doe.com",
        "joinDate", "2014-11-12",
        "organization": "Bertha Foundation",
        "price": 500,
        "flag": "m_office",
        "recurrence: "1 mon"
    }

#### Result

    HTTP/1.1 200 OK

    {
        "id": 199,
        "fullName": "John Doe",
        "email": "john@doe.com",
        "joinDate", "2014-11-12",
        "organization": "Bertha Foundation",
        "price": 500,
        "flag": "m_office",
        "recurrence: "1 mon"
    }

### Create member

Creates a new member.

    POST /v1/members

    {
        "account": 56,
        "fullName": "John Doe",
        "email": "john@doe.com",
        "joinDate", "2014-11-12",
        "organization": "Bertha Foundation",
        "price": 500,
        "flag": "m_office",
        "recurrence: "1 mon"
    }

#### Result

    HTTP/1.1 201 Created
    Location: http://localhost:8080/v1/members/199

    {
        "id": 199,
        "fullName": "John Doe",
        "email": "john@doe.com",
        "joinDate", "2014-11-12",
        "organization": "Bertha Foundation",
        "price": 500,
        "flag": "m_office",
        "recurrence: "1 mon"
    }


## Account Resource

One account can have multiple members.

### List accounts

    GET /v1/accounts

#### Result

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8

    {
        "page: "0"
        "count": "1",
        "total": "4",
        "result": [
            {
                "id": 56,
                "name": "john",
                "lastBilled": "2014-11-12"
            }
        ]
    }


### Get account by id

Returns a single account by id.

    GET /v1/accounts/{id}

#### Result

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8

    {
        "id": 56,
        "name": "john",
        "lastBilled": "2014-11-12",
        "members:" [
            {
                "id": 199,
                "fullName": "John Doe",
                "href": "http://localhost:8080/v1/members/199"
            }
        ]
    }

### Delete account

Deletes a single account.

    DELETE /v1/accounts/{id}

#### Result

    HTTP/1.1 204 No Content


### Update account

Updates an existing account.

    POST /v1/accounts/{id}

    {
        "name": "john",
        "lastBilled": "2014-11-12"
    }

#### Result

    HTTP/1.1 200 OK

    {
        "id": 56,
        "name": "john",
        "lastBilled": "2014-11-12",
        "members:" [
            {
                "id": 199,
                "fullName": "John Doe",
                "href": "http://localhost:8080/v1/members/199"
            }
        ]
    }

### Create account

Creates a new account.

    POST /v1/accounts

    {
        "name": "john",
        "lastBilled": "2014-11-12"
    }

#### Result

    HTTP/1.1 201 Created
    Location: http://localhost:8080/v1/accounts/59

    {
        "id": 56,
        "name": "john",
        "lastBilled": "2014-11-12",
        "members:" [
            {
                "id": 199,
                "fullName": "John Doe",
                "href": "http://localhost:8080/v1/members/199"
            }
        ]
    }

