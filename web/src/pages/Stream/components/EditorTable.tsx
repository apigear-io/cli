import { useState } from 'react';
import { Table, Checkbox, ActionIcon, Text, Group, Button, LoadingOverlay, Box } from '@mantine/core';
import { IconStar, IconStarFilled } from '@tabler/icons-react';
import { useEditorContext } from './EditorContext';
import { useEditorMessages } from '@/api/queries';

export function EditorTable() {
  const {
    sessionStats,
    currentFilters,
    selectedIndices,
    toggleSelection,
    setSelectedIndices,
    markedIndices,
    toggleMarked,
    deletedIndices,
    hideDeleted,
    showMarkedOnly,
  } = useEditorContext();

  const [offset, setOffset] = useState(0);
  const limit = 100;

  const { data: messagesData, isLoading } = useEditorMessages(
    sessionStats?.sessionId || null,
    offset,
    limit,
    currentFilters
  );

  if (!sessionStats) return null;

  if (!messagesData) {
    return (
      <Box pos="relative" mih={400}>
        <LoadingOverlay visible={isLoading} />
      </Box>
    );
  }

  // Filter client-side for deleted/marked
  let visibleMessages = messagesData.messages;
  if (hideDeleted) {
    visibleMessages = visibleMessages.filter((m) => !deletedIndices.has(m.index));
  }
  if (showMarkedOnly) {
    visibleMessages = visibleMessages.filter((m) => markedIndices.has(m.index));
  }

  const formatTimestamp = (ts: number) => {
    const date = new Date(ts);
    return date.toLocaleTimeString();
  };

  const handleSelectAll = (checked: boolean) => {
    const newSet = new Set(selectedIndices);
    visibleMessages.forEach((m) => {
      if (checked) {
        newSet.add(m.index);
      } else {
        newSet.delete(m.index);
      }
    });
    setSelectedIndices(newSet);
  };

  const allSelected = visibleMessages.length > 0 && visibleMessages.every((m) => selectedIndices.has(m.index));

  return (
    <div>
      <div style={{ overflowX: 'auto' }}>
        <Table highlightOnHover stickyHeader>
          <Table.Thead>
            <Table.Tr>
              <Table.Th style={{ width: 40 }}>
                <Checkbox
                  checked={allSelected}
                  onChange={(e) => handleSelectAll(e.currentTarget.checked)}
                  indeterminate={!allSelected && visibleMessages.some((m) => selectedIndices.has(m.index))}
                />
              </Table.Th>
              <Table.Th style={{ width: 40 }}>★</Table.Th>
              <Table.Th style={{ width: 60 }}>#</Table.Th>
              <Table.Th>Time</Table.Th>
              <Table.Th>Proxy</Table.Th>
              <Table.Th>Dir</Table.Th>
              <Table.Th>Type</Table.Th>
              <Table.Th>Symbol</Table.Th>
              <Table.Th>ReqID</Table.Th>
              <Table.Th>Args</Table.Th>
            </Table.Tr>
          </Table.Thead>
          <Table.Tbody>
            {visibleMessages.length === 0 ? (
              <Table.Tr>
                <Table.Td colSpan={10}>
                  <Text size="sm" c="dimmed" ta="center" py="xl">
                    No messages found
                  </Text>
                </Table.Td>
              </Table.Tr>
            ) : (
              visibleMessages.map((message) => (
                <Table.Tr
                  key={message.index}
                  style={{
                    backgroundColor: deletedIndices.has(message.index)
                      ? '#ffe0e0'
                      : selectedIndices.has(message.index)
                      ? '#e7f5ff'
                      : undefined,
                  }}
                >
                  <Table.Td>
                    <Checkbox
                      checked={selectedIndices.has(message.index)}
                      onChange={() => toggleSelection(message.index)}
                    />
                  </Table.Td>
                  <Table.Td>
                    <ActionIcon
                      variant="subtle"
                      color={markedIndices.has(message.index) ? 'yellow' : 'gray'}
                      onClick={() => toggleMarked(message.index)}
                    >
                      {markedIndices.has(message.index) ? (
                        <IconStarFilled size={16} />
                      ) : (
                        <IconStar size={16} />
                      )}
                    </ActionIcon>
                  </Table.Td>
                  <Table.Td>
                    <Text size="sm" c="dimmed">
                      {message.index}
                    </Text>
                  </Table.Td>
                  <Table.Td>
                    <Text size="sm" ff="monospace">
                      {formatTimestamp(message.timestamp)}
                    </Text>
                  </Table.Td>
                  <Table.Td>
                    <Text size="sm">{message.proxy}</Text>
                  </Table.Td>
                  <Table.Td>
                    <Text size="sm" c={message.direction === 'SEND' ? 'blue' : 'green'}>
                      {message.direction}
                    </Text>
                  </Table.Td>
                  <Table.Td>
                    <Text size="sm">{message.parsed.msgTypeName}</Text>
                  </Table.Td>
                  <Table.Td>
                    <Text size="sm">{message.parsed.symbol || '-'}</Text>
                  </Table.Td>
                  <Table.Td>
                    <Text size="sm">{message.parsed.requestId || '-'}</Text>
                  </Table.Td>
                  <Table.Td>
                    <Text size="sm" lineClamp={1}>
                      {message.parsed.args ? JSON.stringify(message.parsed.args) : '-'}
                    </Text>
                  </Table.Td>
                </Table.Tr>
              ))
            )}
          </Table.Tbody>
        </Table>
      </div>

      {/* Pagination controls */}
      <Group p="md" justify="center">
        <Button size="xs" onClick={() => setOffset(Math.max(0, offset - limit))} disabled={offset === 0}>
          Previous
        </Button>
        <Text size="sm">
          {offset + 1} - {Math.min(offset + limit, messagesData.total)} of {messagesData.total}
        </Text>
        <Button
          size="xs"
          onClick={() => setOffset(offset + limit)}
          disabled={offset + limit >= messagesData.total}
        >
          Next
        </Button>
      </Group>
    </div>
  );
}
