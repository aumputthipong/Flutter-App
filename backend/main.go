package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewProductStore()
	// ใส่ข้อมูลตัวอย่างไว้ 2 ชิ้น เพื่อให้ทดสอบได้ทันทีตอนเปิดเซิร์ฟเวอร์
	store.Create(Product{Name: "กาแฟ", Price: 45, Stock: 100})
	store.Create(Product{Name: "ชาเขียว", Price: 40, Stock: 80})

	h := &ProductHandler{store: store}

	// ServeMux ของ Go 1.22+ ระบุทั้ง method และ path ได้ในบรรทัดเดียว
	mux := http.NewServeMux()
	mux.HandleFunc("GET /products", h.list)
	mux.HandleFunc("POST /products", h.create)
	mux.HandleFunc("GET /products/{id}", h.get)
	mux.HandleFunc("PUT /products/{id}", h.update)
	mux.HandleFunc("DELETE /products/{id}", h.delete)

	addr := ":8080"
	log.Printf("เซิร์ฟเวอร์กำลังทำงานที่ http://localhost%s", addr)
	if err := http.ListenAndServe(addr, corsMiddleware(mux)); err != nil {
		log.Fatal(err)
	}
}

// corsMiddleware อนุญาตให้เว็บจากที่อื่น (เช่น Flutter web) เรียก API นี้ได้
// ไม่งั้นเบราว์เซอร์จะบล็อกด้วยกฎ CORS
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// เบราว์เซอร์จะยิง OPTIONS มาถามก่อน (preflight) — ตอบ 204 ผ่านได้เลย
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
