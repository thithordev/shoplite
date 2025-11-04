module.exports = {
  root: true,
  parser: '@typescript-eslint/parser',
  plugins: ['@typescript-eslint', 'react-hooks'],
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:react-hooks/recommended',
    'prettier'
  ],
  env: { browser: true, es2021: true, node: true },
  ignorePatterns: ['dist', 'node_modules'],
  settings: {
    react: { version: 'detect' }
  },
  rules: {
    '@typescript-eslint/no-explicit-any': 'off'
  }
}
