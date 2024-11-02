package parser

import (
	"fmt"
	"yam-client/pkg/util"
)

//validates the file
func validateRequests(parseData YamlParserConfig)(error){
	if !requestsExists(parseData) {
		return fmt.Errorf("NO REQUESTS WERE FOUND ON THE FILE")
	}

    return nil
}

//checks if the file contains requests
func requestsExists(parseData YamlParserConfig) bool{
	return len(parseData.Requests) > 0
}

//validate specificRequest on the file
func validateRequest(requestIndex int, request YamlRequest) (error) {

	if request.Name== ""{
		return fmt.Errorf("REQUEST %v  IS MISSING A NAME",requestIndex)   //maybe this is not important need to revisit it though
	}

	if request.Method== ""{
		return fmt.Errorf("REQUEST %v  IS MISSING A NAME",requestIndex)
	}

	validateRequestMethod(requestIndex,request.Method)

	if request.RequestBody== "" && request.Method != request_methods.GET { 
		return fmt.Errorf("REQUEST %v  IS MISSING A REQUEST BODY",requestIndex)
	}

	return nil
}

func validateRequestMethod(requestIndex int, requestMethod string) (error){

    switch requestMethod {

    case request_methods.GET, request_methods.POST, request_methods.PUT, 
		 request_methods.DELETE, request_methods.PATCH, 
		 request_methods.OPTIONS, request_methods.HEAD:
        return nil
		
    default:
        return fmt.Errorf("REQUEST %d HAS AN INVALID METHOD: %s", requestIndex, requestMethod)
    }
}