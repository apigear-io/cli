package tasks

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

func TestNewTaskManager(t *testing.T) {
	tm := NewTaskManager()
	if tm == nil {
		t.Fatal("NewTaskManager returned nil")
	}
	if tm.tasks == nil {
		t.Error("tasks map not initialized")
	}
	if tm.hooks == nil {
		t.Error("hooks slice not initialized")
	}
}

func TestTaskManager_Register(t *testing.T) {
	tm := NewTaskManager()

	executed := false
	fn := func(ctx context.Context) error {
		executed = true
		return nil
	}

	task := tm.Register("test-task", nil, fn)
	if task == nil {
		t.Fatal("Register returned nil task")
	}

	if !tm.Has("test-task") {
		t.Error("task not registered")
	}

	// Run the task to verify it was stored correctly
	err := tm.Run(context.Background(), "test-task")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !executed {
		t.Error("task function was not executed")
	}
}

func TestTaskManager_RegisterReplaceExisting(t *testing.T) {
	tm := NewTaskManager()

	firstExecuted := false
	first := func(ctx context.Context) error {
		firstExecuted = true
		return nil
	}

	secondExecuted := false
	second := func(ctx context.Context) error {
		secondExecuted = true
		return nil
	}

	// Register first task
	tm.Register("task", nil, first)

	// Register second task with same name
	tm.Register("task", nil, second)

	// Run should execute the second task
	err := tm.Run(context.Background(), "task")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if firstExecuted {
		t.Error("first task should not have been executed")
	}
	if !secondExecuted {
		t.Error("second task should have been executed")
	}
}

func TestTaskManager_Run(t *testing.T) {
	tm := NewTaskManager()

	executed := false
	fn := func(ctx context.Context) error {
		executed = true
		return nil
	}

	tm.Register("test", nil, fn)

	err := tm.Run(context.Background(), "test")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !executed {
		t.Error("task was not executed")
	}
}

func TestTaskManager_RunNotFound(t *testing.T) {
	tm := NewTaskManager()

	err := tm.Run(context.Background(), "nonexistent")
	if err != ErrTaskNotFound {
		t.Errorf("expected ErrTaskNotFound, got %v", err)
	}
}

