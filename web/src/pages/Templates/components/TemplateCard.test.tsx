import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@/test/utils';
import userEvent from '@testing-library/user-event';
import { TemplateCard } from './TemplateCard';
import type { TemplateInfo } from '@/api/types';

// Mock the entire queries module
vi.mock('@/api/queries', () => ({
  useInstallTemplate: () => ({
    mutateAsync: vi.fn(),
    isPending: false,
  }),
}));

// Mock notifications
vi.mock('@mantine/notifications', () => ({
  notifications: {
    show: vi.fn(),
  },
}));

beforeEach(() => {
  vi.clearAllMocks();
});

describe('TemplateCard', () => {
  const mockTemplate: TemplateInfo = {
    name: 'test-template',
    description: 'A test template',
    latest: '1.0.0',
    version: '',
    git: 'https://github.com/test/template.git',
    inCache: false,
    updateNeeded: false,
    versions: ['1.0.0', '0.9.0', '0.8.0'],
  };

  it('renders template information correctly', () => {
    render(<TemplateCard template={mockTemplate} />);

    expect(screen.getByText('test-template')).toBeInTheDocument();
    expect(screen.getByText('A test template')).toBeInTheDocument();
    expect(screen.getByText('Latest: 1.0.0')).toBeInTheDocument();
  });

  it('shows Install button for non-cached templates', () => {
    render(<TemplateCard template={mockTemplate} />);

    const installButton = screen.getByRole('button', { name: /install/i });
    expect(installButton).toBeInTheDocument();
    expect(installButton).not.toBeDisabled();
  });

  it('shows Installed badge and Update button for cached templates with updates', () => {
    const cachedTemplate: TemplateInfo = {
      ...mockTemplate,
      inCache: true,
      version: '0.9.0',
      updateNeeded: true,
    };

    render(<TemplateCard template={cachedTemplate} />);

    expect(screen.getByText('Installed')).toBeInTheDocument();
    expect(screen.getByText('Update Available')).toBeInTheDocument();
    expect(screen.getByRole('button', { name: /update/i })).toBeInTheDocument();
    expect(screen.getByText(/Installed: 0\.9\.0/)).toBeInTheDocument();
  });

  it('shows Up to Date button (disabled) for up-to-date templates', () => {
    const upToDateTemplate: TemplateInfo = {
      ...mockTemplate,
      inCache: true,
      version: '1.0.0',
      updateNeeded: false,
    };

    render(<TemplateCard template={upToDateTemplate} />);

    const upToDateButton = screen.getByRole('button', { name: /up to date/i });
    expect(upToDateButton).toBeInTheDocument();
    expect(upToDateButton).toBeDisabled();
  });

  it('shows version selector dropdown when multiple versions are available', async () => {
    const user = userEvent.setup();
    render(<TemplateCard template={mockTemplate} />);

    // Find the dropdown button (chevron icon button)
    const dropdownButtons = screen.getAllByRole('button');
    const versionDropdown = dropdownButtons.find(
      (btn) => !btn.textContent?.includes('Install')
    );

    expect(versionDropdown).toBeInTheDocument();
  });

  it('renders GitHub link when git URL is provided', () => {
    render(<TemplateCard template={mockTemplate} />);

    const githubLink = screen.getByRole('link');
    expect(githubLink).toHaveAttribute('href', 'https://github.com/test/template');
    expect(githubLink).toHaveAttribute('target', '_blank');
  });

  it('does not render GitHub link when git URL is not provided', () => {
    const templateWithoutGit = { ...mockTemplate, git: '' };
    render(<TemplateCard template={templateWithoutGit} />);

    const links = screen.queryAllByRole('link');
    expect(links).toHaveLength(0);
  });

  it('shows fallback text when description is not provided', () => {
    const templateWithoutDescription = { ...mockTemplate, description: '' };
    render(<TemplateCard template={templateWithoutDescription} />);

    expect(screen.getByText('No description available')).toBeInTheDocument();
  });
});
