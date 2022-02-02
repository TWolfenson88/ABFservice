package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type BucketConfig struct {
	LogLimit int
	PassLimit int
	IPLimit int
}

type Login struct {
	limiter *rate.Limiter
	lastUsage time.Time
}

type Password struct {
	limiter *rate.Limiter
	lastUsage time.Time
}

type IP struct {
	limiter *rate.Limiter
	lastUsage time.Time
}

var logins = make(map[string]*Login)
var passes = make(map[string]*Password)
var IPs = make(map[string]*IP)
var mu sync.Mutex

func(bc *BucketConfig) getLogin(login string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, ok := logins[login]
	if !ok {
		limiter := rate.NewLimiter(rate.Limit(bc.LogLimit/60), bc.LogLimit)
		logins[login] = &Login{limiter, time.Now()}
		return limiter
	}

	v.lastUsage = time.Now()
	return v.limiter
}

func(bc *BucketConfig) getPass(pass string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, ok := passes[pass]
	if !ok {
		limiter := rate.NewLimiter(rate.Limit(bc.PassLimit/60), bc.PassLimit)
		passes[pass] = &Password{limiter, time.Now()}
		return limiter
	}

	v.lastUsage = time.Now()
	return v.limiter
}

func(bc *BucketConfig) getIP(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, ok := IPs[ip]
	if !ok {
		limiter := rate.NewLimiter(rate.Limit(bc.IPLimit/60), bc.IPLimit)
		IPs[ip] = &IP{limiter, time.Now()}
		return limiter
	}

	v.lastUsage = time.Now()
	return v.limiter
}

func(bc *BucketConfig) LoginBucket(login string) bool {
	limiter := bc.getLogin(login)

	return limiter.Allow()
}

func(bc *BucketConfig) PassBucket(pass string) bool {
	limiter := bc.getPass(pass)

	return limiter.Allow()
}

func(bc *BucketConfig) IPBucket(ip string) bool {
	limiter := bc.getIP(ip)

	return limiter.Allow()
}

func Limit(w http.ResponseWriter, r *http.Request) {

		//HERE we should check in white and black list and if not found

	body, _ := ioutil.ReadAll(r.Body)

	fmt.Println("ALERT: ", r.RemoteAddr, string(body))

		// Get the IP address for the current user.
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		fmt.Println("IP ADDR", ip)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Call the getVisitor function to retreive the rate limiter for
		// the current user.
		//limiter := getVisitor(ip)
		//if limiter.Allow() == false {
		//	http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
		//	return
		//}

		w.Write([]byte("OK"))
	}