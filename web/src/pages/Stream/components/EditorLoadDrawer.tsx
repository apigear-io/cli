import { Drawer, Stack, FileButton, Button, Text, ScrollArea, Paper, Group, Badge } from '@mantine/core';
import { IconUpload, IconFileText, IconRefresh } from '@tabler/icons-react';
import { useTraceFiles } from '@/api/queries';
import { useEditorLoad } from '@/api/queries';
import { useEditorContext } from './EditorContext';
import { notifications } from '@mantine/notifications';

interface EditorLoadDrawerProps {
  opened: boolean;
  onClose: () => void;
}

export function EditorLoadDrawer({ opened, onClose }: EditorLoadDrawerProps) {
  const { data: traceFiles, refetch } = useTraceFiles();
  const loadMutation = useEditorLoad();
  const { setSessionStats } = useEditorContext();

  const handleFileUpload = async (file: File | null) => {
    if (!file) return;

    try {
      const stats = await loadMutation.mutateAsync({ file });
      setSessionStats(stats);
      onClose();
      notifications.show({
        title: 'Trace Loaded',
        message: `Successfully loaded ${file.name}`,
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Upload Failed',
        message: error instanceof Error ? error.message : 'Failed to upload file',
        color: 'red',
      });
    }
  };

  const handleServerFile = async (filename: string) => {
    try {
      const stats = await loadMutation.mutateAsync({ name: filename });
      setSessionStats(stats);
      onClose();
      notifications.show({
        title: 'Trace Loaded',
        message: `Successfully loaded ${filename}`,
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Load Failed',
        message: error instanceof Error ? error.message : 'Failed to load file',
        color: 'red',
      });
    }
  };

  return (
    <Drawer opened={opened} onClose={onClose} title="Set Stream" size="lg" position="right">
      <Stack gap="lg">
        {/* Upload section */}
        <Paper withBorder p="xl" style={{ borderStyle: 'dashed' }}>
          <Stack align="center" gap="md">
            <IconUpload size={48} stroke={1.5} style={{ opacity: 0.5 }} />
            <Text size="sm" c="dimmed" ta="center">
              Drop JSONL trace file here or click to browse
            </Text>
            <FileButton onChange={handleFileUpload} accept=".jsonl,.jsonl.gz">
              {(props) => (
                <Button {...props} loading={loadMutation.isPending} leftSection={<IconUpload size={16} />}>
                  Browse Files
                </Button>
              )}
            </FileButton>
          </Stack>
        </Paper>

        {/* Divider */}
        <Text c="dimmed" ta="center" fw={500}>
          OR SELECT FROM TRACES
        </Text>

        {/* Server files list */}
        <Stack gap="xs">
          <Group justify="space-between">
            <Text fw={500}>Recent Trace Files</Text>
            <Button
              variant="subtle"
              size="xs"
              leftSection={<IconRefresh size={14} />}
              onClick={() => refetch()}
            >
              Refresh
            </Button>
          </Group>
          <ScrollArea h={400}>
            {traceFiles.length === 0 ? (
              <Paper p="md" withBorder>
                <Text size="sm" c="dimmed" ta="center">
                  No trace files found
                </Text>
              </Paper>
            ) : (
              traceFiles.map((file) => (
                <Paper
                  key={file.name}
                  p="md"
                  mb="xs"
                  withBorder
                  style={{ cursor: 'pointer' }}
                  onClick={() => handleServerFile(file.name)}
                >
                  <Group justify="space-between">
                    <Group>
                      <IconFileText size={20} />
                      <Stack gap={0}>
                        <Text size="sm" fw={500}>
                          {file.name}
                        </Text>
                        <Text size="xs" c="dimmed">
                          {file.proxyName}
                        </Text>
                      </Stack>
                    </Group>
                    <Badge>{(file.size / 1024).toFixed(1)} KB</Badge>
                  </Group>
                </Paper>
              ))
            )}
          </ScrollArea>
        </Stack>
      </Stack>
    </Drawer>
  );
}
