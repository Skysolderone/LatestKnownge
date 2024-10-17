package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// curl --location --request POST 'https://solana-mainnet.g.alchemy.com/v2/027Cn34BB8HylGsuSfqvRYXelnJl0_H6' \
// --header 'Content-Type: application/json' \
//
//	--data-raw '  {
//	    "jsonrpc": "2.0",
//	    "id": 1,
//	    "method": "getAccountInfo",
//	    "params": [
//	      "vines1vzrYbzLMRdu58ou5XTby4qAqVRLmqo36NKPTg",
//	      {
//	        "encoding": "base58"
//	      }
//	    ]
//	  }'

func main() {
	// alcht_3IcUtPhbQsBVY57gH7e2ajJzInt2T6
	url := "https://solana-mainnet.g.alchemy.com/v2/alcht_3IcUtPhbQsBVY57gH7e2ajJzInt2T6"

	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"getLatestBlockhash\",\"params\":[\"CYbD9RaToYMtWKA7QZyoLahnHdWq553Vm62Lh6qWtuxq\"]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
