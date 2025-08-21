import React from 'react';
import { Button } from 'primereact/button';

const Header: React.FC = () => {
    return (

        <header style={{ padding: '1rem', background: '#222', color: '#fff' }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', minHeight: '100px' }}>
            <div style={{ flex: '1 1 20%', textAlign: 'left' }}>
                <h1 style={{ margin: 0, fontSize: '2.5rem', display: 'flex', alignItems: 'center' }}>
                🥁 DRUMKIT
                </h1>
            </div>
            <div style={{ flex: '1 1 50%', textAlign: 'left' }}>
               
            </div>
            <div style={{ flex: '1 1 0%', textAlign: 'right' }}>
                
            </div>
            </div>
        </header>
    );
};

export default Header;