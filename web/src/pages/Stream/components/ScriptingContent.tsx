import { useState } from 'react';
import {
  Stack,
  Title,
  Grid,
  Group,
  Button,
  TextInput,
  Modal,
  Text,
  Paper,
  Menu,
  ActionIcon,
} from '@mantine/core';
import {
  IconPlayerPlay,
  IconPlayerStop,
  IconDeviceFloppy,
  IconFilePlus,
  IconBulb,
  IconDots,
} from '@tabler/icons-react';
import { notifications } from '@mantine/notifications';
import {
  useScripts,
  useRunningScripts,
  useSaveScript,
  useDeleteScript,
  useRunScript,
  useRunCode,
  useStopScript,
} from '@/api/queries';
import { ScriptEditor } from './ScriptEditor';
import { ConsoleOutput } from './ConsoleOutput';
import { ScriptList, RunningScripts } from './ScriptList';
import { EXAMPLES } from './examples';

export function ScriptingContent() {
  const { data: scripts } = useScripts();
  const { data: runningScripts } = useRunningScripts();

  const [code, setCode] = useState('');
  const [currentScript, setCurrentScript] = useState<string | null>(null);
  const [currentModTime, setCurrentModTime] = useState<number>(0);
  const [activeScriptId, setActiveScriptId] = useState<string | null>(null);

  const [saveModalOpen, setSaveModalOpen] = useState(false);
  const [scriptName, setScriptName] = useState('');
  const [examplesModalOpen, setExamplesModalOpen] = useState(false);

  const saveScript = useSaveScript();
  const deleteScript = useDeleteScript();
  const runScript = useRunScript();
  const runCode = useRunCode();
  const stopScript = useStopScript();

  const handleNew = () => {
    setCode('');
    setCurrentScript(null);
    setCurrentModTime(0);
    setActiveScriptId(null);
  };

  const handleLoadScript = async (name: string) => {
    try {
      const response = await fetch(`/api/v1/stream/scripts/${encodeURIComponent(name)}`);
      if (!response.ok) {
        throw new Error('Failed to load script');
      }
      const script = await response.json();
      setCode(script.code);
      setCurrentScript(name);
      setCurrentModTime(script.modTime);
      notifications.show({
        title: 'Script Loaded',
        message: `Loaded script: ${name}`,
        color: 'blue',
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: 'Failed to load script',
        color: 'red',
      });
    }
  };

  const handleSave = () => {
    if (currentScript) {
      // Save existing script
      saveScript.mutate(
        {
          name: currentScript,
          code,
          expectedModTime: currentModTime,
        },
        {
          onSuccess: (data) => {
            setCurrentModTime(data.modTime);
            notifications.show({
              title: 'Success',
              message: 'Script saved successfully',
              color: 'green',
            });
          },
          onError: (error: Error) => {
            if (error.message.includes('409')) {
              notifications.show({
                title: 'Conflict',
                message: 'Script was modified by another user. Please reload and try again.',
                color: 'red',
              });
            } else {
              notifications.show({
                title: 'Error',
                message: 'Failed to save script',
                color: 'red',
              });
            }
          },
        }
      );
    } else {
      // Open save dialog for new script
      setScriptName('');
      setSaveModalOpen(true);
    }
  };

  const handleSaveAs = () => {
    setScriptName('');
    setSaveModalOpen(true);
  };

  const handleSaveConfirm = () => {
    if (!scriptName) {
      notifications.show({
        title: 'Error',
        message: 'Script name is required',
        color: 'red',
      });
      return;
    }

    saveScript.mutate(
      {
        name: scriptName,
        code,
      },
      {
        onSuccess: (data) => {
          setCurrentScript(scriptName);
          setCurrentModTime(data.modTime);
          setSaveModalOpen(false);
          notifications.show({
            title: 'Success',
            message: 'Script saved successfully',
            color: 'green',
          });
        },
        onError: () => {
          notifications.show({
            title: 'Error',
            message: 'Failed to save script',
            color: 'red',
          });
        },
      }
    );
  };

  const handleDelete = (name: string) => {
    if (confirm(`Are you sure you want to delete "${name}"?`)) {
      deleteScript.mutate(name, {
        onSuccess: () => {
          if (currentScript === name) {
            handleNew();
          }
          notifications.show({
            title: 'Success',
            message: 'Script deleted successfully',
            color: 'green',
          });
        },
        onError: () => {
          notifications.show({
            title: 'Error',
            message: 'Failed to delete script',
            color: 'red',
          });
        },
      });
    }
  };

  const handleRun = () => {
    if (currentScript) {
      // Run saved script
      runScript.mutate(currentScript, {
        onSuccess: (data) => {
          setActiveScriptId(data.id);
          notifications.show({
            title: 'Success',
            message: 'Script started',
            color: 'green',
          });
        },
        onError: () => {
          notifications.show({
            title: 'Error',
            message: 'Failed to start script',
            color: 'red',
          });
        },
      });
    } else {
      // Run ad-hoc code
      runCode.mutate(
        { code, name: 'ad-hoc' },
        {
          onSuccess: (data) => {
            setActiveScriptId(data.id);
            notifications.show({
              title: 'Success',
              message: 'Code started',
              color: 'green',
            });
          },
          onError: () => {
            notifications.show({
              title: 'Error',
              message: 'Failed to run code',
              color: 'red',
            });
          },
        }
      );
    }
  };

  const handleStop = (id: string) => {
    stopScript.mutate(id, {
      onSuccess: () => {
        if (activeScriptId === id) {
          setActiveScriptId(null);
        }
        notifications.show({
          title: 'Success',
          message: 'Script stopped',
          color: 'blue',
        });
      },
      onError: () => {
        notifications.show({
          title: 'Error',
          message: 'Failed to stop script',
          color: 'red',
        });
      },
    });
  };

  const handleLoadExample = (exampleCode: string) => {
    setCode(exampleCode);
    setCurrentScript(null);
    setCurrentModTime(0);
    setExamplesModalOpen(false);
    notifications.show({
      title: 'Example Loaded',
      message: 'Example code loaded into editor',
      color: 'blue',
    });
  };

  return (
    <Stack gap="md">
      <Group justify="space-between">
        <Title order={2}>Scripting</Title>
        <Group gap="xs">
          <Button
            leftSection={<IconFilePlus size={16} />}
            variant="default"
            size="sm"
            onClick={handleNew}
          >
            New
          </Button>
          <Button
            leftSection={<IconBulb size={16} />}
            variant="default"
            size="sm"
            onClick={() => setExamplesModalOpen(true)}
          >
            Examples
          </Button>
          <Menu position="bottom-end">
            <Menu.Target>
              <ActionIcon variant="default" size="lg">
                <IconDots size={16} />
              </ActionIcon>
            </Menu.Target>
            <Menu.Dropdown>
              <Menu.Item onClick={handleSave} disabled={!code}>
                Save{currentScript && ' (Ctrl+S)'}
              </Menu.Item>
              <Menu.Item onClick={handleSaveAs} disabled={!code}>
                Save As...
              </Menu.Item>
            </Menu.Dropdown>
          </Menu>
          <Button
            leftSection={<IconDeviceFloppy size={16} />}
            variant="default"
            size="sm"
            onClick={handleSave}
            disabled={!code}
          >
            Save
          </Button>
          <Button
            leftSection={<IconPlayerPlay size={16} />}
            color="green"
            size="sm"
            onClick={handleRun}
            disabled={!code}
          >
            Run
          </Button>
          {activeScriptId && (
            <Button
              leftSection={<IconPlayerStop size={16} />}
              color="red"
              size="sm"
              onClick={() => handleStop(activeScriptId)}
            >
              Stop
            </Button>
          )}
        </Group>
      </Group>

      {currentScript && (
        <Paper p="xs" withBorder>
          <Text size="sm">
            Editing: <strong>{currentScript}</strong>
          </Text>
        </Paper>
      )}

      <Grid>
        <Grid.Col span={3}>
          <ScriptList
            scripts={scripts}
            currentScript={currentScript}
            onSelect={handleLoadScript}
            onDelete={handleDelete}
            onRun={(name) => {
              handleLoadScript(name);
              setTimeout(() => handleRun(), 100);
            }}
          />
          <RunningScripts scripts={runningScripts} onStop={handleStop} />
        </Grid.Col>

        <Grid.Col span={9}>
          <Stack gap="md">
            <Paper withBorder>
              <ScriptEditor code={code} onChange={setCode} height="500px" />
            </Paper>

            {activeScriptId && <ConsoleOutput scriptId={activeScriptId} height="300px" />}
          </Stack>
        </Grid.Col>
      </Grid>

      {/* Save Modal */}
      <Modal
        opened={saveModalOpen}
        onClose={() => setSaveModalOpen(false)}
        title="Save Script"
      >
        <Stack gap="md">
          <TextInput
            label="Script Name"
            placeholder="my-script"
            value={scriptName}
            onChange={(e) => setScriptName(e.currentTarget.value)}
            data-autofocus
          />
          <Group justify="flex-end">
            <Button variant="default" onClick={() => setSaveModalOpen(false)}>
              Cancel
            </Button>
            <Button onClick={handleSaveConfirm} disabled={!scriptName}>
              Save
            </Button>
          </Group>
        </Stack>
      </Modal>

      {/* Examples Modal */}
      <Modal
        opened={examplesModalOpen}
        onClose={() => setExamplesModalOpen(false)}
        title="Script Examples"
        size="lg"
      >
        <Stack gap="md">
          {EXAMPLES.map((example) => (
            <Paper key={example.name} withBorder p="md" style={{ cursor: 'pointer' }} onClick={() => handleLoadExample(example.code)}>
              <Stack gap="xs">
                <Group justify="space-between">
                  <Text fw={600}>{example.name}</Text>
                  <Button size="xs" variant="light" onClick={(e) => {
                    e.stopPropagation();
                    handleLoadExample(example.code);
                  }}>
                    Load
                  </Button>
                </Group>
                <Text size="sm" c="dimmed">
                  {example.description}
                </Text>
              </Stack>
            </Paper>
          ))}
        </Stack>
      </Modal>
    </Stack>
  );
}
