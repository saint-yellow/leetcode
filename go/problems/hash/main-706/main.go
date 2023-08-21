// LeetCode 主站 Problem Nr. 706: 设计哈希映射

package main

import "container/list"

const base = 769

type MyHashMap struct {
	hm []list.List
}

type entry struct {
	key, value int
}

func Constructor() MyHashMap {
	return MyHashMap{
		hm: make([]list.List, base),
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % base
}

func (this *MyHashMap) Put(key int, value int) {
	h := this.hash(key)
	for ele := this.hm[h].Front(); ele != nil; ele = ele.Next() {
		if etr := ele.Value.(entry); etr.key == key {
			ele.Value = entry{key, value}
			return
		}
	}
	this.hm[h].PushBack(entry{key, value})
}

func (this *MyHashMap) Get(key int) int {
	h := this.hash(key)
	for ele := this.hm[h].Front(); ele != nil; ele = ele.Next() {
		if etr := ele.Value.(entry); etr.key == key {
			return etr.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	h := this.hash(key)
	for ele := this.hm[h].Front(); ele != nil; ele = ele.Next() {
		if etr := ele.Value.(entry); etr.key == key {
			this.hm[h].Remove(ele)
		}
	}
}
