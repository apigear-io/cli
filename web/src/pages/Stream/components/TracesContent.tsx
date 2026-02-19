import { useState, Suspense } from 'react';
import {
  Stack,
  Title,
  Grid,
  Paper,
  Group,
  Text,
  Badge,
  Button,
  ActionIcon,
  Table,
  Modal,
} from '@mantine/core';
import { IconTrash, IconEye, IconRefresh, IconFileText } from '@tabler/icons-react';
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
      <Group justify="space-between">
        <Title order={2}>Trace Files</Title>
        <Button
          leftSection={<IconRefresh size={16} />}
          variant="default"
          size="sm"
          onClick={() => window.location.reload()}
        >
          Refresh
        </Button>
      </Group>

      {/* Stats Cards */}
      <Grid>
        <Grid.Col span={4}>
          <Paper withBorder p="md">
            <Stack gap="xs">
              <Group gap="xs">
                <IconFileText size={20} />
                <Text size="sm" c="dimmed">
                  Total Files
                </Text>
              </Group>
              <Text size="xl" fw={700}>
                {stats.fileCount}
              </Text>
            </Stack>
          </Paper>
        </Grid.Col>
        <Grid.Col span={4}>
          <Paper withBorder p="md">
            <Stack gap="xs">
              <Text size="sm" c="dimmed">
                Total Size
              </Text>
              <Text size="xl" fw={700}>
                {stats.totalMB.toFixed(2)} MB
              </Text>
            </Stack>
          </Paper>
        </Grid.Col>
        <Grid.Col span={4}>
          <Paper withBorder p="md">
            <Stack gap="xs">
              <Text size="sm" c="dimmed">
                Trace Directory
              </Text>
              <Text size="sm" ff="monospace">
                {stats.traceDir}
              </Text>
            </Stack>
          </Paper>
        </Grid.Col>
      </Grid>

      {/* Files Table */}
      <Paper withBorder>
        <Table>
          <Table.Thead>
            <Table.Tr>
              <Table.Th>Filename</Table.Th>
              <Table.Th>Proxy</Table.Th>
              <Table.Th>Size</Table.Th>
              <Table.Th>Modified</Table.Th>
              <Table.Th>Actions</Table.Th>
            </Table.Tr>
          </Table.Thead>
          <Table.Tbody>
            {files.length === 0 && (
              <Table.Tr>
                <Table.Td colSpan={5}>
                  <Text ta="center" c="dimmed" py="xl">
                    No trace files found
                  </Text>
                </Table.Td>
              </Table.Tr>
            )}
            {files.map((file) => (
              <Table.Tr key={file.name}>
                <Table.Td>
                  <Group gap="xs">
                    <IconFileText size={16} />
                    <Text ff="monospace" size="sm">
                      {file.name}
                    </Text>
                  </Group>
                </Table.Td>
                <Table.Td>
                  <Badge variant="light">{file.proxyName}</Badge>
                </Table.Td>
                <Table.Td>
                  <Text size="sm">{formatSize(file.size)}</Text>
                </Table.Td>
                <Table.Td>
                  <Text size="sm">{formatDate(file.modTime)}</Text>
                </Table.Td>
                <Table.Td>
                  <Group gap="xs">
                    <ActionIcon
                      variant="subtle"
                      color="blue"
                      onClick={() => handleView(file.name)}
                      title="View trace"
                    >
                      <IconEye size={16} />
                    </ActionIcon>
                    <ActionIcon
                      variant="subtle"
                      color="red"
                      onClick={() => handleDelete(file.name)}
                      title="Delete trace"
                    >
                      <IconTrash size={16} />
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
