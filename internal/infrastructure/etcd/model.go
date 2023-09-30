package etcd



// /worker/ID
// /worker/ID/platform
type Worker struct {
	ID       string `etcd:"id"`
	Platform string `etcd:"platform"`
}

// /product/ID
// /product/ID/platform
// /product/ID/symbol
// /product/ID/worker
type Product struct {
	ID       string `etcd:"id"`
	Platform string `etcd:"platform"`
	Symbol   string `etcd:"symbol"`
	Worker 	 string `etcd:"worker"`
}