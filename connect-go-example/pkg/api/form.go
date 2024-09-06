package api

import (
	"connectrpc.com/connect"
	"context"
	greetv1 "example/gen/example/v1"
	"example/pkg/model"
	"github.com/uptrace/bun"
	"log"
	"time"
)

// FormServer is the struct implementing the form service
type FormServer struct {
	DB *bun.DB
}

// SaveFormModel saves or updates a form model and its todos in the database
func (f FormServer) SaveFormModel(ctx context.Context, req *connect.Request[greetv1.CreateFormModelRequest]) (*connect.Response[greetv1.GetFormModelResponse], error) {
	// Parse the DateTime
	t, err := time.Parse(time.RFC3339, req.Msg.FormModel.DateTime)
	if err != nil {
		log.Printf("invalid datetime format: %v", err)
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	formModel := &model.FormModel{
		Name:              req.Msg.FormModel.Name,
		Age:               req.Msg.FormModel.Age,
		DateTime:          t,
		Skills:            req.Msg.FormModel.Skills,
		City:              req.Msg.FormModel.City,
		Range:             req.Msg.FormModel.Range,
		TermsAccepted:     req.Msg.FormModel.TermsAccepted,
		SatisfactionLevel: req.Msg.FormModel.SatisfactionLevel,
		Gender:            req.Msg.FormModel.Gender,
	}

	// Start a transaction
	tx, err := f.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("failed to begin transaction: %v", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	defer tx.Rollback()

	// Check if we're updating or creating
	if req.Msg.FormModel.Id != 0 {
		formModel.ID = req.Msg.FormModel.Id // Set the existing ID for updates

		// Update the form model in the database
		_, err = tx.NewUpdate().Model(formModel).Where("id = ?", formModel.ID).Exec(ctx)
		if err != nil {
			log.Printf("failed to update form model: %v", err)
			return nil, connect.NewError(connect.CodeInternal, err)
		}

		// Delete existing todos for this form model (to replace with the updated ones)
		_, err = tx.NewDelete().Model((*model.Todo)(nil)).Where("form_model_id = ?", formModel.ID).Exec(ctx)
		if err != nil {
			log.Printf("failed to delete existing todos: %v", err)
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	} else {
		// Insert a new form model in the database
		_, err = tx.NewInsert().Model(formModel).Exec(ctx)
		if err != nil {
			log.Printf("failed to insert new form model: %v", err)
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}

	// Insert new todos for this form model
	for _, todoReq := range req.Msg.FormModel.Todos {
		todo := &model.Todo{
			Task:        todoReq.Task,
			FormModelID: formModel.ID, // Associate todo with form model
		}
		_, err = tx.NewInsert().Model(todo).Exec(ctx)
		if err != nil {
			log.Printf("failed to insert todos: %v", err)
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		log.Printf("failed to commit transaction: %v", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Prepare the response
	response := &greetv1.GetFormModelResponse{
		FormModel: &greetv1.FormModel{
			Id:                formModel.ID,
			Name:              formModel.Name,
			Age:               formModel.Age,
			DateTime:          formModel.DateTime.Format(time.RFC3339),
			Skills:            formModel.Skills,
			City:              formModel.City,
			Range:             formModel.Range,
			TermsAccepted:     formModel.TermsAccepted,
			SatisfactionLevel: formModel.SatisfactionLevel,
			Gender:            formModel.Gender,
			Todos:             []*greetv1.Todo{},
		},
	}

	// Return response with the saved todos
	for _, todoReq := range req.Msg.FormModel.Todos {
		response.FormModel.Todos = append(response.FormModel.Todos, &greetv1.Todo{
			Task: todoReq.Task,
		})
	}

	return connect.NewResponse(response), nil
}

// GetFormModel retrieves a form model by its ID from the database along with todos
func (f FormServer) GetFormModel(ctx context.Context, req *connect.Request[greetv1.GetFormModelRequest]) (*connect.Response[greetv1.GetFormModelResponse], error) {
	formModel := new(model.FormModel)

	// Find the form model by ID
	err := f.DB.NewSelect().Model(formModel).Where("id = ?", req.Msg.Id).Scan(ctx)
	if err != nil {
		log.Printf("failed to get form model: %v", err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	// Fetch associated todos
	var todos []model.Todo
	err = f.DB.NewSelect().Model(&todos).Where("form_model_id = ?", formModel.ID).Scan(ctx)
	if err != nil {
		log.Printf("failed to get todos for form model: %v", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Map the found data to the response
	response := &greetv1.GetFormModelResponse{
		FormModel: &greetv1.FormModel{
			Id:                formModel.ID,
			Name:              formModel.Name,
			Age:               formModel.Age,
			DateTime:          formModel.DateTime.Format(time.RFC3339),
			Skills:            formModel.Skills,
			City:              formModel.City,
			Range:             formModel.Range,
			TermsAccepted:     formModel.TermsAccepted,
			SatisfactionLevel: formModel.SatisfactionLevel,
			Gender:            formModel.Gender,
			Todos:             []*greetv1.Todo{},
		},
	}

	// Add todos to the response
	for _, todo := range todos {
		response.FormModel.Todos = append(response.FormModel.Todos, &greetv1.Todo{
			Task: todo.Task,
		})
	}

	return connect.NewResponse(response), nil
}

// ListFormModels lists all form models with optional pagination and their todos
func (f FormServer) ListFormModels(ctx context.Context, req *connect.Request[greetv1.ListFormModelsRequest]) (*connect.Response[greetv1.ListFormModelsResponse], error) {
	formModels := make([]model.FormModel, 0)

	// Query the database for form models with pagination
	query := f.DB.NewSelect().Model(&formModels)
	if req.Msg.PageSize > 0 {
		query = query.Limit(int(req.Msg.PageSize)).Offset(int(req.Msg.Page) * int(req.Msg.PageSize))
	}

	// Execute the query
	if err := query.Scan(ctx); err != nil {
		log.Printf("failed to list form models: %v", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Map the form models to the response
	responseModels := make([]*greetv1.FormModel, len(formModels))
	for i, formModel := range formModels {
		responseModels[i] = &greetv1.FormModel{
			Id:                formModel.ID,
			Name:              formModel.Name,
			Age:               formModel.Age,
			DateTime:          formModel.DateTime.Format(time.RFC3339),
			Skills:            formModel.Skills,
			City:              formModel.City,
			Range:             formModel.Range,
			TermsAccepted:     formModel.TermsAccepted,
			SatisfactionLevel: formModel.SatisfactionLevel,
			Gender:            formModel.Gender,
			Todos:             []*greetv1.Todo{},
		}

		// Fetch todos for each form model
		var todos []model.Todo
		err := f.DB.NewSelect().Model(&todos).Where("form_model_id = ?", formModel.ID).Scan(ctx)
		if err != nil {
			log.Printf("failed to get todos for form model %v: %v", formModel.ID, err)
			return nil, connect.NewError(connect.CodeInternal, err)
		}

		// Add todos to the response
		for _, todo := range todos {
			responseModels[i].Todos = append(responseModels[i].Todos, &greetv1.Todo{
				Id:     todo.ID,
				Task:   todo.Task,
				FormId: formModel.ID,
			})
		}
	}

	// Prepare and send the response
	response := &greetv1.ListFormModelsResponse{
		FormModels: responseModels,
	}
	return connect.NewResponse(response), nil
}

// DeleteFormModel deletes a form model by its ID and associated todos
func (f FormServer) DeleteFormModel(ctx context.Context, req *connect.Request[greetv1.DeleteFormModelRequest]) (*connect.Response[greetv1.DeleteFormModelResponse], error) {
	// Start a transaction to ensure all deletions succeed or none
	tx, err := f.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("failed to begin transaction: %v", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	defer tx.Rollback()

	// Delete associated todos first
	_, err = tx.NewDelete().Model((*model.Todo)(nil)).Where("form_model_id = ?", req.Msg.Id).Exec(ctx)
	if err != nil {
		log.Printf("failed to delete todos for form model %v: %v", req.Msg.Id, err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Delete the form model
	_, err = tx.NewDelete().Model((*model.FormModel)(nil)).Where("id = ?", req.Msg.Id).Exec(ctx)
	if err != nil {
		log.Printf("failed to delete form model %v: %v", req.Msg.Id, err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		log.Printf("failed to commit transaction: %v", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Return success response
	return connect.NewResponse(&greetv1.DeleteFormModelResponse{Success: true}), nil
}
