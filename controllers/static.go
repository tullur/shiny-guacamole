package controllers

import "aproj/views"

type Static struct {
	Home    *views.View
	Contact *views.View
	About   *views.View
}

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "home/index"),
		Contact: views.NewView("bootstrap", "home/contact"),
		About:   views.NewView("bootstrap", "home/about"),
	}
}
