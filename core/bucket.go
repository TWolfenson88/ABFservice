package core

import (
	"errors"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type BucketConfig struct {
	LogLimit  int
	PassLimit int
	IPLimit   int
}

type Login struct {
	limiter   *rate.Limiter
	lastUsage time.Time
}

type Password struct {
	limiter   *rate.Limiter
	lastUsage time.Time
}

type IP struct {
	limiter   *rate.Limiter
	lastUsage time.Time
}

var logins = make(map[string]*Login)
var passes = make(map[string]*Password)
var IPs = make(map[string]*IP)
var mu sync.Mutex

func (bc *BucketConfig) getLogin(login string) *rate.Limiter {
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

func (bc *BucketConfig) getPass(pass string) *rate.Limiter {
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

func (bc *BucketConfig) getIP(ip string) *rate.Limiter {
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

func (bc *BucketConfig) LoginBucket(login string) bool {
	limiter := bc.getLogin(login)

	return limiter.Allow()
}

func (bc *BucketConfig) PassBucket(pass string) bool {
	limiter := bc.getPass(pass)

	return limiter.Allow()
}

func (bc *BucketConfig) IPBucket(ip string) bool {
	limiter := bc.getIP(ip)

	return limiter.Allow()
}

type RequestParams struct {
	Login    string
	Password string
	IPaddr   string
}

func Limit(params RequestParams) (bool, error) {

	//TODO: TEMP SHIT, REMOVE BEFORE FLIGHT
	bc := BucketConfig{
		LogLimit:  10,
		PassLimit: 100,
		IPLimit:   1000,
	}

	//TODO:NEED TO CHECK W/B LISTS

	if !bc.LoginBucket(params.Login) {
		return false, errors.New("too much tries for login")
	}
	if !bc.PassBucket(params.Login) {
		return false, errors.New("too much tries for pass")
	}
	if !bc.IPBucket(params.Login) {
		return false, errors.New("too much tries for ip")
	}

	return true, nil
}

func ResetBuckets(params RequestParams)  {

	mu.Lock()
	defer mu.Unlock()

	if _, ok := logins[params.Login]; ok {
		delete(logins, params.Login)
	}
	if _, ok := IPs[params.IPaddr]; ok {
		delete(IPs, params.IPaddr)
	}

}

func cleanupBuckets() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for log, v := range logins {
			if time.Since(v.lastUsage) > 3*time.Minute {
				delete(logins, log)
			}
		}
		for pass, v := range passes {
			if time.Since(v.lastUsage) > 3*time.Minute {
				delete(passes, pass)
			}
		}
		for ip, v := range IPs {
			if time.Since(v.lastUsage) > 3*time.Minute {
				delete(IPs, ip)
			}
		}
		mu.Unlock()
	}
}
