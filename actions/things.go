package actions

import (
	"fmt"
	"hundred_go_buffalo/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/x/responder"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Thing)
// DB Table: Plural (things)
// Resource: Plural (Things)
// Path: Plural (/things)
// View Template Folder: Plural (/templates/things/)

// ThingsResource is the resource for the Thing model
type ThingsResource struct {
	buffalo.Resource
}

// List gets all Things. This function is mapped to the path
// GET /things
func (v ThingsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	things := &models.Things{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Things from the DB
	if err := q.All(things); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("things", things)
		return c.Render(http.StatusOK, r.HTML("/things/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(things))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(things))
	}).Respond(c)
}

// Show gets the data for one Thing. This function is mapped to
// the path GET /things/{thing_id}
func (v ThingsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Thing
	thing := &models.Thing{}

	// To find the Thing the parameter thing_id is used.
	if err := tx.Find(thing, c.Param("thing_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("thing", thing)

		return c.Render(http.StatusOK, r.HTML("/things/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(thing))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(thing))
	}).Respond(c)
}

// New renders the form for creating a new Thing.
// This function is mapped to the path GET /things/new
func (v ThingsResource) New(c buffalo.Context) error {
	c.Set("thing", &models.Thing{})

	return c.Render(http.StatusOK, r.HTML("/things/new.plush.html"))
}

// Create adds a Thing to the DB. This function is mapped to the
// path POST /things
func (v ThingsResource) Create(c buffalo.Context) error {
	// Allocate an empty Thing
	thing := &models.Thing{}

	// Bind thing to the html form elements
	if err := c.Bind(thing); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(thing)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("thing", thing)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("/things/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "thing.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/things/%v", thing.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(thing))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(thing))
	}).Respond(c)
}

// Edit renders a edit form for a Thing. This function is
// mapped to the path GET /things/{thing_id}/edit
func (v ThingsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Thing
	thing := &models.Thing{}

	if err := tx.Find(thing, c.Param("thing_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("thing", thing)
	return c.Render(http.StatusOK, r.HTML("/things/edit.plush.html"))
}

// Update changes a Thing in the DB. This function is mapped to
// the path PUT /things/{thing_id}
func (v ThingsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Thing
	thing := &models.Thing{}

	if err := tx.Find(thing, c.Param("thing_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Thing to the html form elements
	if err := c.Bind(thing); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(thing)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("thing", thing)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("/things/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "thing.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/things/%v", thing.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(thing))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(thing))
	}).Respond(c)
}

// Destroy deletes a Thing from the DB. This function is mapped
// to the path DELETE /things/{thing_id}
func (v ThingsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Thing
	thing := &models.Thing{}

	// To find the Thing the parameter thing_id is used.
	if err := tx.Find(thing, c.Param("thing_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(thing); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "thing.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/things")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(thing))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(thing))
	}).Respond(c)
}
