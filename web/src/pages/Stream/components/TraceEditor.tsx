import { useState } from 'react';
import {
  Stack,
  Title,
  Select,
  Group,
  Button,
  Paper,
  Text,
  TextInput,
  NumberInput,
  Switch,
  Divider,
  Badge,
  ScrollArea,
  Modal,
  Tabs,
  MultiSelect,
} from '@mantine/core';
import { IconEdit, IconGitMerge, IconDownload, IconRefresh } from '@tabler/icons-react';
import { notifications } from '@mantine/notifications';
import {
  useTraceFiles,
  useTraceFilePreview,
  useEditTrace,
  useMergeTraces,
  useExportTrace,
} from '@/api/queries';
import type {
  EditTraceRequest,
  MergeTracesRequest,
  ExportTraceRequest,
  TraceFileInfo,
} from '@/api/types';

export function TraceEditor() {
  const { data: files } = useTraceFiles();
  const editTrace = useEditTrace();
  const mergeTraces = useMergeTraces();
  const exportTrace = useExportTrace();

  // Edit tab state
  const [sourceFile, setSourceFile] = useState<string>('');
  const [outputFile, setOutputFile] = useState<string>('');
  const [direction, setDirection] = useState<string>('');
  const [startTime, setStartTime] = useState<number | undefined>();
  const [endTime, setEndTime] = useState<number | undefined>();
  const [messageTypes, setMessageTypes] = useState<string[]>([]);
  const [containsText, setContainsText] = useState<string>('');
  const [normalizeTime, setNormalizeTime] = useState(false);
  const [remapProxyName, setRemapProxyName] = useState<string>('');
  const [timestampOffset, setTimestampOffset] = useState<number | undefined>();

  // Merge tab state
  const [mergeSourceFiles, setMergeSourceFiles] = useState<string[]>([]);
  const [mergeOutputFile, setMergeOutputFile] = useState<string>('');
  const [sortByTime, setSortByTime] = useState(true);
  const [normalize, setNormalize] = useState(false);

  // Export tab state
  const [exportSourceFile, setExportSourceFile] = useState<string>('');
  const [exportFormat, setExportFormat] = useState<'json' | 'jsonl'>('jsonl');
  const [exportDirection, setExportDirection] = useState<string>('');
  const [exportLimit, setExportLimit] = useState<number | undefined>();

  // Preview state
  const [previewFile, setPreviewFile] = useState<string | null>(null);
  const [previewOpen, setPreviewOpen] = useState(false);

  const { data: previewData } = useTraceFilePreview(previewFile, {
    limit: 50,
    enabled: previewOpen && !!previewFile,
  });

  const fileOptions = files.map((f: TraceFileInfo) => ({
    value: f.name,
    label: `${f.name} (${f.proxyName}, ${(f.size / 1024).toFixed(1)} KB)`,
  }));

  const messageTypeOptions = [
    { value: '10', label: 'LINK (10)' },
    { value: '11', label: 'INIT (11)' },
    { value: '30', label: 'INVOKE (30)' },
    { value: '31', label: 'INVOKE_REPLY (31)' },
    { value: '40', label: 'SIGNAL (40)' },
    { value: '50', label: 'PROPERTY_CHANGE (50)' },
  ];

  const handleEdit = async () => {
    if (!sourceFile || !outputFile) {
      notifications.show({
        title: 'Validation Error',
        message: 'Source file and output file are required',
        color: 'red',
      });
      return;
    }

    const request: EditTraceRequest = {
      sourceFile,
      outputFile,
      direction: direction || undefined,
      startTime: startTime || undefined,
      endTime: endTime || undefined,
      messageTypes: messageTypes.length > 0 ? messageTypes.map(Number) : undefined,
      containsText: containsText || undefined,
      normalizeTime,
      remapProxyName: remapProxyName || undefined,
      timestampOffset: timestampOffset || undefined,
    };

    try {
      await editTrace.mutateAsync(request);
      notifications.show({
        title: 'Success',
        message: `Trace edited and saved to ${outputFile}`,
        color: 'green',
      });
      // Reset form
      setOutputFile('');
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to edit trace',
        color: 'red',
      });
    }
  };

  const handleMerge = async () => {
    if (mergeSourceFiles.length < 2 || !mergeOutputFile) {
      notifications.show({
        title: 'Validation Error',
        message: 'At least 2 source files and output file are required',
        color: 'red',
      });
      return;
    }

    const request: MergeTracesRequest = {
      sourceFiles: mergeSourceFiles,
      outputFile: mergeOutputFile,
      sortByTime,
      normalize,
    };

    try {
      await mergeTraces.mutateAsync(request);
      notifications.show({
        title: 'Success',
        message: `Traces merged and saved to ${mergeOutputFile}`,
        color: 'green',
      });
      // Reset form
      setMergeSourceFiles([]);
      setMergeOutputFile('');
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to merge traces',
        color: 'red',
      });
    }
  };

  const handleExport = async () => {
    if (!exportSourceFile) {
      notifications.show({
        title: 'Validation Error',
        message: 'Source file is required',
        color: 'red',
      });
      return;
    }

    const request: ExportTraceRequest = {
      sourceFile: exportSourceFile,
      format: exportFormat,
      direction: exportDirection || undefined,
      limit: exportLimit || undefined,
    };

    try {
      const blob = await exportTrace.mutateAsync(request);

      // Create download link
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = `${exportSourceFile}.${exportFormat}`;
      document.body.appendChild(a);
      a.click();
      window.URL.revokeObjectURL(url);
      document.body.removeChild(a);

      notifications.show({
        title: 'Success',
        message: `Trace exported as ${exportFormat}`,
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to export trace',
        color: 'red',
      });
    }
  };

  const handlePreview = (filename: string) => {
    setPreviewFile(filename);
    setPreviewOpen(true);
  };

  return (
    <Stack gap="md">
      <Group justify="space-between">
        <Title order={2}>Stream Editor</Title>
        <Text c="dimmed" size="sm">
          Edit, merge, and export trace files
        </Text>
      </Group>

      <Tabs defaultValue="edit">
        <Tabs.List>
          <Tabs.Tab value="edit" leftSection={<IconEdit size={16} />}>
            Edit Trace
          </Tabs.Tab>
          <Tabs.Tab value="merge" leftSection={<IconGitMerge size={16} />}>
            Merge Traces
          </Tabs.Tab>
          <Tabs.Tab value="export" leftSection={<IconDownload size={16} />}>
            Export Trace
          </Tabs.Tab>
        </Tabs.List>

        <Tabs.Panel value="edit" pt="md">
          <Paper p="md" withBorder>
            <Stack gap="md">
              <Title order={4}>Source & Output</Title>
              <Group grow>
                <Select
                  label="Source File"
                  placeholder="Select trace file"
                  data={fileOptions}
                  value={sourceFile}
                  onChange={(value) => setSourceFile(value || '')}
                  searchable
                  required
                />
                <TextInput
                  label="Output File"
                  placeholder="output-trace.jsonl"
                  value={outputFile}
                  onChange={(e) => setOutputFile(e.currentTarget.value)}
                  required
                />
              </Group>

              <Divider label="Filters" />

              <Group grow>
                <Select
                  label="Direction"
                  placeholder="All directions"
                  data={[
                    { value: 'SEND', label: 'SEND' },
                    { value: 'RECV', label: 'RECV' },
                  ]}
                  value={direction}
                  onChange={(value) => setDirection(value || '')}
                  clearable
                />
                <MultiSelect
                  label="Message Types"
                  placeholder="All types"
                  data={messageTypeOptions}
                  value={messageTypes}
                  onChange={setMessageTypes}
                  clearable
                />
              </Group>

              <Group grow>
                <NumberInput
                  label="Start Time (Unix ms)"
                  placeholder="Start timestamp"
                  value={startTime}
                  onChange={(value) => setStartTime(value as number | undefined)}
                  hideControls
                />
                <NumberInput
                  label="End Time (Unix ms)"
                  placeholder="End timestamp"
                  value={endTime}
                  onChange={(value) => setEndTime(value as number | undefined)}
                  hideControls
                />
              </Group>

              <TextInput
                label="Contains Text"
                placeholder="Search in messages"
                value={containsText}
                onChange={(e) => setContainsText(e.currentTarget.value)}
              />

              <Divider label="Transformations" />

              <Group>
                <Switch
                  label="Normalize timestamps"
                  description="Start from time 0"
                  checked={normalizeTime}
                  onChange={(e) => setNormalizeTime(e.currentTarget.checked)}
                />
              </Group>

              <Group grow>
                <TextInput
                  label="Remap Proxy Name"
                  placeholder="New proxy name"
                  value={remapProxyName}
                  onChange={(e) => setRemapProxyName(e.currentTarget.value)}
                />
                <NumberInput
                  label="Timestamp Offset (ms)"
                  placeholder="Add/subtract milliseconds"
                  value={timestampOffset}
                  onChange={(value) => setTimestampOffset(value as number | undefined)}
                  hideControls
                />
              </Group>

              <Group justify="space-between" mt="md">
                <Button
                  variant="light"
                  leftSection={<IconRefresh size={16} />}
                  onClick={() => handlePreview(sourceFile)}
                  disabled={!sourceFile}
                >
                  Preview Source
                </Button>
                <Button
                  leftSection={<IconEdit size={16} />}
                  onClick={handleEdit}
                  loading={editTrace.isPending}
                >
                  Apply & Save
                </Button>
              </Group>
            </Stack>
          </Paper>
        </Tabs.Panel>

        <Tabs.Panel value="merge" pt="md">
          <Paper p="md" withBorder>
            <Stack gap="md">
              <Title order={4}>Merge Multiple Traces</Title>

              <MultiSelect
                label="Source Files"
                placeholder="Select at least 2 files"
                data={fileOptions}
                value={mergeSourceFiles}
                onChange={setMergeSourceFiles}
                searchable
                required
              />

              <TextInput
                label="Output File"
                placeholder="merged-trace.jsonl"
                value={mergeOutputFile}
                onChange={(e) => setMergeOutputFile(e.currentTarget.value)}
                required
              />

              <Group>
                <Switch
                  label="Sort by timestamp"
                  description="Chronological order"
                  checked={sortByTime}
                  onChange={(e) => setSortByTime(e.currentTarget.checked)}
                />
                <Switch
                  label="Normalize timestamps"
                  description="Start from time 0"
                  checked={normalize}
                  onChange={(e) => setNormalize(e.currentTarget.checked)}
                />
              </Group>

              <Group justify="flex-end" mt="md">
                <Button
                  leftSection={<IconGitMerge size={16} />}
                  onClick={handleMerge}
                  loading={mergeTraces.isPending}
                >
                  Merge Traces
                </Button>
              </Group>
            </Stack>
          </Paper>
        </Tabs.Panel>

        <Tabs.Panel value="export" pt="md">
          <Paper p="md" withBorder>
            <Stack gap="md">
              <Title order={4}>Export Trace</Title>

              <Select
                label="Source File"
                placeholder="Select trace file"
                data={fileOptions}
                value={exportSourceFile}
                onChange={(value) => setExportSourceFile(value || '')}
                searchable
                required
              />

              <Select
                label="Export Format"
                data={[
                  { value: 'json', label: 'JSON (single array)' },
                  { value: 'jsonl', label: 'JSONL (line-delimited)' },
                ]}
                value={exportFormat}
                onChange={(value) => setExportFormat(value as 'json' | 'jsonl')}
              />

              <Group grow>
                <Select
                  label="Direction Filter"
                  placeholder="All directions"
                  data={[
                    { value: 'SEND', label: 'SEND' },
                    { value: 'RECV', label: 'RECV' },
                  ]}
                  value={exportDirection}
                  onChange={(value) => setExportDirection(value || '')}
                  clearable
                />
                <NumberInput
                  label="Limit Entries"
                  placeholder="No limit"
                  value={exportLimit}
                  onChange={(value) => setExportLimit(value as number | undefined)}
                  min={1}
                />
              </Group>

              <Group justify="space-between" mt="md">
                <Button
                  variant="light"
                  leftSection={<IconRefresh size={16} />}
                  onClick={() => handlePreview(exportSourceFile)}
                  disabled={!exportSourceFile}
                >
                  Preview
                </Button>
                <Button
                  leftSection={<IconDownload size={16} />}
                  onClick={handleExport}
                  loading={exportTrace.isPending}
                >
                  Download
                </Button>
              </Group>
            </Stack>
          </Paper>
        </Tabs.Panel>
      </Tabs>

      {/* Preview Modal */}
      <Modal
        opened={previewOpen}
        onClose={() => setPreviewOpen(false)}
        title={`Preview: ${previewFile}`}
        size="xl"
      >
        {previewData && (
          <Stack gap="sm">
            <Group>
              <Badge>Total: {previewData.count} entries</Badge>
            </Group>
            <ScrollArea h={500}>
              <Stack gap="xs">
                {previewData.entries.slice(0, 50).map((entry, idx) => (
                  <Paper key={idx} p="xs" withBorder>
                    <Group justify="space-between" mb="xs">
                      <Badge color={entry.dir === 'SEND' ? 'blue' : 'green'}>
                        {entry.dir}
                      </Badge>
                      <Text size="xs" c="dimmed">
                        {new Date(entry.ts).toLocaleString()}
                      </Text>
                    </Group>
                    <Text size="xs" ff="monospace">
                      {JSON.stringify(entry.msg, null, 2)}
                    </Text>
                  </Paper>
                ))}
                {previewData.entries.length > 50 && (
                  <Text size="sm" c="dimmed" ta="center">
                    Showing first 50 of {previewData.count} entries
                  </Text>
                )}
              </Stack>
            </ScrollArea>
          </Stack>
        )}
      </Modal>
    </Stack>
  );
}
