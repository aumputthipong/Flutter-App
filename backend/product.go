package main

import (
	"errors"
	"sort"
	"sync"
)

// Product คือข้อมูลสินค้า 1 ชิ้น
// แท็ก `json:"..."` บอกว่าเวลาแปลงเป็น/จาก JSON ให้ใช้ชื่อคีย์แบบไหน
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// ErrNotFound ใช้แจ้งว่าหาสินค้าตาม id ไม่เจอ
var ErrNotFound = errors.New("ไม่พบสินค้า")

// ProductStore เก็บสินค้าไว้ในหน่วยความจำ (RAM)
// ใช้ Mutex ล็อกกันไม่ให้หลาย request แก้ข้อมูลชนกันพร้อมกัน
type ProductStore struct {
	mu     sync.Mutex
	items  map[int]Product
	nextID int
}

func NewProductStore() *ProductStore {
	return &ProductStore{
		items:  make(map[int]Product),
		nextID: 1,
	}
}

// List คืนสินค้าทั้งหมด (เรียงตาม id)
func (s *ProductStore) List() []Product {
	s.mu.Lock()
	defer s.mu.Unlock()

	list := make([]Product, 0, len(s.items))
	for _, p := range s.items {
		list = append(list, p)
	}
	sort.Slice(list, func(i, j int) bool { return list[i].ID < list[j].ID })
	return list
}

// Get คืนสินค้าตาม id
func (s *ProductStore) Get(id int) (Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	p, ok := s.items[id]
	if !ok {
		return Product{}, ErrNotFound
	}
	return p, nil
}

// Create เพิ่มสินค้าใหม่ แล้วคืนสินค้าพร้อม id ที่ระบบสร้างให้
func (s *ProductStore) Create(p Product) Product {
	s.mu.Lock()
	defer s.mu.Unlock()

	p.ID = s.nextID
	s.nextID++
	s.items[p.ID] = p
	return p
}

// Update แก้ไขสินค้าตาม id
func (s *ProductStore) Update(id int, p Product) (Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.items[id]; !ok {
		return Product{}, ErrNotFound
	}
	p.ID = id // กัน id ในตัว body ให้ตรงกับ id ใน URL เสมอ
	s.items[id] = p
	return p, nil
}

// Delete ลบสินค้าตาม id
func (s *ProductStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.items[id]; !ok {
		return ErrNotFound
	}
	delete(s.items, id)
	return nil
}
