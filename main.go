package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"container/list"
	"strconv"
)

type Alumno struct {
	ANombre       string
	AMateria      string
	ACalificacion float64
}

var listaAlumno = list.New()

func cargarHtml(a string) string {
	html, _ := ioutil.ReadFile(a)

	return string(html)
}

func form(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("form.html"),
	)
}

func menu(res http.ResponseWriter, req *http.Request){
	switch req.Method{
	case "POST":
		if error:= req.ParseForm();error != nil{
			fmt.Fprintf(res, "ParseForm() error %v", error)
			return
		}
		opc:=req.FormValue("opcion")
		if opc=="agregarAlumno"{
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("agregarAlumno.html"),
			)
		}else if opc=="promedioAlumno"{
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("promedioAlumno.html"),
			)
		}else if opc=="promedioGeneral"{
			promedioGeneral(res,req)
		}else if opc=="promedioMateria"{
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("promedioMateria.html"),
			)
		}
	}
}

func agrAlum(res http.ResponseWriter, req *http.Request){
	switch req.Method{
	case "POST":
		if error:= req.ParseForm();error != nil{
			fmt.Fprintf(res, "ParseForm() error %v", error)
			return
		}
		aux, _:=strconv.ParseFloat(req.FormValue("calificacion"), 64)
		alumno:=Alumno{ANombre:req.FormValue("nombreAlumno"), AMateria: req.FormValue("nombreMateria"), ACalificacion: aux}
		listaAlumno.PushBack(alumno)
	}
	
}

func promAlum(res http.ResponseWriter, req *http.Request){
	switch req.Method{
	case "POST":
		if error:= req.ParseForm();error != nil{
			fmt.Fprintf(res, "ParseForm() error %v", error)
			return
		}
		alumno:=req.FormValue("nombreAlumno")
		var prom float64
		var i float64
		prom=0
		i=0
		for e:=listaAlumno.Front();e!=nil;e=e.Next(){
			aux:=e.Value.(Alumno)
			if alumno == aux.ANombre{
				prom += aux.ACalificacion
				i++
			}
		}
		if i==0{
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("error.html"),
				prom,
			)
		}else{
			prom=prom/i
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("respuesta.html"),
				prom,
			)
		}
		
	}
	
}

func promMat(res http.ResponseWriter, req *http.Request){
	switch req.Method{
	case "POST":
		if error:= req.ParseForm();error != nil{
			fmt.Fprintf(res, "ParseForm() error %v", error)
			return
		}
		materia:=req.FormValue("nombreMateria")
		var prom float64
		var i float64
		prom=0
		i=0
		for e:=listaAlumno.Front();e!=nil;e=e.Next(){
			aux:=e.Value.(Alumno)
			if materia == aux.AMateria{
				prom += aux.ACalificacion
				i++
			}
		}
		if i==0{
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("error.html"),
				prom,
			)
		}else{
			prom=prom/i
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("respuesta.html"),
				prom,
			)
		}
		
	}
	
}

func promedioGeneral(res http.ResponseWriter, req *http.Request){
	var prom float64
	var i float64
	prom=0
	i=0
	for e:=listaAlumno.Front();e!=nil;e=e.Next(){
		aux:=e.Value.(Alumno)
		prom += aux.ACalificacion
		i++
	}
	if i==0{
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("error.html"),
		)
	}else{
		prom=prom/i
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("respuesta.html"),
			prom,
		)
	}

}

func main() {
	http.HandleFunc("/form", form)
	http.HandleFunc("/menu", menu)
	http.HandleFunc("/agrAlum", agrAlum)
	http.HandleFunc("/promAlum", promAlum)
	http.HandleFunc("/promMat", promMat)
	fmt.Println("Corriendo servirdor de tareas...")
	http.ListenAndServe(":9000", nil)
}