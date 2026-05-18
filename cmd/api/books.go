package main

import "net/http"

func (app *application) getBooksHandler(w http.ResponseWriter, r *http.Request) {

	books, err := app.models.Book.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"books": books}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}
