import { useState, Suspense } from 'react';
import {
  Stack,
  Title,
  Paper,
  Group,
  Text,
  Badge,
  Button,
  ActionIcon,
  Table,
  Modal,
  Select,
  Alert,
} from '@mantine/core';
import {
  IconTrash,
  IconEye,
  IconRefresh,
  IconFileText,
  IconDownload,
  IconFolder,
} from '@tabler/icons-react';
import { notifications } from '@mantine/notifications';
import { useTraceFiles, useTraceStats, useDeleteTraceFile } from '@/api/queries';
import { TraceViewer } from './TraceViewer';
import { LoadingFallback } from '@/components/LoadingFallback';
import { ErrorBoundary } from '@/components/ErrorBoundary';

export function TracesContent() {
  const { data: files } = useTraceFiles();
  const { data: stats } = useTraceStats();
  const deleteTrace = useDeleteTraceFile();

  const [selectedFile, setSelectedFile] = useState<string | null>(null);
  const [viewerOpen, setViewerOpen] = useState(false);
  const [proxyFilter, setProxyFilter] = useState<string>('all');
  const [dateFilter, setDateFilter] = useState<string>('all');

  // Get unique proxy names for filter
  const proxyOptions = [
    { value: 'all', label: 'All Proxies' },
    ...Array.from(new Set(files.map((f) => f.proxyName))).map((proxy) => ({
      value: proxy,
      label: proxy,
    })),
  ];

  // Filter files based on selected filters
  const filteredFiles = files.filter((file) => {
    if (proxyFilter !== 'all' && file.proxyName !== proxyFilter) return false;
    // Date filter can be enhanced later
    return true;
  });

  const handleView = (filename: string) => {
    setSelectedFile(filename);
    setViewerOpen(true);
  };

  const handleDelete = (filename: string) => {
    if (confirm(`Are you sure you want to delete "${filename}"?`)) {
      deleteTrace.mutate(filename, {
        onSuccess: () => {
          notifications.show({
            title: 'Success',
            message: `Trace file "${filename}" deleted`,
            color: 'green',
          });
        },
        onError: () => {
          notifications.show({
            title: 'Error',
            message: 'Failed to delete trace file',
            color: 'red',
          });
        },
      });
    }
  };

  const formatSize = (bytes: number): string => {
    if (bytes < 1024) return `${bytes} B`;
    if (bytes < 1024 * 1024) return `${(bytes / 1024).toFixed(1)} KB`;
    return `${(bytes / (1024 * 1024)).toFixed(1)} MB`;
  };

  const formatDate = (dateStr: string): string => {
    const date = new Date(dateStr);
    return date.toLocaleString();
  };

  return (
    <Stack gap="md">
      {/* Header */}
      <Group justify="space-between">
        <Title order={2}>Stream Files</Title>
        <Button
          leftSection={<IconRefresh size={16} />}
          variant="outline"
          size="sm"
          onClick={() => window.location.reload()}
        >
          Refresh
        </Button>
      </Group>

      {/* Trace Directory Info */}
      <Alert icon={<IconFolder size={16} />} color="blue" variant="light">
        <Text size="sm">
          <strong>Trace Directory:</strong> {stats.traceDir}
        </Text>
      </Alert>

      {/* Filters and File Count */}
      <Group justify="space-between">
        <Group gap="md">
          <Text size="sm" fw={500}>
            Filter:
          </Text>
          <Select
            value={proxyFilter}
            onChange={(value) => setProxyFilter(value || 'all')}
            data={proxyOptions}
            size="sm"
            style={{ width: 180 }}
          />
          <Select
            value={dateFilter}
            onChange={(value) => setDateFilter(value || 'all')}
            data={[
              { value: 'all', label: 'All Dates' },
              { value: 'today', label: 'Today' },
              { value: 'week', label: 'This Week' },
            ]}
            size="sm"
            style={{ width: 150 }}
          />
        </Group>
        <Badge size="lg" variant="light">
          {filteredFiles.length} files
        </Badge>
      </Group>

      {/* Files Table */}
      <Paper withBorder>
        <Table>
          <Table.Thead>
            <Table.Tr>
              <Table.Th>FILENAME</Table.Th>
              <Table.Th>PROXY</Table.Th>
              <Table.Th>DATE</Table.Th>
              <Table.Th>SIZE</Table.Th>
              <Table.Th>ACTIONS</Table.Th>
            </Table.Tr>
          </Table.Thead>
          <Table.Tbody>
            {filteredFiles.length === 0 && (
              <Table.Tr>
                <Table.Td colSpan={5}>
                  <Text ta="center" c="dimmed" py="xl">
                    No trace files found
                  </Text>
                </Table.Td>
              </Table.Tr>
            )}
            {filteredFiles.map((file) => (
              <Table.Tr key={file.name}>
                <Table.Td>
                  <Group gap="xs">
                    <IconFileText size={16} color="var(--mantine-color-gray-6)" />
                    <Text
                      ff="monospace"
                      size="sm"
                      c="blue"
                      style={{ cursor: 'pointer' }}
                      onClick={() => handleView(file.name)}
                    >
                      {file.name}
                    </Text>
                  </Group>
                </Table.Td>
                <Table.Td>
                  <Text size="sm">{file.proxyName}</Text>
                </Table.Td>
                <Table.Td>
                  <Text size="sm">{formatDate(file.modTime)}</Text>
                </Table.Td>
                <Table.Td>
                  <Text size="sm">{formatSize(file.size)}</Text>
                </Table.Td>
                <Table.Td>
                  <Group gap={4}>
                    <ActionIcon
                      variant="outline"
                      color="blue"
                      size="sm"
                      onClick={() => handleView(file.name)}
                      title="View trace"
                    >
                      <IconEye size={14} />
                    </ActionIcon>
                    <ActionIcon
                      variant="outline"
                      color="green"
                      size="sm"
                      onClick={() => window.open(`/api/v1/stream/traces/${encodeURIComponent(file.name)}`, '_blank')}
                      title="Download trace"
                    >
                      <IconDownload size={14} />
                    </ActionIcon>
                    <ActionIcon
                      variant="outline"
                      color="red"
                      size="sm"
                      onClick={() => handleDelete(file.name)}
                      title="Delete trace"
                    >
                      <IconTrash size={14} />
                    </ActionIcon>
                  </Group>
                </Table.Td>
              </Table.Tr>
            ))}
          </Table.Tbody>
        </Table>
      </Paper>

      {/* Viewer Modal */}
      <Modal
        opened={viewerOpen}
        onClose={() => setViewerOpen(false)}
        title={`Trace: ${selectedFile}`}
        size="xl"
      >
        {selectedFile && (
          <ErrorBoundary>
            <Suspense fallback={<LoadingFallback message="Loading trace entries..." />}>
              <TraceViewer filename={selectedFile} />
            </Suspense>
          </ErrorBoundary>
        )}
      </Modal>
    </Stack>
  );
}
