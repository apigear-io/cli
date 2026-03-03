import { Suspense, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import {
  Stack,
  Text,
  Paper,
  Group,
} from '@mantine/core';
import {
  IconFolder,
  IconApi,
  IconSettings,
  IconFileDescription,
} from '@tabler/icons-react';
import { notifications } from '@mantine/notifications';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { LoadingFallback } from '@/components/LoadingFallback';
import { Breadcrumbs } from '@/components/Breadcrumbs';
import { FileEditorModal } from '@/components/FileEditorModal';
import { DocumentInfoDrawer, DocumentSection } from './components';
import { useProject, useOpenFileExternal } from '@/api/queries';
import type { DocumentInfo } from '@/api/types';

function groupDocuments(docs: DocumentInfo[]) {
  const modules: DocumentInfo[] = [];
  const solutions: DocumentInfo[] = [];
  const others: DocumentInfo[] = [];

  for (const doc of docs) {
    switch (doc.type?.toLowerCase()) {
      case 'module':
        modules.push(doc);
        break;
      case 'solution':
        solutions.push(doc);
        break;
      default:
        others.push(doc);
        break;
    }
  }

  return { modules, solutions, others };
}

function ProjectDetailContent() {
  const { encodedPath } = useParams<{ encodedPath: string }>();
  const navigate = useNavigate();
  const openExternal = useOpenFileExternal();

  const projectPath = encodedPath ? decodeURIComponent(encodedPath) : '';
  const { data: project } = useProject(projectPath);

  const [editingFile, setEditingFile] = useState<{ path: string; name: string } | null>(null);
  const [openingExternalPath, setOpeningExternalPath] = useState<string | null>(null);
  const [selectedDocument, setSelectedDocument] = useState<DocumentInfo | null>(null);

  if (!encodedPath) {
    navigate('/codegen/projects');
    return null;
  }

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
        <Stack gap="lg">
          {(() => {
            const { modules, solutions, others } = groupDocuments(project.documents);
            return (
              <>
                <DocumentSection
                  title="API Modules"
                  icon={<IconApi size={20} />}
                  documents={modules}
                  onEdit={handleEdit}
                  onOpenExternal={handleOpenExternal}
                  onDocumentClick={handleDocumentClick}
                  openingExternalPath={openingExternalPath}
                />
                <DocumentSection
                  title="Solutions"
                  icon={<IconSettings size={20} />}
                  documents={solutions}
                  onEdit={handleEdit}
                  onOpenExternal={handleOpenExternal}
                  onDocumentClick={handleDocumentClick}
                  onGenerate={handleGenerate}
                  showGenerateButton
                  openingExternalPath={openingExternalPath}
                />
                <DocumentSection
                  title="Other Files"
                  icon={<IconFileDescription size={20} />}
                  documents={others}
                  onEdit={handleEdit}
                  onOpenExternal={handleOpenExternal}
                  onDocumentClick={handleDocumentClick}
                  openingExternalPath={openingExternalPath}
                />
              </>
            );
          })()}
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
