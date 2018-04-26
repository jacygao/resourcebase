/*
Package resourcebase contains a very simple pooling algorithm.
It's a perfect addon to your worker pool where it sets the maximum 
number of concurrent processes so your worker don't go out of control.
A perfect use case is controlling the max number of your database connections.*/
package resourcebase

import (
	"log"
)

// ResourceBase Manages the pooling of access to a fixed number of resources.
type ResourceBase interface {
	Take()
	Return()
}

type resourceBase struct {
	resourceName string
	pool       chan bool
}

// Take retrieves any available resource from pool.
// If no resource is available Take will hang and wait.
func (r *resourceBase) Take() {
	select {
	case <-r.pool:
		return
	default:
		log.Printf("[WARNING] %s TakeLease blocked", r.resourceName)
		<-r.pool
	}
}

// Return returns a resource back to the pool.
func (r *resourceBase) Return() {
	r.pool <- true
}

// NewResourceBase creates a new instance of resourcebase.
// ResourceName is a reference for the resources stored.
// Size defines the maximum number of concurrent resource to be taken from the pool.
func NewResourceBase(resourceName string, size int) ResourceBase {
	r := resourceBase{
		resourceName: resourceName,
		pool:       make(chan bool, size),
	}
	for i := 0; i < size; i++ {
		r.pool <- true
	}
	return &r
}
