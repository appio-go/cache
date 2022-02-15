# go-cache
Allows you to cache some data, without any extra storages


## Usage
Add CACHE_LIFETIME="5" (minutes for cache to life) to project .env file


```
package main
import "github.com/appio-go/cache"
import "github.com/joho/godotenv"


func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	rand.Seed(time.Now().UnixNano())
}


func main(){
	var cache cache.Singleton
	
	//get your instance
	cache = GetInstance()
    
	data, found := cache.Get("Some cache key")

	if found {
		//use your data here
	} else {
	    //Marshal some data
	    marshaledData, _ := json.Marshal(u)
	    
	    //add it to cache
		cache.Set("Some cache key", marshaledData)
	}

}
```
