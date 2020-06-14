package gerrors

// go:generate
// Copyright 2020 Nima Ghoroubi. All rights reserved.
// Use of this source code is free by importing in your projects

/*
Package gerrors is a light wrapper on golang standard errors package
Our aim is to make it easy to handle the errors in microservice arch,
And Pass errors both upstream or downstream without an concern about
Decode and encode them.
The returned type all are error which any method can return it and your
Debug mode ( developing mode or production mode ) decides about how to show the errors,
Both in terminal/file output or pass into another service via gRPC,http,etc.

The main features of this package are the below methods:

WrapError:
  Assume you have some custom errors in your service structure which they  describe
  Your in-app service behaviours, and you will get some errors from the methods of
  External packages, like json.Marshal errors.
  You need to know whats the problem in json encoding/decoding when you are developing
  Or debugging your service, but it is unnecessary to show the end-user those errors and
  Only a general error from your service is pretty enough.

  With this method you can wrap your service errors with external packages's errors.
  This method separates your service error and external one with a dash (-) and returns
  Error where you call it.

HandleError:
  The generated and wrapped errors from WrapError method will be displayed in 2 mode;
    First: Develop and debug mode, where you need the complete error message
    Second: Production mode, where you would send the end-user only your customized error
			Without any extra information
*/
