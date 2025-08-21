import React from 'react';
import { Card } from 'primereact/card';
import { typography, colors, textStyles } from '../styles/typography';

interface ErrorDisplayProps {
    error: any;
}

const ErrorDisplay: React.FC<ErrorDisplayProps> = ({ error }) => {
    if (!error) return null;

    return (
        <Card 
            title="Error" 
            style={{
                background: colors.background.primary,
                marginBottom: '20px',
                borderLeft: `4px solid ${colors.status.error}`,
                padding: '1.5rem'
            }}
        >
            <div style={{ 
                color: colors.status.error,
                fontFamily: typography.fonts.primary
            }}>
                <p style={{
                    ...textStyles.body,
                    marginBottom: '1rem',
                    color: colors.status.error
                }}>
                    Failed to load data:
                </p>
                <pre style={{ 
                    fontSize: typography.sizes.xs,
                    fontFamily: typography.fonts.monospace,
                    background: colors.background.secondary,
                    padding: '1rem',
                    borderRadius: '6px',
                    overflow: 'auto',
                    color: colors.text.secondary
                }}>
                    {JSON.stringify(error, null, 2)}
                </pre>
            </div>
        </Card>
    );
};

export default ErrorDisplay;
