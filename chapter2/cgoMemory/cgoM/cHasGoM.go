package cgoM

import "sync"

type ObjectId int32

func (id ObjectId) IsNil() bool {
	return id == 0
}

func (id ObjectId) Get() interface{} {
	refs.Lock()
	defer refs.Unlock()

	return refs.objS[id]
}

func (id *ObjectId) Free() interface{} {
	refs.Lock()
	defer refs.Unlock()

	obj := refs.objS[*id]
	delete(refs.objS, *id)
	*id = 0

	return obj
}

var refs struct {
	sync.Mutex
	objS map[ObjectId]interface{}
	next ObjectId
}

func NewObjectId(obj interface{}) ObjectId {
	refs.Lock()
	defer refs.Unlock()

	id := refs.next
	refs.next++
	refs.objS[id] = obj
	return id
}

func init() {
	refs.Lock()
	defer refs.Unlock()

	refs.objS = make(map[ObjectId]interface{})
	refs.next = 1000
}
