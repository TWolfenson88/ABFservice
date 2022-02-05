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

//var logins = make(map[string]*Login)
//var passes = make(map[string]*Password)
//var IPs = make(map[string]*IP)
//var mu sync.Mutex

type BucketStorage struct {
		Logins map[string]*Login
		Passes map[string]*Password
		IPs    map[string]*IP
		mu     sync.Mutex
	}

func (bs *BucketStorage) GetLogin(login string, n int) *rate.Limiter {
	bs.mu.Lock()
	defer bs.mu.Unlock()

		v, ok := bs.Logins[login]
		if !ok {
			limiter := rate.NewLimiter(rate.Limit(n/60), n)
			bs.Logins[login] = &Login{limiter, time.Now()}
			return limiter
		}

		v.lastUsage = time.Now()
		return v.limiter
}

func (bs *BucketStorage) GetPass(pass string, m int) *rate.Limiter {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	v, ok := bs.Passes[pass]
	if !ok {
		limiter := rate.NewLimiter(rate.Limit(m/60), m)
		bs.Passes[pass] = &Password{limiter, time.Now()}
		return limiter
	}

	v.lastUsage = time.Now()
	return v.limiter
}

func (bs *BucketStorage) GetIP(ip string, k int) *rate.Limiter {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	v, ok := bs.IPs[ip]
	if !ok {
		limiter := rate.NewLimiter(rate.Limit(k/60), k)
		bs.IPs[ip] = &IP{limiter, time.Now()}
		return limiter
	}

	v.lastUsage = time.Now()
	return v.limiter
}

func (bs *BucketStorage) LoginBucket(login string, n int) bool {
	limiter := bs.GetLogin(login, n)

	return limiter.Allow()
}

func (bs *BucketStorage) PassBucket(pass string, m int) bool {
	limiter := bs.GetPass(pass, m)

	return limiter.Allow()
}

func (bs *BucketStorage) IPBucket(ip string, k int) bool {
	limiter := bs.GetIP(ip, k)

	return limiter.Allow()
}

type RequestParams struct {
	Login    string
	Password string
	IPaddr   string
}

func(bs *BucketStorage) Limit(params RequestParams, bc BucketConfig) (bool, error) {

	//TODO:NEED TO CHECK W/B LISTS

	if !bs.LoginBucket(params.Login, bc.LogLimit) {
		return false, errors.New("too much tries for login")
	}
	if !bs.PassBucket(params.Login, bc.PassLimit) {
		return false, errors.New("too much tries for pass")
	}
	if !bs.IPBucket(params.Login, bc.IPLimit) {
		return false, errors.New("too much tries for ip")
	}

	return true, nil
}

func(bs *BucketStorage) ResetBuckets(params RequestParams) {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	if _, ok := bs.Logins[params.Login]; ok {
		delete(bs.Logins, params.Login)
	}
	if _, ok := bs.IPs[params.IPaddr]; ok {
		delete(bs.IPs, params.IPaddr)
	}

}

func(bs *BucketStorage) CleanupBuckets() {
	for {
		time.Sleep(time.Minute)

		bs.mu.Lock()
		for log, v := range bs.Logins {
			if time.Since(v.lastUsage) > 3*time.Minute {
				delete(bs.Logins, log)
			}
		}
		for pass, v := range bs.Passes {
			if time.Since(v.lastUsage) > 3*time.Minute {
				delete(bs.Passes, pass)
			}
		}
		for ip, v := range bs.IPs {
			if time.Since(v.lastUsage) > 3*time.Minute {
				delete(bs.IPs, ip)
			}
		}
		bs.mu.Unlock()
	}
}
