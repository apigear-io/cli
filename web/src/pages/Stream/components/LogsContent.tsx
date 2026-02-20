import { useState } from 'react';
import {
  Stack,
  Title,
  Group,
  Select,
  TextInput,
  Table,
  Badge,
  Text,
  Paper,
  ActionIcon,
  Tooltip,
  Code,
} from '@mantine/core';
import { IconSearch, IconDownload, IconTrash } from '@tabler/icons-react';
import { notifications } from '@mantine/notifications';
import { useLogs, useClearLogs } from '@/api/queries';
import type { LogLevel } from '@/api/types';

export function LogsContent() {
  const [level, setLevel] = useState<LogLevel | ''>('');
  const [search, setSearch] = useState('');
  const { data } = useLogs(level || undefined, search || undefined);
  const clearLogs = useClearLogs();

  const levelOptions = [
    { value: '', label: 'All Levels' },
    { value: 'DEBUG', label: 'DEBUG' },
    { value: 'INFO', label: 'INFO' },
    { value: 'WARN', label: 'WARN' },
    { value: 'ERROR', label: 'ERROR' },
  ];

  const handleClear = async () => {
    if (!confirm('Are you sure you want to clear all logs?')) {
      return;
    }

    try {
      await clearLogs.mutateAsync();
      notifications.show({
        title: 'Success',
        message: 'Logs cleared',
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to clear logs',
        color: 'red',
      });
    }
  };

  const handleExport = () => {
    const jsonData = JSON.stringify(data.entries, null, 2);
    const blob = new Blob([jsonData], { type: 'application/json' });
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `logs-${new Date().toISOString()}.json`;
    document.body.appendChild(a);
    a.click();
    window.URL.revokeObjectURL(url);
    document.body.removeChild(a);

    notifications.show({
      title: 'Success',
      message: 'Logs exported',
      color: 'green',
    });
  };

  const getLevelColor = (lvl: LogLevel) => {
    switch (lvl) {
      case 'ERROR':
        return 'red';
      case 'WARN':
        return 'orange';
      case 'INFO':
        return 'blue';
      case 'DEBUG':
        return 'gray';
      default:
        return 'gray';
    }
  };

  const formatTimestamp = (timestamp: string) => {
    const date = new Date(timestamp);
    const timeStr = date.toLocaleTimeString('en-US', {
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit',
    });
    const ms = date.getMilliseconds().toString().padStart(3, '0');
    return `${timeStr}.${ms}`;
  };

  return (
    <Stack gap="md">
      {/* Header */}
      <Group justify="space-between">
        <Title order={2}>Application Logs</Title>
        <Group gap="xs">
          <Tooltip label="Export logs">
            <ActionIcon variant="filled" color="blue" size="lg" onClick={handleExport}>
              <IconDownload size={20} />
            </ActionIcon>
          </Tooltip>
          <Tooltip label="Clear logs">
            <ActionIcon
              variant="filled"
              color="gray"
              size="lg"
              onClick={handleClear}
              loading={clearLogs.isPending}
            >
              <IconTrash size={20} />
            </ActionIcon>
          </Tooltip>
        </Group>
      </Group>

      {/* Filters */}
      <Group justify="space-between">
        <Group gap="md">
          <Text size="sm" fw={500}>
            Filter:
          </Text>
          <Select
            value={level}
            onChange={(value) => setLevel((value as LogLevel) || '')}
            data={levelOptions}
            size="sm"
            style={{ width: 150 }}
          />
          <TextInput
            placeholder="Search..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            leftSection={<IconSearch size={16} />}
            size="sm"
            style={{ width: 300 }}
          />
        </Group>
        <Text size="sm" c="dimmed">
          {data.count} entries
        </Text>
      </Group>

      {/* Table */}
      <Paper withBorder>
        <Table>
          <Table.Thead>
            <Table.Tr>
              <Table.Th>TIME</Table.Th>
              <Table.Th>LEVEL</Table.Th>
              <Table.Th>MESSAGE</Table.Th>
              <Table.Th>FIELDS</Table.Th>
            </Table.Tr>
          </Table.Thead>
          <Table.Tbody>
            {data.entries.length === 0 && (
              <Table.Tr>
                <Table.Td colSpan={4}>
                  <Text ta="center" c="dimmed" py="xl">
                    No log entries
                  </Text>
                </Table.Td>
              </Table.Tr>
            )}
            {data.entries.map((entry, index) => (
              <Table.Tr key={index}>
                <Table.Td>
                  <Text size="sm" ff="monospace">
                    {formatTimestamp(entry.timestamp)}
                  </Text>
                </Table.Td>
                <Table.Td>
                  <Badge size="sm" color={getLevelColor(entry.level)}>
                    {entry.level}
                  </Badge>
                </Table.Td>
                <Table.Td>
                  <Text size="sm">{entry.message}</Text>
                </Table.Td>
                <Table.Td>
                  {entry.fields && Object.keys(entry.fields).length > 0 && (
                    <Code block>{JSON.stringify(entry.fields, null, 2)}</Code>
                  )}
                </Table.Td>
              </Table.Tr>
            ))}
          </Table.Tbody>
        </Table>
      </Paper>
    </Stack>
  );
}
