package etcd


// etcd type policy:
// first compartment means the type of data.
// it starts with / which means it folows described policy.
// second compartment means unique id
// third and more compartment means the field of data.
// if any fourth and more compartment exists,
// it is a record of third compartment which is an object



// /worker/ID
// /worker/ID/platform
// /worker/ID/status
type Worker struct {
	ID       string `etcd:"id"`       // uuid format
	Platform string `etcd:"platform"` // kis, polygon, buycycle, ...
	Status   string `etcd:"status"`   // active, waiting, dead
}

// /product/ID
// /product/ID/platform
// /product/ID/symbol
// /product/ID/worker
type Product struct {
	ID       string `etcd:"id"`       // product_type.name.region
	Platform string `etcd:"platform"` // kis, polygon, buycycle, ...
	Symbol   string `etcd:"symbol"`   // identifier inside platform
	Worker 	 string `etcd:"worker"`   // uuid format
	Status   string `etcd:"status"`   // onsubscribe, notsubscribed
}