// Typography system for the application
export const typography = {
  // Font families
  fonts: {
    primary: '-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif',
    monospace: 'ui-monospace, SFMono-Regular, "SF Mono", Monaco, Inconsolata, "Roboto Mono", monospace'
  },

  // Font sizes
  sizes: {
    xs: '12px',
    sm: '13px',
    base: '14px',
    lg: '16px',
    xl: '18px',
    '2xl': '20px',
    '3xl': '24px',
    '4xl': '28px'
  },

  // Font weights
  weights: {
    normal: '400',
    medium: '500',
    semibold: '600',
    bold: '700'
  },

  // Line heights
  lineHeights: {
    tight: '1.25',
    normal: '1.5',
    relaxed: '1.625'
  },

  // Letter spacing
  letterSpacing: {
    tight: '-0.025em',
    normal: '0',
    wide: '0.025em',
    wider: '0.05em'
  }
};

// Color system
export const colors = {
  // Text colors
  text: {
    primary: '#202124',
    secondary: '#5f6368',
    tertiary: '#80868b',
    disabled: '#dadce0'
  },

  // Background colors
  background: {
    primary: '#ffffff',
    secondary: '#f8f9fa',
    tertiary: '#f1f3f4',
    overlay: 'rgba(0, 0, 0, 0.6)'
  },

  // Brand colors
  brand: {
    primary: '#1976d2',
    secondary: '#1565c0'
  },

  // Status colors
  status: {
    success: '#0f9d58',
    warning: '#ff9800',
    error: '#dc3545',
    info: '#2196f3'
  },

  // Border colors
  border: {
    light: '#e8eaed',
    medium: '#dadce0',
    dark: '#5f6368'
  }
};

// Common text styles
export const textStyles = {
  h1: {
    fontSize: typography.sizes['4xl'],
    fontWeight: typography.weights.bold,
    lineHeight: typography.lineHeights.tight,
    color: colors.text.primary,
    fontFamily: typography.fonts.primary
  },

  h2: {
    fontSize: typography.sizes['3xl'],
    fontWeight: typography.weights.semibold,
    lineHeight: typography.lineHeights.tight,
    color: colors.text.primary,
    fontFamily: typography.fonts.primary
  },

  h3: {
    fontSize: typography.sizes.xl,
    fontWeight: typography.weights.medium,
    lineHeight: typography.lineHeights.normal,
    color: colors.text.primary,
    fontFamily: typography.fonts.primary
  },

  body: {
    fontSize: typography.sizes.base,
    fontWeight: typography.weights.normal,
    lineHeight: typography.lineHeights.normal,
    color: colors.text.primary,
    fontFamily: typography.fonts.primary
  },

  caption: {
    fontSize: typography.sizes.sm,
    fontWeight: typography.weights.normal,
    lineHeight: typography.lineHeights.normal,
    color: colors.text.secondary,
    fontFamily: typography.fonts.primary
  },

  label: {
    fontSize: typography.sizes.sm,
    fontWeight: typography.weights.medium,
    lineHeight: typography.lineHeights.normal,
    color: colors.text.secondary,
    fontFamily: typography.fonts.primary,
    textTransform: 'uppercase' as const,
    letterSpacing: typography.letterSpacing.wide
  }
};
