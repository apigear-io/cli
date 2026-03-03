import { useState } from 'react';
import {
  Modal,
  Stack,
  Text,
  Button,
  Group,
  ScrollArea,
  ActionIcon,
  TextInput,
  Loader,
  Alert,
} from '@mantine/core';
import {
  IconFolder,
  IconFolderOpen,
  IconHome,
  IconChevronRight,
  IconLock,
  IconArrowLeft,
} from '@tabler/icons-react';
import { useQuery } from '@tanstack/react-query';
import { apiClient } from '@/api/client';
import type { DirectoryListResponse } from '@/api/types';

interface DirectoryBrowserProps {
  opened: boolean;
  onClose: () => void;
  onSelect: (path: string) => void;
  initialPath?: string;
}

export function DirectoryBrowser({
  opened,
  onClose,
  onSelect,
  initialPath,
}: DirectoryBrowserProps) {
  const [currentPath, setCurrentPath] = useState(initialPath || '');

  const { data, isLoading, error } = useQuery({
    queryKey: ['browse-directories', currentPath],
    queryFn: () =>
      apiClient.get<DirectoryListResponse>(
        `/projects/browse?path=${encodeURIComponent(currentPath)}`
      ),
    enabled: opened,
  });

  const handleNavigate = (path: string) => {
    setCurrentPath(path);
  };

  const handleSelect = () => {
    if (currentPath) {
      onSelect(currentPath);
      onClose();
    }
  };

  const handleGoHome = async () => {
    // Navigate to home by not providing a path
    setCurrentPath('');
  };

  return (
    <Modal
      opened={opened}
      onClose={onClose}
      title="Select Directory"
      size="lg"
      styles={{
        body: { height: '500px', display: 'flex', flexDirection: 'column' },
      }}
    >
      <Stack gap="md" style={{ flex: 1, minHeight: 0 }}>
        {/* Current Path Display */}
        <Group gap="xs">
          <ActionIcon variant="light" onClick={handleGoHome} title="Go to home directory">
            <IconHome size={16} />
          </ActionIcon>
          <TextInput
            value={data?.currentPath || currentPath || 'Loading...'}
            readOnly
            style={{ flex: 1 }}
            leftSection={<IconFolderOpen size={16} />}
          />
        </Group>

        {/* Error State */}
        {error && (
          <Alert color="red" title="Error">
            {error instanceof Error ? error.message : 'Failed to load directory'}
          </Alert>
        )}

        {/* Loading State */}
        {isLoading && (
          <Group justify="center" p="xl">
            <Loader size="md" />
            <Text>Loading directory...</Text>
          </Group>
        )}

        {/* Directory List */}
        {data && !isLoading && (
          <ScrollArea style={{ flex: 1 }} offsetScrollbars>
            <Stack gap="xs">
              {/* Parent Directory */}
              {data.parentPath && (
                <Button
                  variant="light"
                  leftSection={<IconArrowLeft size={16} />}
                  onClick={() => handleNavigate(data.parentPath)}
                  fullWidth
                  justify="flex-start"
                >
                  .. (Parent Directory)
                </Button>
              )}

              {/* Subdirectories */}
              {data.directories.length === 0 && (
                <Text c="dimmed" ta="center" p="xl">
                  No subdirectories found
                </Text>
              )}

              {data.directories.map((dir) => (
                <Button
                  key={dir.path}
                  variant="subtle"
                  leftSection={
                    dir.accessible ? <IconFolder size={16} /> : <IconLock size={16} />
                  }
                  rightSection={dir.accessible ? <IconChevronRight size={16} /> : null}
                  onClick={() => dir.accessible && handleNavigate(dir.path)}
                  disabled={!dir.accessible}
                  fullWidth
                  justify="flex-start"
                  styles={{
                    label: {
                      overflow: 'hidden',
                      textOverflow: 'ellipsis',
                      whiteSpace: 'nowrap',
                    },
                  }}
                >
                  {dir.name}
                </Button>
              ))}
            </Stack>
          </ScrollArea>
        )}

        {/* Action Buttons */}
        <Group justify="space-between">
          <Button variant="default" onClick={onClose}>
            Cancel
          </Button>
          <Button onClick={handleSelect} disabled={!currentPath}>
            Select This Directory
          </Button>
        </Group>
      </Stack>
    </Modal>
  );
}
