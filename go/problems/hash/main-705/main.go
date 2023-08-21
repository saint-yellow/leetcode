// LeetCode 主站 Problem Nr. 705: 设计哈希集合

package main

const base = 167

type entry struct {
	key  int
	next *entry
}

type MyHashSet struct {
	hs []*entry
}

func Constructor() MyHashSet {
	hs := make([]*entry, base)
	for i := range hs {
		hs[i] = new(entry)
	}
	return MyHashSet{
		hs: hs,
	}
}

func (this *MyHashSet) hash(key int) int {
	return key % base
}

func (this *MyHashSet) Add(key int) {
	h := this.hash(key)
	p := this.hs[h]
	for p = this.hs[h]; p.next != nil; p = p.next {
		if p.next.key == key {
			return
		}
	}
	p.next = &entry{key, nil}
}

func (this *MyHashSet) Remove(key int) {
	h := this.hash(key)
	for p := this.hs[h]; p.next != nil; p = p.next {
		if p.next.key == key {
			p.next = p.next.next
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	h := this.hash(key)
	for p := this.hs[h].next; p != nil; p = p.next {
		if p.key == key {
			return true
		}
	}
	return false
}
