package service

import (
	"context"

	"github.com/caioeverest/fed-its/adapter/database"
	"github.com/caioeverest/fed-its/internal/config"
	"github.com/caioeverest/fed-its/internal/logger"
	"github.com/caioeverest/fed-its/internal/validate"
	"github.com/caioeverest/fed-its/model"
)

type MethodI interface {
	Create(ctx context.Context, method model.Method) (model.Method, error)
	Get(ctx context.Context, method string) (model.Method, error)
	List(ctx context.Context) ([]model.Method, error)
}

type Method struct {
	cfg      *config.Config
	log      *logger.Logger
	db       *database.Database
	validate *validate.Validate
}

func NewMethod(cfg *config.Config, log *logger.Logger, db *database.Database, validate *validate.Validate) MethodI {
	return &Method{cfg, log, db, validate}
}

// Create a new method in the database and return it
func (m *Method) Create(ctx context.Context, method model.Method) (result model.Method, err error) {
	m.log.Infof("New method %s requested to be created", method.Name)

	//Validate input
	m.log.Info("Validating method")
	if err = m.validate.Struct(method); err != nil {
		m.log.Errorf("Validation error: %+v", err)
		return
	}

	//Create method
	if err = m.db.WithContext(ctx).Create(&method).Error; err != nil {
		m.log.Errorf("Error creating method - %+v", err)
		return
	}

	return
}

// Get a method from the database
func (m *Method) Get(ctx context.Context, methodName string) (method model.Method, err error) {
	m.log.Infof("Get method %s requested", methodName)
	if err = m.db.WithContext(ctx).Where("name = ?", methodName).First(&method).Error; err != nil {
		m.log.Errorf("Error getting method - %+v", err)
		return
	}
	m.log.Infof("Method %s found", methodName)
	return
}

// List all methods in the database
func (m *Method) List(ctx context.Context) (methods []model.Method, err error) {
	m.log.Info("List methods requested")
	if err = m.db.WithContext(ctx).Find(&methods).Error; err != nil {
		m.log.Errorf("Error listing methods - %+v", err)
		return
	}
	m.log.Infof("Found %d methods", len(methods))
	return
}
