import { useState } from 'react';
import {
  Stack,
  Title,
  Text,
  Group,
  Button,
  NumberInput,
  TextInput,
  Menu,
  Paper,
  Code,
  Textarea,
  ActionIcon,
  Tooltip,
} from '@mantine/core';
import {
  IconPlayerPlay,
  IconDeviceFloppy,
  IconFolderOpen,
  IconChevronDown,
  IconCodePlus,
  IconSparkles,
  IconReload,
} from '@tabler/icons-react';
import { notifications } from '@mantine/notifications';
import {
  useGeneratorPreview,
  useGeneratorSave,
  useGeneratorExamples,
  useGeneratorTemplates,
  useGeneratorLoadTemplate,
  useGeneratorSaveTemplate,
} from '@/api/queries';

const DEFAULT_TEMPLATE = `// Return a JS object - it's automatically JSON stringified
function generate() {
  return {
    ts: new Date().toISOString(),
    msg: [2, "demo.Counter/count", faker.int(0, 1000)]
  };
}`;

export function GeneratorContent() {
  const { data: examples } = useGeneratorExamples();
  const { data: templatesData } = useGeneratorTemplates();
  const [template, setTemplate] = useState(DEFAULT_TEMPLATE);
  const [count, setCount] = useState(100);
  const [filename, setFilename] = useState('generated.jsonl');
  const [proxyName, setProxyName] = useState('generator');
  const [preview, setPreview] = useState<unknown[] | null>(null);

  const previewMutation = useGeneratorPreview();
  const saveMutation = useGeneratorSave();
  const loadTemplateMutation = useGeneratorLoadTemplate();
  const saveTemplateMutation = useGeneratorSaveTemplate();

  const handlePreview = async () => {
    try {
      const result = await previewMutation.mutateAsync({
        template,
        count: Math.min(count, 100), // Limit preview to 100
      });
      setPreview(result.entries);
      notifications.show({
        title: 'Preview Generated',
        message: `Generated ${result.count} entries`,
        color: 'blue',
      });
    } catch (error) {
      notifications.show({
        title: 'Preview Failed',
        message: error instanceof Error ? error.message : 'Failed to generate preview',
        color: 'red',
      });
    }
  };

  const handleSave = async () => {
    try {
      const result = await saveMutation.mutateAsync({
        template,
        count,
        filename,
        proxyName,
      });
      notifications.show({
        title: 'Trace Saved',
        message: `Saved ${result.count} entries to ${result.filename}`,
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Save Failed',
        message: error instanceof Error ? error.message : 'Failed to save trace',
        color: 'red',
      });
    }
  };

  const handleLoadExample = (exampleName: string) => {
    const exampleTemplate = examples[exampleName];
    if (exampleTemplate) {
      setTemplate(exampleTemplate);
      setPreview(null);
      notifications.show({
        title: 'Example Loaded',
        message: `Loaded "${exampleName}" example`,
        color: 'blue',
      });
    }
  };

  const handleLoadTemplate = async (templateName: string) => {
    try {
      const result = await loadTemplateMutation.mutateAsync(templateName);
      setTemplate(result.template);
      setPreview(null);
      notifications.show({
        title: 'Template Loaded',
        message: `Loaded template "${templateName}"`,
        color: 'blue',
      });
    } catch (error) {
      notifications.show({
        title: 'Load Failed',
        message: error instanceof Error ? error.message : 'Failed to load template',
        color: 'red',
      });
    }
  };

  const handleSaveTemplate = async () => {
    const name = prompt('Enter template name:');
    if (!name) return;

    try {
      await saveTemplateMutation.mutateAsync({ name, template });
      notifications.show({
        title: 'Template Saved',
        message: `Saved template "${name}"`,
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Save Failed',
        message: error instanceof Error ? error.message : 'Failed to save template',
        color: 'red',
      });
    }
  };

  return (
    <Stack gap="md">
      {/* Header */}
      <Group justify="space-between">
        <Stack gap={4}>
          <Title order={2}>Trace Generator</Title>
          <Text size="sm" c="dimmed">
            Generate JSONL trace files using Go templates with faker functions
          </Text>
        </Stack>

        <Group gap="xs">
          <Menu>
            <Menu.Target>
              <Button leftSection={<IconFolderOpen size={16} />} variant="default" rightSection={<IconChevronDown size={14} />}>
                Load
              </Button>
            </Menu.Target>
            <Menu.Dropdown>
              <Menu.Label>Examples</Menu.Label>
              {Object.keys(examples).map((name) => (
                <Menu.Item key={name} onClick={() => handleLoadExample(name)}>
                  {name}
                </Menu.Item>
              ))}
              {templatesData.templates.length > 0 && (
                <>
                  <Menu.Divider />
                  <Menu.Label>Saved Templates</Menu.Label>
                  {templatesData.templates.map((name) => (
                    <Menu.Item key={name} onClick={() => handleLoadTemplate(name)}>
                      {name}
                    </Menu.Item>
                  ))}
                </>
              )}
            </Menu.Dropdown>
          </Menu>

          <Button
            leftSection={<IconDeviceFloppy size={16} />}
            variant="default"
            onClick={handleSaveTemplate}
            loading={saveTemplateMutation.isPending}
          >
            Save Template
          </Button>

          <Button
            leftSection={<IconPlayerPlay size={16} />}
            variant="light"
            onClick={handlePreview}
            loading={previewMutation.isPending}
          >
            Preview
          </Button>

          <Button
            leftSection={<IconDeviceFloppy size={16} />}
            color="green"
            onClick={handleSave}
            loading={saveMutation.isPending}
          >
            Save as Trace
          </Button>

          <Tooltip label="Reset to default">
            <ActionIcon
              variant="subtle"
              onClick={() => {
                setTemplate(DEFAULT_TEMPLATE);
                setPreview(null);
              }}
            >
              <IconReload size={18} />
            </ActionIcon>
          </Tooltip>
        </Group>
      </Group>

      {/* Template Editor */}
      <Stack gap="xs">
        <Group justify="space-between">
          <Text size="sm" fw={500}>
            Template
          </Text>
          <Group gap="xs">
            <Tooltip label="Insert faker function">
              <Button size="xs" variant="subtle" leftSection={<IconSparkles size={14} />}>
                Insert Faker
              </Button>
            </Tooltip>
            <Tooltip label="Examples menu" >
              <Menu>
                <Menu.Target>
                  <Button size="xs" variant="subtle" leftSection={<IconCodePlus size={14} />} rightSection={<IconChevronDown size={12} />}>
                    Examples
                  </Button>
                </Menu.Target>
                <Menu.Dropdown>
                  {Object.keys(examples).map((name) => (
                    <Menu.Item key={name} onClick={() => handleLoadExample(name)}>
                      {name}
                    </Menu.Item>
                  ))}
                </Menu.Dropdown>
              </Menu>
            </Tooltip>
          </Group>
        </Group>

        <Textarea
          value={template}
          onChange={(e) => setTemplate(e.target.value)}
          minRows={15}
          maxRows={15}
          styles={{
            input: {
              fontFamily: 'monospace',
              fontSize: '13px',
              backgroundColor: 'var(--mantine-color-dark-8)',
              color: 'var(--mantine-color-gray-3)',
            },
          }}
        />
      </Stack>

      {/* Configuration */}
      <Group>
        <NumberInput
          label="Lines to generate"
          value={count}
          onChange={(val) => setCount(typeof val === 'number' ? val : 100)}
          min={1}
          max={10000}
          w={150}
        />
        <TextInput
          label="Proxy Name"
          value={proxyName}
          onChange={(e) => setProxyName(e.target.value)}
          w={200}
        />
        <TextInput
          label="Filename"
          value={filename}
          onChange={(e) => setFilename(e.target.value)}
          flex={1}
        />
      </Group>

      {/* Preview Output */}
      {preview && (
        <Stack gap="xs">
          <Text size="sm" fw={500}>
            Preview Output ({preview.length} entries)
          </Text>
          <Paper withBorder p="md" style={{ backgroundColor: 'var(--mantine-color-dark-8)' }}>
            <Code block style={{ maxHeight: 400, overflow: 'auto' }}>
              {JSON.stringify(preview, null, 2)}
            </Code>
          </Paper>
        </Stack>
      )}
    </Stack>
  );
}
