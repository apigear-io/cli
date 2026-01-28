# tasks

Task management and execution framework with file watching.

## Purpose

The `tasks` package provides a framework for registering, running, and monitoring tasks with support for:

- One-time task execution
- File/directory watching with automatic re-execution on changes
- Task lifecycle management (creation, execution, cancellation)
- Event-driven notifications for task state changes

## Key Exports

### Types
- `TaskFunc` - Function type: `func(ctx context.Context) error`
- `TaskItem` - Individual task with execution control
- `TaskManager` - Central manager for task lifecycle
- `TaskEvent` - Event emitted on state changes
- `TaskState` - States: Idle, Added, Removed, Watching, Running, Finished, Stopped, Failed

### TaskItem Methods
- `NewTaskItem()` - Create new task item
- `Run()` - Execute task once
- `Watch()` - Monitor dependencies for changes
- `Cancel()`, `CancelWatch()` - Cancel operations
- `UpdateMeta()` - Update task metadata

### TaskManager Methods
- `NewTaskManager()` - Create new manager
- `Register()` - Create and register task
- `AddTask()`, `RmTask()` - Collection management
- `Get()`, `Has()` - Task lookup
- `Run()`, `Watch()` - Execute or watch task
- `Cancel()`, `CancelAll()` - Cancel tasks
- `Names()` - List registered task names

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Configuration access |
| `helper` | IsDir utility, Hook pattern |
| `log` | Logging |
