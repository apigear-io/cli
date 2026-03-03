import { useNavigate } from 'react-router-dom';
import { notifications } from '@mantine/notifications';
import { apiClient } from '@/api/client';
import { DirectoryBrowser } from './DirectoryBrowser';
import type { ProjectInfo } from '@/api/types';

interface OpenProjectModalProps {
  opened: boolean;
  onClose: () => void;
}

export function OpenProjectModal({ opened, onClose }: OpenProjectModalProps) {
  const navigate = useNavigate();

  const handleSelect = async (path: string) => {
    try {
      await apiClient.get<ProjectInfo>(
        `/projects/get?path=${encodeURIComponent(path)}`
      );
      const encodedPath = encodeURIComponent(path);
      onClose();
      navigate(`/codegen/projects/${encodedPath}`);
    } catch {
      notifications.show({
        title: 'Not a valid project',
        message: 'No apigear/ folder found in the selected directory. Please choose a directory that contains an ApiGear project.',
        color: 'red',
      });
    }
  };

  return (
    <DirectoryBrowser
      opened={opened}
      onClose={onClose}
      onSelect={handleSelect}
      initialPath=""
    />
  );
}
