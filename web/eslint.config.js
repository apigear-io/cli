// ESLint Flat Config (v9+)
//
// LINTING ONLY - NO FORMATTING
// This config focuses exclusively on code quality and correctness.
// Formatting is handled by Prettier or not at all.
//
// Configs used:
// - js.configs.recommended: Basic JavaScript linting (no style rules)
// - tseslint.configs.recommended: TypeScript linting (no style rules)
// - react-hooks: React Hooks rules of hooks
// - react-refresh: Fast refresh compatibility
//
// We explicitly do NOT use:
// - @stylistic/* plugins
// - Any formatting-related rules
// - tseslint.configs.stylistic

import js from '@eslint/js';
import globals from 'globals';
import reactHooks from 'eslint-plugin-react-hooks';
import reactRefresh from 'eslint-plugin-react-refresh';
import tseslint from 'typescript-eslint';

export default tseslint.config(
  { ignores: ['dist', 'node_modules', 'coverage', 'test-results', 'playwright-report'] },
  {
    extends: [js.configs.recommended, ...tseslint.configs.recommended],
    files: ['**/*.{ts,tsx}'],
    languageOptions: {
      ecmaVersion: 2020,
      globals: globals.browser,
    },
    plugins: {
      'react-hooks': reactHooks,
      'react-refresh': reactRefresh,
    },
    rules: {
      ...reactHooks.configs.recommended.rules,
      'react-refresh/only-export-components': [
        'warn',
        { allowConstantExport: true },
      ],
      '@typescript-eslint/no-unused-vars': [
        'error',
        {
          argsIgnorePattern: '^_',
          varsIgnorePattern: '^_',
        },
      ],
    },
  },
  {
    files: ['**/*.test.{ts,tsx}', '**/test/**/*.{ts,tsx}'],
    rules: {
      'react-refresh/only-export-components': 'off',
    },
  },
  {
    files: ['**/vite.config.ts', '**/vitest.config.ts', '**/playwright.config.ts'],
    rules: {
      'react-refresh/only-export-components': 'off',
    },
  }
);
