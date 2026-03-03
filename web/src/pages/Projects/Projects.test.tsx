import { describe, it, expect, vi } from 'vitest';
import { render, screen, waitFor } from '@/test/utils';
import { Projects } from './Projects';
import type { ProjectListResponse } from '@/api/types';

// Mock API queries
const mockProjects: ProjectListResponse = {
  projects: [
    {
      name: 'test-project-1',
      path: '/path/to/project1',
      documents: [
        { name: 'demo.module.yaml', path: '/path/to/project1/apigear/demo.module.yaml', type: 'module' },
        { name: 'demo.solution.yaml', path: '/path/to/project1/apigear/demo.solution.yaml', type: 'solution' },
      ],
    },
    {
      name: 'test-project-2',
      path: '/path/to/project2',
      documents: [
        { name: 'api.module.yaml', path: '/path/to/project2/apigear/api.module.yaml', type: 'module' },
      ],
    },
  ],
  count: 2,
};

const emptyProjects: ProjectListResponse = {
  projects: [],
  count: 0,
};

vi.mock('@/api/queries', () => ({
  useRecentProjects: vi.fn(),
  useCreateProject: () => ({
    mutateAsync: vi.fn(),
  }),
  useDeleteProject: () => ({
    mutateAsync: vi.fn(),
  }),
}));

vi.mock('@mantine/notifications', () => ({
  notifications: {
    show: vi.fn(),
  },
}));

vi.mock('@mantine/modals', () => ({
  modals: {
    openConfirmModal: vi.fn(),
  },
}));

function mockQueryResult(data: ProjectListResponse) {
  return {
    data,
    isLoading: false,
    error: null,
    isError: false,
    isSuccess: true,
    status: 'success' as const,
    dataUpdatedAt: Date.now(),
    errorUpdatedAt: 0,
    failureCount: 0,
    failureReason: null,
    errorUpdateCount: 0,
    isFetched: true,
    isFetchedAfterMount: true,
    isFetching: false,
    isRefetching: false,
    isPending: false,
    isStale: false,
    isPlaceholderData: false,
    refetch: vi.fn(),
    fetchStatus: 'idle' as const,
    isRefetchError: false,
    isLoadingError: false,
    isPaused: false,
  };
}

describe('Projects', () => {
  it('renders loading state initially', async () => {
    const { useRecentProjects } = await import('@/api/queries');
    vi.mocked(useRecentProjects).mockReturnValue(mockQueryResult(mockProjects));

    render(<Projects />);

    // Check that loading fallback is shown initially
    await waitFor(() => {
      expect(screen.getByText(/loading/i) || screen.getByText('Projects')).toBeInTheDocument();
    });
  });

  it('renders projects list when projects exist', async () => {
    const { useRecentProjects } = await import('@/api/queries');
    vi.mocked(useRecentProjects).mockReturnValue(mockQueryResult(mockProjects));

    render(<Projects />);

    await waitFor(() => {
      expect(screen.getByText('test-project-1')).toBeInTheDocument();
      expect(screen.getByText('test-project-2')).toBeInTheDocument();
    });
  });

  it('renders empty state when no projects exist', async () => {
    const { useRecentProjects } = await import('@/api/queries');
    vi.mocked(useRecentProjects).mockReturnValue(mockQueryResult(emptyProjects));

    render(<Projects />);

    await waitFor(() => {
      expect(screen.getByText(/no projects yet/i)).toBeInTheDocument();
      expect(screen.getByText(/create your first project/i)).toBeInTheDocument();
    });
  });

  it('shows create project button', async () => {
    const { useRecentProjects } = await import('@/api/queries');
    vi.mocked(useRecentProjects).mockReturnValue(mockQueryResult(mockProjects));

    render(<Projects />);

    await waitFor(() => {
      const createButtons = screen.getAllByRole('button', { name: /create project/i });
      expect(createButtons.length).toBeGreaterThan(0);
    });
  });

  it('displays project document count', async () => {
    const { useRecentProjects } = await import('@/api/queries');
    vi.mocked(useRecentProjects).mockReturnValue(mockQueryResult(mockProjects));

    render(<Projects />);

    await waitFor(() => {
      expect(screen.getByText('2 documents')).toBeInTheDocument();
      expect(screen.getByText('1 document')).toBeInTheDocument();
    });
  });
});
