package main

import (
	"errors"
	"testing"
)

func TestCreateAndGet(t *testing.T) {
	store := NewProductStore()

	created := store.Create(Product{Name: "กาแฟ", Price: 45, Stock: 10})
	if created.ID != 1 {
		t.Fatalf("อยากได้ id = 1 แต่ได้ %d", created.ID)
	}

	got, err := store.Get(1)
	if err != nil {
		t.Fatalf("ไม่ควร error แต่ได้: %v", err)
	}
	if got.Name != "กาแฟ" {
		t.Errorf("อยากได้ชื่อ 'กาแฟ' แต่ได้ '%s'", got.Name)
	}
}

func TestGetNotFound(t *testing.T) {
	store := NewProductStore()

	_, err := store.Get(999)
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("อยากได้ ErrNotFound แต่ได้: %v", err)
	}
}

func TestDelete(t *testing.T) {
	store := NewProductStore()
	store.Create(Product{Name: "ชาเขียว"})

	if err := store.Delete(1); err != nil {
		t.Fatalf("ลบไม่ควร error แต่ได้: %v", err)
	}
	if _, err := store.Get(1); !errors.Is(err, ErrNotFound) {
		t.Error("ลบแล้วแต่ยังหาสินค้าเจอ")
	}
}
