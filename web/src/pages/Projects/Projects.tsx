import { Suspense, useState } from 'react';
import { Stack, Button, Group, SimpleGrid } from '@mantine/core';
import { IconPlus, IconFolderOpen } from '@tabler/icons-react';
import { notifications } from '@mantine/notifications';
import { modals } from '@mantine/modals';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';
import { Breadcrumbs } from '@/components/Breadcrumbs';
import {
  useRecentProjects,
  useCreateProject,
  useDeleteProject,
} from '@/api/queries';
import {
  ProjectCard,
  CreateProjectModal,
  OpenProjectModal,
  EmptyState,
} from './components';

function ProjectsContent() {
  const { data } = useRecentProjects();
  const createProject = useCreateProject();
  const deleteProject = useDeleteProject();

  const [createModalOpen, setCreateModalOpen] = useState(false);
  const [openModalOpen, setOpenModalOpen] = useState(false);

  const handleCreate = async (req: { name: string; path: string }) => {
    try {
      await createProject.mutateAsync(req);
      notifications.show({
        title: 'Success',
        message: 'Project created successfully',
        color: 'green',
      });
      setCreateModalOpen(false);
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to create project',
        color: 'red',
      });
      throw error; // Re-throw to prevent modal from closing
    }
  };

  const handleDelete = (path: string, name: string) => {
    modals.openConfirmModal({
      title: 'Delete Project',
      children: `Are you sure you want to delete "${name}"? This will permanently remove the project directory from your disk. This action cannot be undone.`,
      labels: { confirm: 'Delete', cancel: 'Cancel' },
      confirmProps: { color: 'red' },
      onConfirm: async () => {
        try {
          await deleteProject.mutateAsync(path);
          notifications.show({
            title: 'Success',
            message: 'Project deleted successfully',
            color: 'green',
          });
        } catch (error) {
          notifications.show({
            title: 'Error',
            message: error instanceof Error ? error.message : 'Failed to delete project',
            color: 'red',
          });
        }
      },
    });
  };

  return (
    <Stack gap="lg">
      <Group justify="space-between" align="flex-start">
        <Breadcrumbs items={[{ label: 'Projects' }]} />
        <Group>
          <Button
            variant="light"
            onClick={() => setOpenModalOpen(true)}
            leftSection={<IconFolderOpen size={16} />}
          >
            Open Project
          </Button>
          <Button
            onClick={() => setCreateModalOpen(true)}
            leftSection={<IconPlus size={16} />}
          >
            Create Project
          </Button>
        </Group>
      </Group>

      {data.count === 0 ? (
        <EmptyState
          onCreate={() => setCreateModalOpen(true)}
          onOpen={() => setOpenModalOpen(true)}
        />
      ) : (
        <SimpleGrid cols={{ base: 1, sm: 2, lg: 3 }} spacing="lg">
          {data.projects.map((project) => (
            <ProjectCard
              key={project.path}
              project={project}
              onDelete={(path) => handleDelete(path, project.name)}
            />
          ))}
        </SimpleGrid>
      )}

      <CreateProjectModal
        opened={createModalOpen}
        onClose={() => setCreateModalOpen(false)}
        onSubmit={handleCreate}
      />

      <OpenProjectModal
        opened={openModalOpen}
        onClose={() => setOpenModalOpen(false)}
      />
    </Stack>
  );
}

export function Projects() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingFallback message="Loading projects..." />}>
        <ProjectsContent />
      </Suspense>
    </ErrorBoundary>
  );
}
