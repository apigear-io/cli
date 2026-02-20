import { useState } from 'react';
import {
  Stack,
  Title,
  Text,
  Tabs,
  Paper,
  Group,
  TextInput,
  Switch,
  NumberInput,
  Alert,
  Code,
  Divider,
} from '@mantine/core';
import {
  IconSettings,
  IconFileText,
  IconDatabase,
  IconInfoCircle,
  IconAlertTriangle,
} from '@tabler/icons-react';

export function SettingsContent() {
  const [activeTab, setActiveTab] = useState<string | null>('general');

  // Default values (these would come from backend in a real implementation)
  const settings = {
    listenAddress: ':8080',
    traceDirectory: './data/traces',
    configFile: './stream.yaml',
    verbose: false,
    trace: true,
    watchConfig: false,
    messageBatching: true,
    batchInterval: 50,
    maxFileSizeMB: 10,
    maxBackups: 5,
    maxAgeDays: 7,
    compress: true,
    bufferSize: 1000,
    currentMessages: 0,
  };

  return (
    <Stack gap="md">
      {/* Header */}
      <Group justify="space-between">
        <Stack gap={4}>
          <Title order={2}>Settings</Title>
          <Text size="sm" c="dimmed">
            Configuration is managed via config file or CLI flags
          </Text>
        </Stack>
      </Group>

      <Alert icon={<IconInfoCircle />} color="blue">
        Settings are read from the configuration file (stream.yaml) or CLI flags.
        Most changes require restarting the server to take effect.
      </Alert>

      {/* Tabs */}
      <Tabs value={activeTab} onChange={setActiveTab}>
        <Tabs.List>
          <Tabs.Tab value="general" leftSection={<IconSettings size={16} />}>
            General
          </Tabs.Tab>
          <Tabs.Tab value="traces" leftSection={<IconFileText size={16} />}>
            Traces
          </Tabs.Tab>
          <Tabs.Tab value="buffer" leftSection={<IconDatabase size={16} />}>
            Buffer
          </Tabs.Tab>
        </Tabs.List>

        {/* General Tab */}
        <Tabs.Panel value="general" pt="md">
          <Stack gap="lg">
            {/* Application */}
            <Paper withBorder p="md">
              <Stack gap="md">
                <Text fw={600} size="lg">
                  Application
                </Text>

                <TextInput
                  label="Web UI Listen Address"
                  value={settings.listenAddress}
                  description="Address for the web UI (e.g., :8080 or 0.0.0.0:8080)"
                  readOnly
                />

                <TextInput
                  label="Trace Directory"
                  value={settings.traceDirectory}
                  description="Directory for trace files (relative to executable)"
                  readOnly
                />

                <TextInput
                  label="Config File"
                  value={settings.configFile}
                  description="Path to YAML configuration file"
                  readOnly
                />
              </Stack>
            </Paper>

            {/* Logging & Behavior */}
            <Paper withBorder p="md">
              <Stack gap="md">
                <Text fw={600} size="lg">
                  Logging & Behavior
                </Text>

                <Switch
                  label="Verbose Mode"
                  description="Log all WebSocket messages to console output"
                  checked={settings.verbose}
                  disabled
                />

                <Switch
                  label="Trace Mode"
                  description="Write all messages to JSONL files for later analysis"
                  checked={settings.trace}
                  disabled
                />

                <Switch
                  label="Watch Config File"
                  description="Automatically reload when config file changes on disk"
                  checked={settings.watchConfig}
                  disabled
                />

                <Divider />

                <Text size="sm" fw={500} c="dimmed">
                  Performance
                </Text>

                <Switch
                  label="Message Batching"
                  description="Batch SSE messages for better performance at high throughput"
                  checked={settings.messageBatching}
                  disabled
                />

                <NumberInput
                  label="Batch Interval (ms)"
                  value={settings.batchInterval}
                  description="Lower = more responsive, Higher = better performance (10-200ms)"
                  min={10}
                  max={200}
                  disabled
                />
              </Stack>
            </Paper>

            <Alert icon={<IconAlertTriangle />} color="orange" title="Configuration">
              To change these settings, edit the <Code>stream.yaml</Code> configuration file
              or use CLI flags when starting the server. Run <Code>apigear stream --help</Code>{' '}
              to see available options.
            </Alert>
          </Stack>
        </Tabs.Panel>

        {/* Traces Tab */}
        <Tabs.Panel value="traces" pt="md">
          <Stack gap="lg">
            {/* Trace File Rotation */}
            <Paper withBorder p="md">
              <Stack gap="md">
                <Text fw={600} size="lg">
                  Trace File Rotation
                </Text>
                <Text size="sm" c="dimmed">
                  Configure how trace files are rotated and managed. These settings control the
                  lumberjack log rotation.
                </Text>

                <NumberInput
                  label="Max File Size (MB)"
                  value={settings.maxFileSizeMB}
                  description="Maximum size in MB before rotation (default: 10)"
                  min={1}
                  max={1000}
                  disabled
                />

                <NumberInput
                  label="Max Backups"
                  value={settings.maxBackups}
                  description="Maximum number of old files to keep (default: 5, 0 = keep all)"
                  min={0}
                  max={100}
                  disabled
                />

                <NumberInput
                  label="Max Age (Days)"
                  value={settings.maxAgeDays}
                  description="Maximum age in days to keep old files (default: 7, 0 = no limit)"
                  min={0}
                  max={365}
                  disabled
                />

                <Switch
                  label="Compress Rotated Files"
                  description="Gzip compress rotated trace files to save disk space"
                  checked={settings.compress}
                  disabled
                />
              </Stack>
            </Paper>

            {/* About Trace Files */}
            <Alert icon={<IconInfoCircle />} color="blue" title="JSONL Format">
              <Text size="sm">
                Trace files are written in JSONL format (one JSON object per line). Each entry
                contains timestamp, direction, proxy name, and message content.
              </Text>
            </Alert>

            <Alert icon={<IconAlertTriangle />} color="orange" title="Restart Required">
              <Text size="sm">
                Changes to trace rotation settings will take effect after saving and restarting the
                proxies. Active trace files will continue using the old settings until rotated.
              </Text>
            </Alert>
          </Stack>
        </Tabs.Panel>

        {/* Buffer Tab */}
        <Tabs.Panel value="buffer" pt="md">
          <Stack gap="lg">
            {/* Message Buffer */}
            <Paper withBorder p="md">
              <Stack gap="md">
                <Text fw={600} size="lg">
                  Message Buffer
                </Text>
                <Text size="sm" c="dimmed">
                  The server keeps recent messages in memory for the live stream view. This allows
                  you to see message history when you first open the page or refresh.
                </Text>

                <Alert icon={<IconInfoCircle />} color="blue">
                  <Text size="sm" fw={500}>
                    Buffer Size: {settings.bufferSize} messages
                  </Text>
                  <Text size="sm" mt="xs">
                    Messages are stored in memory and will be lost when the server restarts.
                  </Text>
                </Alert>
              </Stack>
            </Paper>

            {/* Buffer Statistics */}
            <Paper withBorder p="md">
              <Stack gap="md">
                <Text fw={600} size="lg">
                  Buffer Statistics
                </Text>

                <Group>
                  <Stack gap={4} flex={1}>
                    <Text size="sm" c="dimmed">
                      Current Messages
                    </Text>
                    <Text size="xl" fw={600}>
                      {settings.currentMessages}
                    </Text>
                  </Stack>

                  <Stack gap={4} flex={1}>
                    <Text size="sm" c="dimmed">
                      Max Capacity
                    </Text>
                    <Text size="xl" fw={600}>
                      {settings.bufferSize}
                    </Text>
                  </Stack>

                  <Stack gap={4} flex={1}>
                    <Text size="sm" c="dimmed">
                      Usage
                    </Text>
                    <Text size="xl" fw={600}>
                      {((settings.currentMessages / settings.bufferSize) * 100).toFixed(1)}%
                    </Text>
                  </Stack>
                </Group>
              </Stack>
            </Paper>

            <Alert icon={<IconInfoCircle />} color="blue">
              <Text size="sm">
                The message buffer is implemented using a circular buffer. When full, the oldest
                messages are automatically discarded to make room for new ones.
              </Text>
            </Alert>
          </Stack>
        </Tabs.Panel>
      </Tabs>
    </Stack>
  );
}
