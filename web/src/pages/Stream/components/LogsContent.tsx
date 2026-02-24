import { useState } from 'react';
import {
  Stack,
  Title,
  Group,
  Select,
  TextInput,
  Badge,
  Text,
  ActionIcon,
  Tooltip,
  JsonInput,
  Paper,
} from '@mantine/core';
import { DataTable } from 'mantine-datatable';
import { IconSearch, IconDownload, IconTrash } from '@tabler/icons-react';
import { notifications } from '@mantine/notifications';
import { useLogs, useClearLogs } from '@/api/queries';
import type { LogLevel, LogEntry } from '@/api/types';

export function LogsContent() {
  const [level, setLevel] = useState<LogLevel | ''>('');
  const [search, setSearch] = useState('');
  const [page, setPage] = useState(1);
  const pageSize = 50;
  const { data } = useLogs(level || undefined, search || undefined);
  const clearLogs = useClearLogs();

  // Calculate pagination
  const totalRecords = data.entries.length;
  const startIndex = (page - 1) * pageSize;
  const endIndex = startIndex + pageSize;
  const paginatedEntries = data.entries.slice(startIndex, endIndex);

  // Reset to page 1 when filters change
  const handleLevelChange = (value: string | null) => {
    setLevel((value as LogLevel) || '');
    setPage(1);
  };

  const handleSearchChange = (value: string) => {
    setSearch(value);
    setPage(1);
  };

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
      hour12: false,
    });
    const ms = date.getMilliseconds().toString().padStart(3, '0');
    return `${timeStr}.${ms}`;
  };

  const formatFieldsCompact = (fields: Record<string, unknown> | undefined) => {
    if (!fields || Object.keys(fields).length === 0) {
      return '-';
    }
    // Single line JSON
    return JSON.stringify(fields);
  };

  const formatFieldsExpanded = (fields: Record<string, unknown> | undefined) => {
    if (!fields || Object.keys(fields).length === 0) {
      return '';
    }
    // Pretty printed JSON
    return JSON.stringify(fields, null, 2);
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
            onChange={handleLevelChange}
            data={levelOptions}
            size="sm"
            style={{ width: 150 }}
          />
          <TextInput
            placeholder="Search..."
            value={search}
            onChange={(e) => handleSearchChange(e.target.value)}
            leftSection={<IconSearch size={16} />}
            size="sm"
            style={{ width: 300 }}
          />
        </Group>
        <Text size="sm" c="dimmed">
          {totalRecords} entries
          {totalRecords > pageSize && ` (showing ${startIndex + 1}-${Math.min(endIndex, totalRecords)})`}
        </Text>
      </Group>

      {/* DataTable */}
      <DataTable
        withTableBorder
        borderRadius="sm"
        striped
        highlightOnHover
        records={paginatedEntries}
        totalRecords={totalRecords}
        recordsPerPage={pageSize}
        page={page}
        onPageChange={setPage}
        columns={[
          {
            accessor: 'timestamp',
            title: 'Time',
            width: 120,
            render: (entry: LogEntry) => (
              <Text size="xs" ff="monospace">
                {formatTimestamp(entry.timestamp)}
              </Text>
            ),
          },
          {
            accessor: 'level',
            title: 'Level',
            width: 80,
            render: (entry: LogEntry) => (
              <Badge size="sm" color={getLevelColor(entry.level)}>
                {entry.level}
              </Badge>
            ),
          },
          {
            accessor: 'message',
            title: 'Message',
            width: 300,
            render: (entry: LogEntry) => (
              <Text size="sm" lineClamp={1}>
                {entry.message}
              </Text>
            ),
          },
          {
            accessor: 'fields',
            title: 'Fields',
            render: (entry: LogEntry) => {
              const compactJson = formatFieldsCompact(entry.fields);
              const expandedJson = formatFieldsExpanded(entry.fields);

              if (compactJson === '-') {
                return (
                  <Text size="xs" c="dimmed">
                    -
                  </Text>
                );
              }

              return (
                <Tooltip
                  label={
                    <Paper withBorder shadow="md" p="xs" style={{ width: 500 }}>
                      <JsonInput
                        value={expandedJson}
                        readOnly
                        autosize
                        minRows={3}
                        maxRows={15}
                        styles={{
                          input: {
                            fontSize: '11px',
                            fontFamily: 'monospace',
                            padding: '8px',
                            border: 'none',
                            backgroundColor: 'transparent',
                          },
                        }}
                      />
                    </Paper>
                  }
                  multiline
                  withArrow={false}
                  position="left"
                  offset={10}
                  styles={{
                    tooltip: {
                      padding: 0,
                      maxWidth: 'none',
                      backgroundColor: 'transparent',
                      border: 'none',
                    },
                  }}
                >
                  <Text
                    size="xs"
                    ff="monospace"
                    c="dimmed"
                    style={{
                      cursor: 'pointer',
                      maxWidth: '100%',
                      overflow: 'hidden',
                      textOverflow: 'ellipsis',
                      whiteSpace: 'nowrap',
                    }}
                  >
                    {compactJson}
                  </Text>
                </Tooltip>
              );
            },
          },
        ]}
        noRecordsText="No log entries"
        minHeight={400}
        verticalSpacing="xs"
      />
    </Stack>
  );
}
