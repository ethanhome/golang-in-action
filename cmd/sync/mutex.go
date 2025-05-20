package main

import "sync"

/*
For any sync.Mutex or sync.RWMutex variable l and n < m,
call n of l.Unlock() is synchronized before call m of l.Lock() returns.

For any call to l.RLock on a sync.RWMutex variable l,
there is an n such that the nth call to l.Unlock is synchronized before the return from l.RLock,
and the matching call to l.RUnlock is synchronized before the return from call n+1 to l.Lock.

A successful call to l.TryLock (or l.TryRLock) is equivalent to a call to l.Lock (or l.RLock).
An unsuccessful call has no synchronizing effect at all.
As far as the memory model is concerned, l.TryLock (or l.TryRLock) may be considered to be able to return false even when the mutex l is unlocked.
*/
var l sync.Mutex
var a string

func f() {
	a = "hello, world"
	l.Unlock()
	//l.Unlock() will lead deadlock
}

func main() {
	l.Lock()
	go f()
	l.Lock()
	print(a)
}
