package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB
var tmpl *template.Template

type Department struct {
	ID       int          `json:"id"`
	ParentID *int         `json:"parent_id"`
	Name     string       `json:"name"`
	Code     string       `json:"code"`
	Level    int          `json:"level"`
	Children []Department `json:"children"`
}

func init() {
	tmpl = template.Must(template.ParseFiles("templates/index.html", "templates/department_tree.html"))
}

func main() {

	//Read ENV
	dir, err1 := os.Getwd()
	if err1 != nil {
		log.Fatal(err1)
	}
	environmentPath := filepath.Join(dir, ".env")
	err1 = godotenv.Load(environmentPath)

	if err1 != nil {
		log.Fatal(err1)
	}
	//Read ENV

	// Update these values with your local MySQL credentials
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)
	var err error
	db, err = sql.Open("mysql", dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/departments/create", createHandler)
	http.HandleFunc("/departments/read", readHandler)
	http.HandleFunc("/departments/update", updateHandler)
	http.HandleFunc("/departments/delete", deleteHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, parent_id, name, code, level FROM departments ORDER BY level, id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	departmentMap := make(map[int][]Department)
	var rootDepartments []Department

	for rows.Next() {
		var dept Department
		err := rows.Scan(&dept.ID, &dept.ParentID, &dept.Name, &dept.Code, &dept.Level)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if dept.ParentID == nil {
			rootDepartments = append(rootDepartments, dept)
		} else {
			departmentMap[*dept.ParentID] = append(departmentMap[*dept.ParentID], dept)
		}
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rootDepartments = buildDepartmentHierarchy(rootDepartments)

	data := TemplateData{
		RootDepartments: rootDepartments,
		EditDepartment:  Department{}, // Empty struct if on the index page
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	parentID, _ := strconv.Atoi(r.FormValue("parent_id"))
	name := r.FormValue("name")
	code := r.FormValue("code")
	level, _ := strconv.Atoi(r.FormValue("level"))

	_, err := db.Exec("INSERT INTO departments (parent_id, name, code, level) VALUES (?, ?, ?, ?)", parentID, name, code, level)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

type TemplateData struct {
	RootDepartments []Department
	EditDepartment  Department
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var dept Department

	// Query to fetch the department details based on ID
	err := db.QueryRow("SELECT id, parent_id, name, code, level FROM departments WHERE id = ?", id).
		Scan(&dept.ID, &dept.ParentID, &dept.Name, &dept.Code, &dept.Level)
	if err != nil {
		http.Error(w, "Department not found", http.StatusNotFound)
		return
	}

	// Fetch all departments for the tree and table
	allDepartments, err := fetchAllDepartments()
	if err != nil {
		http.Error(w, "Unable to fetch departments", http.StatusInternalServerError)
		return
	}

	// Create the hierarchical structure for the tree view
	rootDepartments := buildDepartmentHierarchy(allDepartments)

	// Create the template data
	data := TemplateData{
		RootDepartments: rootDepartments,
		EditDepartment:  dept, // Pass the department to be edited
	}

	// Render the index template with the department to populate the form
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Helper function to fetch all departments
func fetchAllDepartments() ([]Department, error) {
	rows, err := db.Query("SELECT id, parent_id, name, code, level FROM departments ORDER BY level, id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []Department
	for rows.Next() {
		var dept Department
		err := rows.Scan(&dept.ID, &dept.ParentID, &dept.Name, &dept.Code, &dept.Level)
		if err != nil {
			return nil, err
		}
		departments = append(departments, dept)
	}
	return departments, nil
}

// Helper function to build department hierarchy
func buildDepartmentHierarchy(departments []Department) []Department {
	departmentMap := make(map[int][]Department)
	var rootDepartments []Department

	for _, dept := range departments {
		if dept.ParentID == nil {
			rootDepartments = append(rootDepartments, dept)
		} else {
			departmentMap[*dept.ParentID] = append(departmentMap[*dept.ParentID], dept)
		}
	}

	var buildHierarchy func([]Department) []Department
	buildHierarchy = func(departments []Department) []Department {
		for i := range departments {
			if children, ok := departmentMap[departments[i].ID]; ok {
				departments[i].Children = buildHierarchy(children)
			}
		}
		return departments
	}

	return buildHierarchy(rootDepartments)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	parentID, _ := strconv.Atoi(r.FormValue("parent_id"))
	name := r.FormValue("name")
	code := r.FormValue("code")
	level, _ := strconv.Atoi(r.FormValue("level"))

	_, err := db.Exec("UPDATE departments SET parent_id = ?, name = ?, code = ?, level = ? WHERE id = ?", parentID, name, code, level, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := db.Exec("DELETE FROM departments WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
