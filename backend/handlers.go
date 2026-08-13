package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// ProductHandler ผูก store เข้ากับฟังก์ชันจัดการ HTTP
type ProductHandler struct {
	store *ProductStore
}

// writeJSON ส่ง response กลับเป็น JSON พร้อมกำหนด status code
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// writeError ส่งข้อความ error กลับเป็น JSON
func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

// idFromPath ดึงค่า {id} จาก URL แล้วแปลงเป็นตัวเลข
func idFromPath(r *http.Request) (int, error) {
	return strconv.Atoi(r.PathValue("id"))
}

// list = GET /products (อ่านทั้งหมด)
func (h *ProductHandler) list(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.store.List())
}

// get = GET /products/{id} (อ่าน 1 ชิ้น)
func (h *ProductHandler) get(w http.ResponseWriter, r *http.Request) {
	id, err := idFromPath(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id ต้องเป็นตัวเลข")
		return
	}
	p, err := h.store.Get(id)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "ไม่พบสินค้า")
		return
	}
	writeJSON(w, http.StatusOK, p)
}

// create = POST /products (เพิ่มใหม่)
func (h *ProductHandler) create(w http.ResponseWriter, r *http.Request) {
	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeError(w, http.StatusBadRequest, "ข้อมูล JSON ไม่ถูกต้อง")
		return
	}
	if p.Name == "" {
		writeError(w, http.StatusBadRequest, "ต้องระบุชื่อสินค้า (name)")
		return
	}
	created := h.store.Create(p)
	writeJSON(w, http.StatusCreated, created)
}

// update = PUT /products/{id} (แก้ไข)
func (h *ProductHandler) update(w http.ResponseWriter, r *http.Request) {
	id, err := idFromPath(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id ต้องเป็นตัวเลข")
		return
	}
	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeError(w, http.StatusBadRequest, "ข้อมูล JSON ไม่ถูกต้อง")
		return
	}
	updated, err := h.store.Update(id, p)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "ไม่พบสินค้า")
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

// delete = DELETE /products/{id} (ลบ)
func (h *ProductHandler) delete(w http.ResponseWriter, r *http.Request) {
	id, err := idFromPath(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id ต้องเป็นตัวเลข")
		return
	}
	if err := h.store.Delete(id); errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "ไม่พบสินค้า")
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204 = สำเร็จ ไม่มีเนื้อหาส่งกลับ
}
