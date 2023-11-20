package internal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockTaskController struct {
	mock.Mock
}

func (m *MockTaskController) AddTaskToDB(description, priority, dueDate string) error {
	args := m.Called(description, priority, dueDate)
	return args.Error(0)
}

func (m *MockTaskController) ViewTasksFromDB() ([]Task, error) {
	args := m.Called()
	return args.Get(0).([]Task), args.Error(1)
}

func (m *MockTaskController) UpdateTaskInDB(id int, description, priority, dueDate string, completed bool) error {
	args := m.Called(id, description, priority, dueDate, completed)
	return args.Error(0)
}

func (m *MockTaskController) DeleteTaskFromDB(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestAddTaskHandler(t *testing.T) {
	// Create a mock task controller
	mockController := new(MockTaskController)

	// Set up the expected behavior and return values for the mock
	testTask := Task{
		Description: "Test Task",
		Priority:    "High",
		DueDate:     "2023-01-01",
	}
	mockController.On("AddTaskToDB", testTask.Description, testTask.Priority, testTask.DueDate).Return(nil)

	// Initialize the handler with the mock controller
	handler := NewTaskHandler(mockController)

	// Create an HTTP request and recorder
	body, _ := json.Marshal(testTask)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Get the HTTP handler and serve the request
	httpHandler := handler.AddTaskHandler(nil) // nil is passed for the pool as it's not used in the handler
	httpHandler.ServeHTTP(rr, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, rr.Code)
	mockController.AssertExpectations(t) // Ensure all expectations are met
}

func TestGetTasksHandler(t *testing.T) {
	mockController := new(MockTaskController)
	handler := NewTaskHandler(mockController)

	// Mock response
	tasks := []Task{{ID: 1, Description: "Test Task 1", Priority: "High", DueDate: "2023-01-01", Completed: false}}
	mockController.On("ViewTasksFromDB").Return(tasks, nil)

	// Create request and recorder
	req, _ := http.NewRequest("GET", "/tasks", nil)
	rr := httptest.NewRecorder()

	// Get the HTTP handler and serve the request
	httpHandler := handler.GetTasksHandler(nil)
	httpHandler.ServeHTTP(rr, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rr.Code)
	mockController.AssertExpectations(t)
}

// func TestUpdateTaskHandler(t *testing.T) {
// 	mockController := new(MockTaskController)
// 	handler := NewTaskHandler(mockController)

// 	// Mock behavior
// 	updatedTask := Task{Description: "Updated Task", Priority: "Medium", DueDate: "2023-02-01", Completed: true}
// 	taskID := 1
// 	mockController.On("UpdateTaskInDB", taskID, updatedTask.Description, updatedTask.Priority, updatedTask.DueDate, updatedTask.Completed).Return(nil)

// 	// Create request and recorder
// 	body, _ := json.Marshal(updatedTask)
// 	req, _ := http.NewRequest("PUT", "/tasks/"+strconv.Itoa(taskID), bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	rr := httptest.NewRecorder()

// 	// Get the HTTP handler and serve the request
// 	httpHandler := handler.UpdateTaskHandler(nil)
// 	httpHandler.ServeHTTP(rr, req)

// 	// Assertions
// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	mockController.AssertExpectations(t)
// }

// func TestDeleteTaskHandler(t *testing.T) {
// 	mockController := new(MockTaskController)
// 	handler := NewTaskHandler(mockController)

// 	// Mock behavior
// 	taskID := 1
// 	mockController.On("DeleteTaskFromDB", taskID).Return(nil)

// 	// Create request and recorder
// 	req, _ := http.NewRequest("DELETE", "/tasks/"+strconv.Itoa(taskID), nil)
// 	rr := httptest.NewRecorder()

// 	// Get the HTTP handler and serve the request
// 	httpHandler := handler.DeleteTaskHandler(nil)
// 	httpHandler.ServeHTTP(rr, req)

// 	// Assertions
// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	mockController.AssertExpectations(t)
// }

// func TestTaskHandlerIntegration(t *testing.T) {
// 	// Connect to the test database
// 	pool, err := ConnectDB()
// 	require.NoError(t, err)
// 	defer pool.Close()

// 	// Initialize handler with real controller
// 	controller := NewTaskController(pool)
// 	handler := NewTaskHandler(controller)

// 	// Create a test task
// 	testTask := Task{
// 		Description: "Integration Test Task",
// 		Priority:    "High",
// 		DueDate:     "2023-01-01",
// 	}
// 	body, _ := json.Marshal(testTask)
// 	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
// 	rr := httptest.NewRecorder()

// 	// Test the handler
// 	httpHandler := handler.AddTaskHandler(pool)
// 	httpHandler.ServeHTTP(rr, req)

// 	// Assert response
// 	assert.Equal(t, http.StatusCreated, rr.Code)
// }

func TestTaskServiceAcceptance(t *testing.T) {

	testTask := Task{
		Description: "Acceptance Test Task",
		Priority:    "Medium",
		DueDate:     "2023-02-01",
	}
	body, _ := json.Marshal(testTask)
	resp, err := http.Post("http://localhost:8081/tasks", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assert response
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	
	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("Response body: %s", string(bodyBytes))
	}
}
