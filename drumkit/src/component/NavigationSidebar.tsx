import React from 'react';
import { typography, colors, textStyles } from '../styles/typography';

const NavigationSidebar: React.FC = () => {
    return (
        <div style={{
            width: '400px',
            background: '#222',
            color: '#fff',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center'
        }}>
            <span style={{
                ...textStyles.body,
                color: 'white',
                fontWeight: typography.weights.medium
            }}>
                Navigation
            </span>
        </div>
    );
};

export default NavigationSidebar;
