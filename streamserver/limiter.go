package main

import "log"

type ConnLimiter struct {
	concurrentConn int
	bucket         chan struct{}
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan struct{}, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}
	cl.bucket <- struct{}{}
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	<-cl.bucket
	log.Printf("Connection Released.")
}
