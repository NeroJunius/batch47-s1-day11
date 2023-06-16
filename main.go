package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type Projects struct {
	Title        string
	Author       string
	detailDate   string
	descProjects string
	nodeJS       string
	reactJS      string
	nextJS       string
	typeScript   string
}

var dataProject = []Projects{
	{
		Title:        "Ameno ameno latire Latiremo Dori me",
		Author:       "Nafiisan N. Achmad",
		detailDate:   "durasi : 4 bulan",
		descProjects: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin quis risus ut mi euismod sodales. Mauris id quam ut massa sodales faucibus consectetur sit amet dolor. ",
		nodeJS:       "nodejs",
		reactJS:      "reactjs",
		nextJS:       "nextjs",
		typeScript:   "typescript",
	},

	{
		Title:        "Ameno ameno, ameen",
		Author:       "Nero002",
		detailDate:   "durasi : 11 bulan",
		descProjects: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin quis risus ut mi euismod sodales. Mauris id quam ut massa sodales faucibus consectetur sit amet dolor. ",
		nodeJS:       "nodejs",
		reactJS:      "reactjs",
		nextJS:       "nextjs",
		typeScript:   "typescript",
	},

	{
		Title:        "latire Latiremo Dori me",
		Author:       "NeroJunius",
		detailDate:   "durasi : 6 bulan",
		descProjects: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin quis risus ut mi euismod sodales. Mauris id quam ut massa sodales faucibus consectetur sit amet dolor. ",
		nodeJS:       "nodejs",
		reactJS:      "reactjs",
		nextJS:       "nextjs",
		typeScript:   "typescript",
	},
}

func main() {
	e := echo.New()
	e.Static("/assets", "assets")

	e.GET("/hello", helloWorld)
	e.GET("/about", about)
	e.GET("/", Home)
	e.GET("/contactMe", contactMe)
	e.GET("/projectPage", projectPage)
	e.GET("/projectDetail/:id", projectDetail)
	e.GET("/add-project", addProject)
	e.GET("/delete-project/:id", deleteProject)

	e.Logger.Fatal(e.Start("localhost:5500"))
}
func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func about(c echo.Context) error {
	return c.String(http.StatusOK, "this is about")
}

func Home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("tabs/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	post := map[string]interface{}{
		"Post": dataProject,
	}
	return tmpl.Execute(c.Response(), post)
}

func contactMe(c echo.Context) error {
	var tmpl, err = template.ParseFiles("tabs/contact-me.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func projectPage(c echo.Context) error {
	var tmpl, err = template.ParseFiles("tabs/project.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var tmpl, err = template.ParseFiles("tabs/project-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data := map[string]interface{}{
		"Id":           id,
		"Title":        "Ameno ameno latire Latiremo Dori me",
		"Author":       "Nafiisan N. Achmad",
		"detailDate":   "durasi : 4 bulan",
		"descProjects": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin quis risus ut mi euismod sodales. Mauris id quam ut massa sodales faucibus consectetur sit amet dolor. ",
		"nodeJS":       "nodejs",
		"reactJS":      "reactjs",
		"nextJS":       "nextjs",
		"typeScript":   "typescript",
	}

	return tmpl.Execute(c.Response(), data)
}

func addProject(c echo.Context) error {
	title := c.FormValue("title")
	DescProjects := c.FormValue("DescProjects")
	NodeJS := c.FormValue("nodeJs")
	ReactJS := c.FormValue("nodeReact")
	NextJS := c.FormValue("TypeScript")
	TypeScript := c.FormValue("reactJs")

	var addProject = Projects{
		Title:        title,
		Author:       "Nafiisan N. Achmad",
		detailDate:   time.Now().String(),
		descProjects: DescProjects,
		nodeJS:       NodeJS,
		reactJS:      ReactJS,
		nextJS:       NextJS,
		typeScript:   TypeScript,
	}

	fmt.Println(addProject)
	dataProject = append(dataProject, addProject)

	// fmt.Println(title)
	// fmt.Println(detailDate)
	// fmt.Println(DescProjects)
	// fmt.Println(nodeJS)
	// fmt.Println(ReactJS)
	// fmt.Println(NextJS)
	// fmt.Println(TypeScript)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	dataProject = append(dataProject,[:id], dataProject[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/index")
}
