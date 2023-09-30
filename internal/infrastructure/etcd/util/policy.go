package etcdutil

// ETCD UTIL POLICY DOCUMENT



// Policy for etcd object:

// First compartment means the type of data.
// It starts with / which means it folows described policy.
// Second compartment means unique id
// Third and more compartment means the field of data.
// If any fourth and more compartment exists, it is a record of third compartment which is an object

// This package provides a Marshal and Unmarshal function to handle object type on etcd.
// Every struct record that needs to be recognized by this package should have a tag named "etcd".



// Policy for etcd semaphore
// Key-value pair for specific usecase does not starts with /, should starts with its function directly.
// Semaphore starts with semaphore
// Second compartment means semaphore name