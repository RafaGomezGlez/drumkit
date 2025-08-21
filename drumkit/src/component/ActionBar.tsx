import React from 'react';
import { Button } from 'primereact/button';
import { typography, colors, textStyles } from '../styles/typography';

interface ActionBarProps {
    onCreateLoad: () => void;
}

const ActionBar: React.FC<ActionBarProps> = ({ onCreateLoad }) => {
    return (
        <div style={{ 
            display: 'flex', 
            justifyContent: 'flex-end', 
            alignItems: 'center', 
            marginBottom: '20px' 
        }}>
            <Button
                label="Create Load"
                icon="pi pi-plus"
                onClick={onCreateLoad}
                style={{
                    backgroundColor: colors.brand.primary,
                    border: 'none',
                    borderRadius: '8px',
                    padding: '0.75rem 1.5rem',
                    fontWeight: typography.weights.medium,
                    fontSize: typography.sizes.base,
                    fontFamily: typography.fonts.primary
                }}
            />
        </div>
    );
};

export default ActionBar;
