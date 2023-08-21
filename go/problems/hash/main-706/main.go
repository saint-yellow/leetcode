// LeetCode 主站 Problem Nr. 706: 设计哈希映射

package main

const base = 769

type entry struct {
	key, value int
	next       *entry
}

type MyHashMap struct {
	hm []*entry
}

func Constructor() MyHashMap {
	hm := make([]*entry, base)
	for i := range hm {
		hm[i] = new(entry)
	}
	return MyHashMap{
		hm: hm,
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % base
}

func (this *MyHashMap) Put(key int, value int) {
	h := this.hash(key)
	var p *entry
	for p = this.hm[h]; p.next != nil; p = p.next {
		if p.next.key == key {
			p.next.value = value
			return
		}
	}
	p.next = &entry{key, value, nil}
}

func (this *MyHashMap) Get(key int) int {
	h := this.hash(key)
	for p := this.hm[h]; p.next != nil; p = p.next {
		if p.next.key == key {
			return p.next.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	h := this.hash(key)
	for p := this.hm[h]; p.next != nil; p = p.next {
		if p.next.key == key {
			p.next = p.next.next
			return
		}
	}
}
