package client

type RequestResponse struct {
    StatusCode int
    Status     string
    Headers    map[string][]string
    Body       string
}