func TestTaskManager_RunWithError(t *testing.T) {
	tm := NewTaskManager()

	expectedErr := errors.New("task error")
	fn := func(ctx context.Context) error {
		return expectedErr
	}

	tm.Register("error-task", nil, fn)

	err := tm.Run(context.Background(), "error-task")
	if err != expectedErr {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
}

func TestTaskManager_Watch(t *testing.T) {
	tm := NewTaskManager()

	execCount := 0
	var mu sync.Mutex

	fn := func(ctx context.Context) error {
		mu.Lock()
		execCount++
		mu.Unlock()
		return nil
	}

	tm.Register("watch-task", nil, fn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start watching (this runs in background)
	err := tm.Watch(ctx, "watch-task")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Give it a moment to execute
	time.Sleep(100 * time.Millisecond)

	mu.Lock()
	count := execCount
	mu.Unlock()

	if count < 1 {
		t.Error("watch task should have executed at least once")
	}

	cancel()
}

func TestTaskManager_WatchNotFound(t *testing.T) {
	tm := NewTaskManager()

	err := tm.Watch(context.Background(), "nonexistent")
	if err != ErrTaskNotFound {
		t.Errorf("expected ErrTaskNotFound, got %v", err)
	}
}

func TestTaskManager_Cancel(t *testing.T) {
	tm := NewTaskManager()

	blocked := make(chan struct{})
	fn := func(ctx context.Context) error {
		<-ctx.Done()
		close(blocked)
		return ctx.Err()
	}

	tm.Register("blocking", nil, fn)

	// Start task in background
	errCh := make(chan error, 1)
	go func() {
		errCh <- tm.Run(context.Background(), "blocking")
	}()

	// Give it time to start
	time.Sleep(50 * time.Millisecond)

	// Cancel the task
	err := tm.Cancel("blocking")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Wait for task to finish
	select {
	case <-blocked:
		// Success - task was cancelled
	case <-time.After(1 * time.Second):
		t.Error("task was not cancelled in time")
	}

	select {
	case runErr := <-errCh:
		if runErr != nil && !errors.Is(runErr, context.Canceled) {
			t.Errorf("unexpected run error: %v", runErr)
		}
	case <-time.After(1 * time.Second):
		t.Error("run did not return after cancellation")
	}
}

func TestTaskManager_CancelNotFound(t *testing.T) {
	tm := NewTaskManager()

	err := tm.Cancel("nonexistent")
	if err != ErrTaskNotFound {
		t.Errorf("expected ErrTaskNotFound, got %v", err)
	}
}

func TestTaskManager_CancelAll(t *testing.T) {
	tm := NewTaskManager()

	count := 3
	blocked := make([]chan struct{}, count)

	for i := 0; i < count; i++ {
		blocked[i] = make(chan struct{})
		ch := blocked[i] // capture for closure
		fn := func(ctx context.Context) error {
			<-ctx.Done()
			close(ch)
			return ctx.Err()
		}
		tm.Register("task-"+string(rune('A'+i)), nil, fn)
	}

	// Start all tasks
	errs := make(chan error, count)
	for i := 0; i < count; i++ {
		name := "task-" + string(rune('A'+i))
		go func(taskName string) {
			errs <- tm.Run(context.Background(), taskName)
		}(name)
	}

	time.Sleep(50 * time.Millisecond)

	// Cancel all
	tm.CancelAll()

	// Wait for all to finish
	for i := 0; i < count; i++ {
		select {
		case <-blocked[i]:
			// Success
		case <-time.After(1 * time.Second):
			t.Errorf("task %d was not cancelled", i)
		}
	}

	for i := 0; i < count; i++ {
		select {
		case runErr := <-errs:
			if runErr != nil && !errors.Is(runErr, context.Canceled) {
				t.Errorf("unexpected run error: %v", runErr)
			}
		case <-time.After(1 * time.Second):
			t.Errorf("run %d did not return after CancelAll", i)
		}
	}
}

func TestTaskManager_Has(t *testing.T) {
	tm := NewTaskManager()

	if tm.Has("test") {
		t.Error("should not have task before registration")
	}

	tm.Register("test", nil, func(ctx context.Context) error { return nil })

	if !tm.Has("test") {
		t.Error("should have task after registration")
	}
}

func TestTaskManager_Names(t *testing.T) {
	tm := NewTaskManager()

	names := tm.Names()
	if len(names) != 0 {
		t.Error("expected empty names initially")
	}

	tm.Register("task1", nil, func(ctx context.Context) error { return nil })
	tm.Register("task2", nil, func(ctx context.Context) error { return nil })
	tm.Register("task3", nil, func(ctx context.Context) error { return nil })

	names = tm.Names()
	if len(names) != 3 {
		t.Errorf("expected 3 names, got %d", len(names))
	}

	// Check all names are present
	nameSet := make(map[string]bool)
	for _, name := range names {
		nameSet[name] = true
	}

	for _, expected := range []string{"task1", "task2", "task3"} {
		if !nameSet[expected] {
			t.Errorf("missing task name: %s", expected)
		}
	}
}

func TestTaskManager_AddHook(t *testing.T) {
	tm := NewTaskManager()

	events := make([]TaskEvent, 0)
	var mu sync.Mutex

	tm.AddHook(func(evt *TaskEvent) {
		mu.Lock()
		events = append(events, *evt)
		mu.Unlock()
	})

	fn := func(ctx context.Context) error { return nil }
	tm.Register("hooked", nil, fn)

	err := tm.Run(context.Background(), "hooked")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Give hooks time to fire
	time.Sleep(50 * time.Millisecond)

	mu.Lock()
	eventCount := len(events)
	mu.Unlock()

	// Should have: Added, Running, Finished
	if eventCount < 3 {
		t.Errorf("expected at least 3 events, got %d", eventCount)
	}

	// Check event states
	mu.Lock()
	hasAdded := false
	hasRunning := false
	hasFinished := false
	for _, evt := range events {
		switch evt.State {
		case TaskStateAdded:
			hasAdded = true
		case TaskStateRunning:
			hasRunning = true
		case TaskStateFinished:
			hasFinished = true
		}
	}
	mu.Unlock()

	if !hasAdded {
		t.Error("missing TaskStateAdded event")
	}
	if !hasRunning {
		t.Error("missing TaskStateRunning event")
	}
	if !hasFinished {
		t.Error("missing TaskStateFinished event")
	}
}

func TestTaskManager_HooksOnError(t *testing.T) {
	tm := NewTaskManager()

	var lastState TaskState
	var mu sync.Mutex

	tm.AddHook(func(evt *TaskEvent) {
		mu.Lock()
		lastState = evt.State
		mu.Unlock()
	})

	fn := func(ctx context.Context) error {
		return errors.New("fail")
	}

	tm.Register("failing", nil, fn)
	if err := tm.Run(context.Background(), "failing"); err == nil {
		t.Fatal("expected run to fail")
	}

	time.Sleep(50 * time.Millisecond)

	mu.Lock()
	state := lastState
	mu.Unlock()

	if state != TaskStateFailed {
		t.Errorf("expected TaskStateFailed, got %v", state)
	}
}

func TestTaskManager_Concurrent(t *testing.T) {
	tm := NewTaskManager()

	var wg sync.WaitGroup
	count := 10

	// Concurrent registrations
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			name := "concurrent-" + string(rune('0'+n))
			tm.Register(name, nil, func(ctx context.Context) error {
				time.Sleep(10 * time.Millisecond)
				return nil
			})
		}(i)
	}

	wg.Wait()

	// Concurrent runs
	errCh := make(chan error, count)
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			name := "concurrent-" + string(rune('0'+n))
			errCh <- tm.Run(context.Background(), name)
		}(i)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			t.Errorf("concurrent run failed: %v", err)
		}
	}

	// Verify all registered
	if len(tm.Names()) != count {
		t.Errorf("expected %d tasks, got %d", count, len(tm.Names()))
	}
}
