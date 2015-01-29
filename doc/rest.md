
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

    $ curl -i http://localhost:8080/v1/members/dummy

    HTTP/1.1 404 OK
    Content-Type: application/json; charset=utf-8

    {
        "status": 404,
        "message": "Member dummy not found"
    }




