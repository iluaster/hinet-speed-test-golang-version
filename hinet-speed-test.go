// Hinet speed test in Golang 
// v1.0

package main

import (
        "io/ioutil"
        "log"
        "net/http"
      	"fmt"
	      "time"
)

func main() {
        client := &http.Client{}
        
        // Hinet speed test file url. 
        req, err := http.NewRequest("GET", "http://tpdb.speed2.hinet.net/test_060m.zip", nil)
        if err != nil {
                log.Fatalln(err)
        }
        
        // Set HTTP Request User-Agent & Referer header
        req.Header.Set("User-Agent", "Mozilla/5.0")
        req.Header.Set("Referer", "http://speed.hinet.net/index_test01.htm")
      	
      	// Set Current Time
      	t1 := time.Now()
	      fmt.Printf("Start: %s\n",t1)
	      
	      // Send HTTP Request
        resp, err := client.Do(req)
        if err != nil {
                log.Fatalln(err)
        }

        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                log.Fatalln(err)
        }
        
        // Set Finished Time
	      t2 := time.Now()
      	fmt.Printf("End  : %s\n",t2)
      	
      	// Conver t2.Sub(t1) value to seconds unit.
	      elapsed_second_time := t2.Sub(t1)/1000000000
	      
      	fmt.Printf("The download took %v to run.\n",t2.Sub(t1))
      	
      	// The file size is 62914560 bytes
	      fmt.Printf("The download rate is %d KB/s.\n",62914560/1024/elapsed_second_time)

        // Write packet into file and set file permission.
      	ioutil.WriteFile("test_060m.zip",body,0644)

}
