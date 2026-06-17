package tasks

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	task := NewTask()
	if task == nil {
		t.Fatal("NewTask returned nil")
	}
}

func TestTask_Run(t *testing.T) {
	task := NewTask()

	executed := false
	fn := func(ctx context.Context) error {
		executed = true
		return nil
	}

	err := task.Run(context.Background(), fn)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !executed {
		t.Error("task function was not executed")
	}
}

func TestTask_RunWithError(t *testing.T) {
	task := NewTask()

	expectedErr := errors.New("task failed")
	fn := func(ctx context.Context) error {
		return expectedErr
	}

	err := task.Run(context.Background(), fn)
	if err != expectedErr {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
}

func TestTask_RunCancellation(t *testing.T) {
	task := NewTask()

	started := make(chan struct{})
	finished := make(chan struct{})

	fn := func(ctx context.Context) error {
		close(started)
		<-ctx.Done()
		close(finished)
		return ctx.Err()
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- task.Run(context.Background(), fn)
	}()

	// Wait for task to start
	<-started

	// Cancel the task
	task.Cancel()

	// Wait for task to finish
	select {
	case <-finished:
		// Success
	case <-time.After(1 * time.Second):
		t.Error("task did not finish after cancellation")
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

func TestTask_RunReplacePrevious(t *testing.T) {
	task := NewTask()

	first := make(chan struct{})
	second := make(chan struct{})

	fn1 := func(ctx context.Context) error {
		<-ctx.Done()
		close(first)
		return ctx.Err()
	}

	fn2 := func(ctx context.Context) error {
		close(second)
		return nil
	}

	// Start first task
	firstErr := make(chan error, 1)
	go func() {
		firstErr <- task.Run(context.Background(), fn1)
	}()
	time.Sleep(50 * time.Millisecond)

	// Start second task (should cancel first)
	err := task.Run(context.Background(), fn2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// First should be cancelled
	select {
	case <-first:
		// Success - first was cancelled
	case <-time.After(1 * time.Second):
		t.Error("first task was not cancelled")
	}

	select {
	case err := <-firstErr:
		if !errors.Is(err, context.Canceled) {
			t.Errorf("expected context canceled from first task, got %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Error("first task run did not return")
	}

	// Second should complete
	select {
	case <-second:
		// Success
	case <-time.After(100 * time.Millisecond):
		t.Error("second task did not complete")
	}
}

func TestTask_Cancel(t *testing.T) {
	task := NewTask()

	// Cancel without running should not panic
	task.Cancel()

	// Cancel multiple times should not panic
	task.Cancel()
	task.Cancel()
}

func TestTask_CancelWatch(t *testing.T) {
	task := NewTask()

	// CancelWatch without watching should not panic
	task.CancelWatch()

	// Multiple calls should not panic
	task.CancelWatch()
	task.CancelWatch()
}

func TestTask_Watch(t *testing.T) {
	task := NewTask()

	// Create a temporary file to watch
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")

	err := os.WriteFile(testFile, []byte("initial"), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	execCount := 0
	fn := func(ctx context.Context) error {
		execCount++
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- task.Watch(ctx, fn, testFile)
	}()

	// Wait for initial execution
	time.Sleep(100 * time.Millisecond)

	if execCount < 1 {
		t.Error("task should have executed at least once initially")
	}

	initialCount := execCount

	// Modify the file
	err = os.WriteFile(testFile, []byte("modified"), 0644)
	if err != nil {
		t.Fatalf("failed to modify test file: %v", err)
	}

	// Wait for watch to trigger
	time.Sleep(200 * time.Millisecond)

	if execCount <= initialCount {
		t.Error("task should have executed again after file modification")
	}

	// Cancel and wait for completion
	cancel()

	select {
	case err := <-done:
		if err != context.Canceled {
			t.Errorf("expected context.Canceled, got %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Error("watch did not stop after cancellation")
	}
}

func TestTask_WatchDirectory(t *testing.T) {
	task := NewTask()

	// Create a temporary directory to watch
	tmpDir := t.TempDir()

	execCount := 0
	fn := func(ctx context.Context) error {
		execCount++
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- task.Watch(ctx, fn, tmpDir)
	}()

	// Wait for initial execution
	time.Sleep(100 * time.Millisecond)

	if execCount < 1 {
		t.Error("task should have executed at least once initially")
	}

	// Wait for timeout
	<-done
}

func TestTask_WatchNonexistentFile(t *testing.T) {
	task := NewTask()

	execCount := 0
	fn := func(ctx context.Context) error {
		execCount++
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	// Watch a file that doesn't exist
	err := task.Watch(ctx, fn, "/nonexistent/file.txt")

	// Should still execute once (initial run)
	if execCount < 1 {
		t.Error("task should have executed at least once initially")
	}

	// Should complete without error (just no watching)
	if err != context.DeadlineExceeded {
		t.Logf("watch completed with: %v", err)
	}
}

func TestTask_WatchCancellation(t *testing.T) {
	task := NewTask()

	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")

	err := os.WriteFile(testFile, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	fn := func(ctx context.Context) error {
		time.Sleep(10 * time.Millisecond)
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan error, 1)
	go func() {
		done <- task.Watch(ctx, fn, testFile)
	}()

	// Let it run briefly
	time.Sleep(100 * time.Millisecond)

	// Cancel via context
	cancel()

	select {
	case <-done:
		// Success
	case <-time.After(1 * time.Second):
		t.Error("watch did not stop after context cancellation")
	}
}

func TestTask_WatchMultipleFiles(t *testing.T) {
	task := NewTask()

	tmpDir := t.TempDir()
	file1 := filepath.Join(tmpDir, "file1.txt")
	file2 := filepath.Join(tmpDir, "file2.txt")

	err := os.WriteFile(file1, []byte("file1"), 0644)
	if err != nil {
		t.Fatalf("failed to create file1: %v", err)
	}

	err = os.WriteFile(file2, []byte("file2"), 0644)
	if err != nil {
		t.Fatalf("failed to create file2: %v", err)
	}

	execCount := 0
	fn := func(ctx context.Context) error {
		execCount++
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- task.Watch(ctx, fn, file1, file2)
	}()

	// Wait for initial execution
	time.Sleep(100 * time.Millisecond)

	if execCount < 1 {
		t.Error("task should have executed at least once initially")
	}

	initialCount := execCount

	// Modify one file
	err = os.WriteFile(file1, []byte("modified"), 0644)
	if err != nil {
		t.Fatalf("failed to modify file1: %v", err)
	}

	// Wait for watch to trigger
	time.Sleep(200 * time.Millisecond)

	if execCount <= initialCount {
		t.Error("task should have executed after file modification")
	}

	cancel()

	select {
	case err := <-done:
		if err != nil && !errors.Is(err, context.Canceled) {
			t.Errorf("unexpected watch error: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Error("watch did not return after cancellation")
	}
}
