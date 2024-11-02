package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"yam-client/pkg/parser"
	"yam-client/pkg/util"
)

func executeRequest(requestIndex int,request *parser.YamlRequest)(RequestResponse,error){

	if request.Name =="" {
		
	}

	err :=validateRequest(request)
	
	if err != nil {
		return RequestResponse{} , err
	}

	return execute(*request)
}


func execute(request parser.YamlRequest) (RequestResponse,error){
	var err error
	var req *http.Request

	client := &http.Client{}

	reqUrl,err := url.Parse(request.Url)

    if err != nil {
        return RequestResponse{}, fmt.Errorf("Error occurred parsing url: %w", err)
    }

   
	if request.QueryParams != nil {
		query:=reqUrl.Query()
		for key ,val :=range request.QueryParams{
			query.Add(key,val)
		}
		reqUrl.RawQuery = query.Encode()  //encode query
	}

	
	
	if request.Method != http.MethodGet {
		req , err  = http.NewRequest("GET",reqUrl.String(),bytes.NewBuffer([]byte(request.RequestBody))) 
		if err !=nil {
			return RequestResponse{}, fmt.Errorf("Error occurred creating new request", err)
		}
	}else{
		req , err  = http.NewRequest("GET",reqUrl.String(),nil) 
		if err !=nil {
			return RequestResponse{}, fmt.Errorf("Error occurred creating new request", err)
		}
	}

	//Add headers
	if request.Headers != nil {
		for key, val := range request.Headers{
			req.Header.Add(key,val)
		}
	}

	res,err := client.Do(req)

	if err!=nil {
		return RequestResponse{}, fmt.Errorf("error executing request: %w", err)
	}
	defer res.Body.Close()

	body , err := io.ReadAll(res.Body)
	
	if err != nil {
		return RequestResponse{}, fmt.Errorf("error occurred reading request body: %w", err)
	}

	return RequestResponse{
		StatusCode: res.StatusCode,
		Headers: res.Header,
		Body: string(body),
	},err
}


//validate specificRequest on the file
func validateRequest(request *parser.YamlRequest) (error) {

	// if request.Name== ""{
	// 	return fmt.Errorf("REQUEST %v  IS MISSING A NAME",request.Name)   //maybe this is not important need to revisit it though
	// }

	if request.Method== ""{
		return fmt.Errorf("REQUEST %v  IS MISSING A NAME",request.Name)
	}

	validateRequestMethod(request.Name,request.Method)

	if request.RequestBody== "" && request.Method != request_methods.GET { 
		return fmt.Errorf("REQUEST %v  IS MISSING A REQUEST BODY",request.Name)
	}

	return nil
}

func validateRequestMethod(requestName string, requestMethod string) (error){

    switch requestMethod {

    case http.MethodGet, http.MethodPut:
        return nil
		
    default:
        return fmt.Errorf("REQUEST %d HAS AN UNSUPPORTED METHOD: %s", requestName, requestMethod)
    }
}