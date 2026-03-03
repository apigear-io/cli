import { useState, useEffect } from 'react';
import { Modal, TextInput, Button, Stack, Text, Group, ActionIcon } from '@mantine/core';
import { IconFolder, IconFolderSearch } from '@tabler/icons-react';
import { useProjectDirectories } from '@/api/queries';
import { DirectoryBrowser } from './DirectoryBrowser';
import type { CreateProjectRequest } from '@/api/types';

interface CreateProjectModalProps {
  opened: boolean;
  onClose: () => void;
  onSubmit: (request: CreateProjectRequest) => Promise<void>;
}

export function CreateProjectModal({ opened, onClose, onSubmit }: CreateProjectModalProps) {
  const { data: directories, isLoading: loadingDirectories } = useProjectDirectories();

  const [name, setName] = useState('');
  const [path, setPath] = useState('');
  const [loading, setLoading] = useState(false);
  const [errors, setErrors] = useState<{ name?: string; path?: string }>({});
  const [browserOpen, setBrowserOpen] = useState(false);

  // Set default path when directories are loaded
  useEffect(() => {
    if (directories && !path && opened) {
      // Use first suggestion (usually home directory) as default
      const defaultPath = directories.suggestions[0] || directories.homeDir || '';
      setPath(defaultPath);
    }
  }, [directories, opened, path]);

  const validate = () => {
    const newErrors: { name?: string; path?: string } = {};

    if (!name.trim()) {
      newErrors.name = 'Project name is required';
    }

    if (!path.trim()) {
      newErrors.path = 'Parent directory path is required';
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = async () => {
    if (!validate()) {
      return;
    }

    setLoading(true);
    try {
      await onSubmit({ name: name.trim(), path: path.trim() });
      // Reset form on success
      setName('');
      setPath('');
      setErrors({});
    } finally {
      setLoading(false);
    }
  };

  const handleClose = () => {
    if (!loading) {
      setName('');
      setPath('');
      setErrors({});
      onClose();
    }
  };

  const handleSelectDirectory = (selectedPath: string) => {
    setPath(selectedPath);
    setBrowserOpen(false);
  };

  return (
    <>
      <Modal opened={opened} onClose={handleClose} title="Create New Project" size="md">
        <Stack gap="md">
          <TextInput
            label="Project Name"
            placeholder="my-project"
            value={name}
            onChange={(e) => setName(e.currentTarget.value)}
            error={errors.name}
            required
            disabled={loading}
            autoFocus
          />

          <TextInput
            label="Parent Directory"
            placeholder="/path/to/parent/directory"
            value={path}
            onChange={(e) => setPath(e.currentTarget.value)}
            error={errors.path}
            required
            disabled={loading || loadingDirectories}
            leftSection={<IconFolder size={16} />}
            rightSection={
              <ActionIcon
                variant="subtle"
                onClick={() => setBrowserOpen(true)}
                disabled={loading}
                title="Browse directories"
              >
                <IconFolderSearch size={18} />
              </ActionIcon>
            }
            description="Click the browse icon to select a directory"
          />

          {name && path && (
            <Text size="sm" c="dimmed">
              <Group gap={4}>
                <Text span fw={500}>
                  Project will be created at:
                </Text>
                <Text span fs="italic">
                  {path}/{name}
                </Text>
              </Group>
            </Text>
          )}

          <Button onClick={handleSubmit} loading={loading} fullWidth>
            Create Project
          </Button>
        </Stack>
      </Modal>

      <DirectoryBrowser
        opened={browserOpen}
        onClose={() => setBrowserOpen(false)}
        onSelect={handleSelectDirectory}
        initialPath={path || directories?.homeDir}
      />
    </>
  );
}
