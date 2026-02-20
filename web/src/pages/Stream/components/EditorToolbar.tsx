import { Group, Button, Text, Menu } from '@mantine/core';
import { useEditorContext } from './EditorContext';
import { useEditorMessages, useEditorExport } from '@/api/queries';
import { notifications } from '@mantine/notifications';

export function EditorToolbar() {
  const {
    sessionStats,
    selectedIndices,
    clearSelection,
    setSelectedIndices,
    deletedIndices,
    setDeletedIndices,
    markedIndices,
    setMarkedIndices,
    currentFilters,
  } = useEditorContext();

  const { data: messagesData } = useEditorMessages(
    sessionStats?.sessionId || null,
    0,
    10000, // Load all for operations
    currentFilters
  );

  const exportMutation = useEditorExport();

  const selectedCount = selectedIndices.size;
  const deletedCount = deletedIndices.size;
  const markedCount = markedIndices.size;

  const handleSelectNone = () => {
    clearSelection();
  };

  const handleSelectFiltered = () => {
    if (!messagesData) return;
    const filteredIndices = new Set(messagesData.messages.map((m) => m.index));
    setSelectedIndices(filteredIndices);
  };

  const handleInvertSelection = () => {
    if (!messagesData) return;
    const allIndices = new Set(messagesData.messages.map((m) => m.index));
    const newSelected = new Set<number>();
    allIndices.forEach((idx) => {
      if (!selectedIndices.has(idx)) {
        newSelected.add(idx);
      }
    });
    setSelectedIndices(newSelected);
  };

  const handleSelectMarked = () => {
    setSelectedIndices(new Set(markedIndices));
  };

  const handleCutSelected = () => {
    const newDeleted = new Set(deletedIndices);
    selectedIndices.forEach((idx) => newDeleted.add(idx));
    setDeletedIndices(newDeleted);
    clearSelection();
    notifications.show({
      title: 'Messages Cut',
      message: `Marked ${selectedIndices.size} messages as deleted`,
      color: 'orange',
    });
  };

  const handleCutBefore = () => {
    if (selectedIndices.size === 0 || !messagesData) return;
    const minSelected = Math.min(...Array.from(selectedIndices));
    const newDeleted = new Set(deletedIndices);
    messagesData.messages.forEach((m) => {
      if (m.index < minSelected) newDeleted.add(m.index);
    });
    setDeletedIndices(newDeleted);
    notifications.show({
      title: 'Messages Cut',
      message: 'Marked all messages before selection as deleted',
      color: 'orange',
    });
  };

  const handleCutAfter = () => {
    if (selectedIndices.size === 0 || !messagesData) return;
    const maxSelected = Math.max(...Array.from(selectedIndices));
    const newDeleted = new Set(deletedIndices);
    messagesData.messages.forEach((m) => {
      if (m.index > maxSelected) newDeleted.add(m.index);
    });
    setDeletedIndices(newDeleted);
    notifications.show({
      title: 'Messages Cut',
      message: 'Marked all messages after selection as deleted',
      color: 'orange',
    });
  };

  const handleUndoAllCuts = () => {
    setDeletedIndices(new Set());
    notifications.show({
      title: 'Cuts Undone',
      message: 'Restored all deleted messages',
      color: 'green',
    });
  };

  const handleMarkSelected = () => {
    const newMarked = new Set(markedIndices);
    selectedIndices.forEach((idx) => newMarked.add(idx));
    setMarkedIndices(newMarked);
    notifications.show({
      title: 'Messages Marked',
      message: `Marked ${selectedIndices.size} messages`,
      color: 'yellow',
    });
  };

  const handleUnmarkSelected = () => {
    const newMarked = new Set(markedIndices);
    selectedIndices.forEach((idx) => newMarked.delete(idx));
    setMarkedIndices(newMarked);
    notifications.show({
      title: 'Messages Unmarked',
      message: `Unmarked ${selectedIndices.size} messages`,
      color: 'gray',
    });
  };

  const handleExport = async (indices?: number[]) => {
    if (!sessionStats) return;

    try {
      const blob = await exportMutation.mutateAsync({
        sessionId: sessionStats.sessionId,
        indices,
      });

      // Download file
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = sessionStats.filename;
      document.body.appendChild(a);
      a.click();
      window.URL.revokeObjectURL(url);
      document.body.removeChild(a);

      notifications.show({
        title: 'Export Complete',
        message: `Downloaded ${sessionStats.filename}`,
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Export Failed',
        message: error instanceof Error ? error.message : 'Export failed',
        color: 'red',
      });
    }
  };

  if (!sessionStats) return null;

  return (
    <Group gap="md" p="sm" style={{ borderTop: '1px solid #dee2e6' }} wrap="wrap" justify="space-between">
      {/* Left side - Selection controls */}
      <Group gap="md">
        <Text size="sm">{selectedCount} selected</Text>

        <Group gap={4}>
          <Button size="xs" variant="subtle" onClick={handleSelectNone}>
            Select None
          </Button>
          <Text size="sm" c="dimmed">
            |
          </Text>
          <Button size="xs" variant="subtle" onClick={handleSelectFiltered}>
            Select Filtered
          </Button>
          <Text size="sm" c="dimmed">
            |
          </Text>
          <Button size="xs" variant="subtle" onClick={handleInvertSelection}>
            Invert
          </Button>
        </Group>

        {/* Mark operations */}
        <Group gap="xs">
          <Button size="xs" variant="light" color="orange" onClick={handleMarkSelected} disabled={selectedCount === 0}>
            Mark
          </Button>
          <Button size="xs" variant="light" color="orange" onClick={handleUnmarkSelected} disabled={selectedCount === 0}>
            Unmark
          </Button>
          <Button size="xs" variant="light" color="orange" onClick={handleSelectMarked} disabled={markedCount === 0}>
            Select Marked
          </Button>
        </Group>
      </Group>

      {/* Right side - Cut operations and Export */}
      <Group gap="md">
        {/* Cut operations */}
        <Group gap="xs">
          <Button size="xs" variant="light" color="orange" onClick={handleCutBefore} disabled={selectedCount === 0}>
            Cut Before
          </Button>
          <Button size="xs" variant="light" color="orange" onClick={handleCutSelected} disabled={selectedCount === 0}>
            Cut Selected
          </Button>
          <Button size="xs" variant="light" color="orange" onClick={handleCutAfter} disabled={selectedCount === 0}>
            Cut After
          </Button>
        </Group>

        {deletedCount > 0 && (
          <Button size="xs" variant="outline" color="orange" onClick={handleUndoAllCuts}>
            Undo All Cuts
          </Button>
        )}

        {/* Export menu */}
        <Menu shadow="md" width={200}>
          <Menu.Target>
            <Button size="xs" variant="filled" color="blue" loading={exportMutation.isPending}>
              Export
            </Button>
          </Menu.Target>

          <Menu.Dropdown>
            <Menu.Item onClick={() => handleExport()}>Export All</Menu.Item>
            <Menu.Item onClick={() => handleExport(Array.from(selectedIndices))} disabled={selectedCount === 0}>
              Export Selected ({selectedCount})
            </Menu.Item>
            <Menu.Item onClick={() => handleExport(Array.from(markedIndices))} disabled={markedCount === 0}>
              Export Marked ({markedCount})
            </Menu.Item>
            <Menu.Item
              onClick={() => {
                if (!messagesData) return;
                const undeleted = messagesData.messages
                  .filter((m) => !deletedIndices.has(m.index))
                  .map((m) => m.index);
                handleExport(undeleted);
              }}
              disabled={deletedCount === 0}
            >
              Export Undeleted
            </Menu.Item>
          </Menu.Dropdown>
        </Menu>
      </Group>
    </Group>
  );
}
