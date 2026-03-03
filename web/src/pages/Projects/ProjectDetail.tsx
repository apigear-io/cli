import { Suspense, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import {
  Stack,
  Text,
  Paper,
  Group,
  Badge,
  Button,
  ActionIcon,
  Tooltip,
} from '@mantine/core';
import {
  IconFolder,
  IconFile,
  IconEdit,
  IconExternalLink,
  IconPlayerPlay,
} from '@tabler/icons-react';
import { notifications } from '@mantine/notifications';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';
import { Breadcrumbs } from '@/components/Breadcrumbs';
import { FileEditorModal } from '@/components/FileEditorModal';
import { DocumentInfoDrawer } from './components';
import { useProject, useOpenFileExternal } from '@/api/queries';
import type { DocumentInfo } from '@/api/types';

const getDocumentTypeColor = (type: string | undefined) => {
  if (!type) {
    return 'gray';
  }

  switch (type.toLowerCase()) {
    case 'module':
      return 'blue';
    case 'solution':
      return 'green';
    case 'simulation':
      return 'orange';
    case 'scenario':
      return 'purple';
    case 'unknown':
      return 'gray';
    default:
      return 'gray';
  }
};

const isSolutionFile = (doc: DocumentInfo) => {
  return doc.name.endsWith('.solution.yaml');
};

function ProjectDetailContent() {
  const { encodedPath } = useParams<{ encodedPath: string }>();
  const navigate = useNavigate();
  const openExternal = useOpenFileExternal();

  const [editingFile, setEditingFile] = useState<{ path: string; name: string } | null>(null);
  const [openingExternalPath, setOpeningExternalPath] = useState<string | null>(null);
  const [selectedDocument, setSelectedDocument] = useState<DocumentInfo | null>(null);

  if (!encodedPath) {
    navigate('/codegen/projects');
    return null;
  }

  const projectPath = decodeURIComponent(encodedPath);
  const { data: project } = useProject(projectPath);

  const handleEdit = (doc: DocumentInfo) => {
    setEditingFile({ path: doc.path, name: doc.name });
    setSelectedDocument(null); // Close drawer when opening editor
  };

  const handleOpenExternal = async (doc: DocumentInfo) => {
    setOpeningExternalPath(doc.path);
    try {
      await openExternal.mutateAsync({ path: doc.path });
      notifications.show({
        title: 'Success',
        message: 'File opened in external editor',
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Error',
        message: error instanceof Error ? error.message : 'Failed to open file',
        color: 'red',
      });
    } finally {
      setOpeningExternalPath(null);
    }
  };

  const handleGenerate = (doc: DocumentInfo) => {
    // Navigate to code generation page with encoded solution path
    const encodedSolutionPath = encodeURIComponent(doc.path);
    navigate(`/codegen/projects/generate/${encodedSolutionPath}`);
  };

  const handleDocumentClick = (doc: DocumentInfo, e: React.MouseEvent) => {
    // Don't open drawer if clicking on action buttons
    const target = e.target as HTMLElement;
    if (target.closest('button') || target.closest('[role="button"]')) {
      return;
    }
    setSelectedDocument(doc);
  };

  return (
    <Stack gap="lg">
      <Group justify="space-between" align="flex-start">
        <div>
          <Breadcrumbs
            items={[
              { label: 'Projects', href: '/codegen/projects' },
              { label: project.name },
            ]}
          />
          <Group gap="xs" mt="xs">
            <IconFolder size={20} style={{ opacity: 0.6 }} />
            <Text size="sm" c="dimmed">
              {project.path}
            </Text>
          </Group>
        </div>
      </Group>

      {/* Documents Section */}
      {project.documents.length === 0 ? (
        <Paper shadow="sm" p="lg" withBorder>
          <Text c="dimmed" ta="center" py="xl">
            No documents found in this project
          </Text>
        </Paper>
      ) : (
        <Stack gap="sm">
          {project.documents.map((doc, index) => (
            <Paper
              key={index}
              shadow="xs"
              p="md"
              withBorder
              style={{ cursor: 'pointer' }}
              onClick={(e) => handleDocumentClick(doc, e)}
            >
              <Group justify="space-between" wrap="nowrap">
                <Group gap="md" style={{ flex: 1, minWidth: 0 }}>
                  <IconFile size={20} style={{ flexShrink: 0 }} />
                  <div style={{ flex: 1, minWidth: 0 }}>
                    <Group gap="xs" mb={4}>
                      <Text fw={500}>{doc.name}</Text>
                      <Badge color={getDocumentTypeColor(doc.type)} size="sm">
                        {doc.type || 'unknown'}
                      </Badge>
                    </Group>
                    <Text
                      size="xs"
                      c="dimmed"
                      style={{ wordBreak: 'break-all', overflow: 'hidden' }}
                    >
                      {doc.path}
                    </Text>
                  </div>
                </Group>

                <Group gap="xs" style={{ flexShrink: 0 }}>
                  {isSolutionFile(doc) && (
                    <Tooltip label="Generate code">
                      <Button
                        size="sm"
                        variant="light"
                        color="green"
                        leftSection={<IconPlayerPlay size={16} />}
                        onClick={() => handleGenerate(doc)}
                      >
                        Generate
                      </Button>
                    </Tooltip>
                  )}

                  <Tooltip label="Edit in browser">
                    <ActionIcon
                      size="lg"
                      variant="light"
                      onClick={() => handleEdit(doc)}
                    >
                      <IconEdit size={18} />
                    </ActionIcon>
                  </Tooltip>

                  <Tooltip label="Open in external editor">
                    <ActionIcon
                      size="lg"
                      variant="light"
                      onClick={() => handleOpenExternal(doc)}
                      loading={openingExternalPath === doc.path}
                    >
                      <IconExternalLink size={18} />
                    </ActionIcon>
                  </Tooltip>
                </Group>
              </Group>
            </Paper>
          ))}
        </Stack>
      )}

      <DocumentInfoDrawer
        document={selectedDocument}
        onClose={() => setSelectedDocument(null)}
        onEdit={handleEdit}
        onOpenExternal={handleOpenExternal}
        onGenerate={handleGenerate}
        isOpeningExternal={openingExternalPath === selectedDocument?.path}
      />

      <FileEditorModal
        opened={!!editingFile}
        onClose={() => setEditingFile(null)}
        filePath={editingFile?.path || null}
        fileName={editingFile?.name}
      />
    </Stack>
  );
}

export function ProjectDetail() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingFallback message="Loading project..." />}>
        <ProjectDetailContent />
      </Suspense>
    </ErrorBoundary>
  );
}
