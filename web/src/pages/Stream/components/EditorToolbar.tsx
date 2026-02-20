import { Group, Button, Text, Badge, Menu } from '@mantine/core';
import {
  IconSelect,
  IconSelectAll,
  IconX,
  IconCut,
  IconRestore,
  IconStar,
  IconStarOff,
  IconDownload,
} from '@tabler/icons-react';
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

  const handleSelectAll = () => {
    if (!messagesData) return;
    const allIndices = new Set(messagesData.messages.map((m) => m.index));
    setSelectedIndices(allIndices);
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
    <Group gap="md" p="xs" style={{ borderTop: '1px solid #dee2e6' }} wrap="wrap">
      <Group gap="xs">
        <IconSelect size={16} />
        <Text size="sm" fw={500}>
          Selection:
        </Text>
        {selectedCount > 0 ? (
          <Badge size="md" color="blue">
            {selectedCount.toLocaleString()} selected
          </Badge>
        ) : (
          <Text size="sm" c="dimmed">
            None
          </Text>
        )}
      </Group>

      {selectedCount > 0 && (
        <Button
          size="xs"
          variant="subtle"
          color="gray"
          leftSection={<IconX size={14} />}
          onClick={clearSelection}
        >
          Clear
        </Button>
      )}

      <Button
        size="xs"
        variant="light"
        leftSection={<IconSelectAll size={14} />}
        onClick={handleSelectAll}
      >
        Select All
      </Button>

      {/* Cut operations */}
      {selectedCount > 0 && (
        <Menu shadow="md" width={200}>
          <Menu.Target>
            <Button size="xs" variant="light" color="orange" leftSection={<IconCut size={14} />}>
              Cut
            </Button>
          </Menu.Target>

          <Menu.Dropdown>
            <Menu.Item onClick={handleCutBefore}>Cut Before Selected</Menu.Item>
            <Menu.Item onClick={handleCutSelected}>Cut Selected</Menu.Item>
            <Menu.Item onClick={handleCutAfter}>Cut After Selected</Menu.Item>
          </Menu.Dropdown>
        </Menu>
      )}

      {deletedCount > 0 && (
        <>
          <Badge size="md" color="orange">
            {deletedCount.toLocaleString()} deleted
          </Badge>
          <Button
            size="xs"
            variant="subtle"
            color="orange"
            leftSection={<IconRestore size={14} />}
            onClick={handleUndoAllCuts}
          >
            Undo All Cuts
          </Button>
        </>
      )}

      {/* Mark operations */}
      {selectedCount > 0 && (
        <>
          <Button
            size="xs"
            variant="light"
            color="yellow"
            leftSection={<IconStar size={14} />}
            onClick={handleMarkSelected}
          >
            Mark Selected
          </Button>
          <Button
            size="xs"
            variant="subtle"
            color="yellow"
            leftSection={<IconStarOff size={14} />}
            onClick={handleUnmarkSelected}
          >
            Unmark
          </Button>
        </>
      )}

      {markedCount > 0 && (
        <Badge size="md" color="yellow">
          {markedCount.toLocaleString()} marked
        </Badge>
      )}

      {/* Export menu */}
      <Menu shadow="md" width={200}>
        <Menu.Target>
          <Button
            size="xs"
            variant="filled"
            color="blue"
            leftSection={<IconDownload size={14} />}
            loading={exportMutation.isPending}
          >
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
  );
}
