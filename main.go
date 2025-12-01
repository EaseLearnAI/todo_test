package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// Todo 结构体
type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	DueAt     *string   `json:"dueAt"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Response 响应结构
type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

var (
	dataDir  = "data"
	dataFile = filepath.Join(dataDir, "todos.json")
)

// 确保数据文件存在
func ensureDataFile() error {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return err
	}
	
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return ioutil.WriteFile(dataFile, []byte("[]"), 0644)
	}
	return nil
}

// 读取所有 Todos
func readTodos() ([]Todo, error) {
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return []Todo{}, nil
	}
	
	var todos []Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return []Todo{}, nil
	}
	
	return todos, nil
}

// 写入 Todos
func writeTodos(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	
	tmpFile := dataFile + ".tmp"
	if err := ioutil.WriteFile(tmpFile, data, 0644); err != nil {
		return err
	}
	
	return os.Rename(tmpFile, dataFile)
}

// 验证 ISO 日期格式
func isValidISODate(dateStr string) bool {
	if dateStr == "" {
		return false
	}
	
	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04:05",
		"2006-01-02T15:04",
	}
	
	for _, layout := range layouts {
		if _, err := time.Parse(layout, dateStr); err == nil {
			return true
		}
	}
	return false
}

// 查找 Todo 索引
func findTodoIndex(todos []Todo, id string) int {
	for i, todo := range todos {
		if todo.ID == id {
			return i
		}
	}
	return -1
}

// API Handlers

// 获取所有 Todos
func getTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := readTodos()
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, Response{Error: "读取数据失败"})
		return
	}
	
	respondJSON(w, http.StatusOK, Response{Data: todos})
}

// 创建新 Todo
func createTodo(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string  `json:"title"`
		DueAt *string `json:"dueAt"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondJSON(w, http.StatusBadRequest, Response{Error: "请求体错误"})
		return
	}
	
	title := strings.TrimSpace(input.Title)
	if title == "" {
		respondJSON(w, http.StatusBadRequest, Response{Error: "标题不能为空"})
		return
	}
	
	if input.DueAt != nil && *input.DueAt != "" && !isValidISODate(*input.DueAt) {
		respondJSON(w, http.StatusBadRequest, Response{Error: "时间格式不正确"})
		return
	}
	
	todos, _ := readTodos()
	now := time.Now()
	
	newTodo := Todo{
		ID:        uuid.New().String(),
		Title:     title,
		DueAt:     input.DueAt,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
	
	todos = append(todos, newTodo)
	
	if err := writeTodos(todos); err != nil {
		respondJSON(w, http.StatusInternalServerError, Response{Error: "保存失败"})
		return
	}
	
	respondJSON(w, http.StatusCreated, Response{Data: newTodo})
}

// 更新 Todo
func updateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["id"]
	
	var input struct {
		Title     *string `json:"title"`
		DueAt     *string `json:"dueAt"`
		Completed *bool   `json:"completed"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondJSON(w, http.StatusBadRequest, Response{Error: "请求体错误"})
		return
	}
	
	todos, _ := readTodos()
	idx := findTodoIndex(todos, todoID)
	
	if idx == -1 {
		respondJSON(w, http.StatusNotFound, Response{Error: "未找到Todo"})
		return
	}
	
	todo := todos[idx]
	
	if input.Title != nil {
		title := strings.TrimSpace(*input.Title)
		if title == "" {
			respondJSON(w, http.StatusBadRequest, Response{Error: "标题不能为空"})
			return
		}
		todo.Title = title
	}
	
	if input.DueAt != nil {
		if *input.DueAt != "" && !isValidISODate(*input.DueAt) {
			respondJSON(w, http.StatusBadRequest, Response{Error: "时间格式不正确"})
			return
		}
		todo.DueAt = input.DueAt
	}
	
	if input.Completed != nil {
		todo.Completed = *input.Completed
	}
	
	todo.UpdatedAt = time.Now()
	todos[idx] = todo
	
	if err := writeTodos(todos); err != nil {
		respondJSON(w, http.StatusInternalServerError, Response{Error: "保存失败"})
		return
	}
	
	respondJSON(w, http.StatusOK, Response{Data: todo})
}

// 切换 Todo 完成状态
func toggleTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["id"]
	
	todos, _ := readTodos()
	idx := findTodoIndex(todos, todoID)
	
	if idx == -1 {
		respondJSON(w, http.StatusNotFound, Response{Error: "未找到Todo"})
		return
	}
	
	todos[idx].Completed = !todos[idx].Completed
	todos[idx].UpdatedAt = time.Now()
	
	if err := writeTodos(todos); err != nil {
		respondJSON(w, http.StatusInternalServerError, Response{Error: "保存失败"})
		return
	}
	
	respondJSON(w, http.StatusOK, Response{Data: todos[idx]})
}

// 删除 Todo
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["id"]
	
	todos, _ := readTodos()
	idx := findTodoIndex(todos, todoID)
	
	if idx == -1 {
		respondJSON(w, http.StatusNotFound, Response{Error: "未找到Todo"})
		return
	}
	
	removed := todos[idx]
	todos = append(todos[:idx], todos[idx+1:]...)
	
	if err := writeTodos(todos); err != nil {
		respondJSON(w, http.StatusInternalServerError, Response{Error: "保存失败"})
		return
	}
	
	respondJSON(w, http.StatusOK, Response{Data: removed})
}

// 响应 JSON
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// SPA 路由处理
func spaHandler(staticPath string, indexPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(staticPath, r.URL.Path)
		
		// 检查文件是否存在
		if _, err := os.Stat(path); os.IsNotExist(err) {
			// 文件不存在，返回 index.html
			http.ServeFile(w, r, filepath.Join(staticPath, indexPath))
			return
		}
		
		// 文件存在，直接提供
		http.FileServer(http.Dir(staticPath)).ServeHTTP(w, r)
	}
}

func main() {
	// 加载环境变量
	godotenv.Load()
	
	// 确保数据文件存在
	if err := ensureDataFile(); err != nil {
		log.Fatal("无法初始化数据文件:", err)
	}
	
	// 创建路由
	router := mux.NewRouter()
	
	// API 路由
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/todos", getTodos).Methods("GET")
	api.HandleFunc("/todos", createTodo).Methods("POST")
	api.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	api.HandleFunc("/todos/{id}/toggle", toggleTodo).Methods("POST")
	api.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")
	
	// 静态文件和 SPA 路由
	staticPath := "dist"
	router.PathPrefix("/").HandlerFunc(spaHandler(staticPath, "index.html"))
	
	// CORS 配置
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	
	handler := c.Handler(router)
	
	// 获取端口
	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}
	
	addr := "0.0.0.0:" + port
	
	fmt.Printf("🚀 Server running at http://localhost:%s\n", port)
	fmt.Println("📁 Serving static files from:", staticPath)
	fmt.Println("💾 Data file:", dataFile)
	
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
