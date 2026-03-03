import { useState, useEffect } from 'react';
import { Modal, Stack, Button, Group, Text, Loader } from '@mantine/core';
import { notifications } from '@mantine/notifications';
import Editor from '@monaco-editor/react';
import { useReadFile, useWriteFile } from '@/api/queries';

interface FileEditorModalProps {
  opened: boolean;
  onClose: () => void;
  filePath: string | null;
  fileName?: string;
}

export function FileEditorModal({ opened, onClose, filePath, fileName }: FileEditorModalProps) {
  const [content, setContent] = useState('');
  const [hasChanges, setHasChanges] = useState(false);

  const { data, isLoading, error } = useReadFile(filePath || '');
  const writeFile = useWriteFile();

  // Load file content when data arrives
  useEffect(() => {
    if (data) {
      setContent(data.content);
      setHasChanges(false);
    }
  }, [data]);

  const handleEditorChange = (value: string | undefined) => {
    if (value !== undefined) {
      setContent(value);
      setHasChanges(value !== data?.content);
    }
  };

  const handleSave = async () => {
    if (!filePath) return;

    try {
      await writeFile.mutateAsync({
        path: filePath,
        content,
      });

      notifications.show({
        title: 'Success',
        message: 'File saved successfully',
        color: 'green',
      });

      setHasChanges(false);
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to save file',
        color: 'red',
      });
    }
  };

  const handleClose = () => {
    if (hasChanges) {
      if (confirm('You have unsaved changes. Are you sure you want to close?')) {
        setContent('');
        setHasChanges(false);
        onClose();
      }
    } else {
      setContent('');
      setHasChanges(false);
      onClose();
    }
  };

  // Determine language based on file extension
  const getLanguage = (path: string) => {
    if (path.endsWith('.yaml') || path.endsWith('.yml')) return 'yaml';
    if (path.endsWith('.json')) return 'json';
    if (path.endsWith('.js')) return 'javascript';
    if (path.endsWith('.ts')) return 'typescript';
    if (path.endsWith('.idl')) return 'plaintext';
    return 'plaintext';
  };

  return (
    <Modal
      opened={opened}
      onClose={handleClose}
      title={fileName || 'Edit File'}
      size="xl"
      styles={{
        body: { height: '70vh', display: 'flex', flexDirection: 'column' },
        content: { height: '80vh' },
      }}
    >
      <Stack gap="md" style={{ flex: 1, minHeight: 0 }}>
        {error && (
          <Text c="red">Failed to load file: {error instanceof Error ? error.message : 'Unknown error'}</Text>
        )}

        {isLoading && (
          <Group justify="center" p="xl">
            <Loader size="md" />
            <Text>Loading file...</Text>
          </Group>
        )}

        {data && (
          <>
            <div style={{ flex: 1, border: '1px solid var(--mantine-color-gray-3)', borderRadius: '4px', overflow: 'hidden' }}>
              <Editor
                height="100%"
                language={getLanguage(filePath || '')}
                value={content}
                onChange={handleEditorChange}
                theme="vs-dark"
                options={{
                  minimap: { enabled: false },
                  fontSize: 14,
                  lineNumbers: 'on',
                  scrollBeyondLastLine: false,
                  automaticLayout: true,
                }}
              />
            </div>

            <Group justify="space-between">
              <Text size="sm" c="dimmed">
                {hasChanges && '• Unsaved changes'}
              </Text>
              <Group>
                <Button variant="default" onClick={handleClose}>
                  Cancel
                </Button>
                <Button
                  onClick={handleSave}
                  loading={writeFile.isPending}
                  disabled={!hasChanges}
                >
                  Save
                </Button>
              </Group>
            </Group>
          </>
        )}
      </Stack>
    </Modal>
  );
}
